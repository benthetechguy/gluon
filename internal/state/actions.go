package state

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ProtonMail/gluon/db"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/ids"
	"github.com/ProtonMail/gluon/reporter"
	"github.com/ProtonMail/gluon/rfc822"
	"github.com/bradenaw/juniper/xslices"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

func (state *State) actionCreateAndGetMailbox(ctx context.Context, tx db.Transaction, name string, uidValidity imap.UID) (*db.Mailbox, error) {
	res, err := state.user.GetRemote().CreateMailbox(ctx, strings.Split(name, state.delimiter))
	if err != nil {
		return nil, err
	}

	exists, err := tx.MailboxExistsWithRemoteID(ctx, res.ID)
	if err != nil {
		return nil, err
	}

	if !exists {
		mbox, err := tx.CreateMailbox(
			ctx,
			res.ID,
			strings.Join(res.Name, state.user.GetDelimiter()),
			res.Flags,
			res.PermanentFlags,
			res.Attributes,
			uidValidity,
		)
		if err != nil {
			return nil, err
		}

		return mbox, nil
	}

	return tx.GetMailboxByRemoteID(ctx, res.ID)
}

func (state *State) actionCreateMailbox(ctx context.Context, tx db.Transaction, name string, uidValidity imap.UID) error {
	res, err := state.user.GetRemote().CreateMailbox(ctx, strings.Split(name, state.delimiter))
	if err != nil {
		return err
	}

	return tx.CreateMailboxIfNotExists(ctx, res, state.delimiter, uidValidity)
}

func (state *State) actionDeleteMailbox(ctx context.Context, tx db.Transaction, mboxID db.MailboxIDPair) ([]Update, error) {
	if err := state.user.GetRemote().DeleteMailbox(ctx, mboxID.RemoteID); err != nil {
		return nil, err
	}

	if err := tx.DeleteMailboxWithRemoteID(ctx, mboxID.RemoteID); err != nil {
		return nil, err
	}

	return []Update{NewMailboxDeletedStateUpdate(mboxID.InternalID)}, nil
}

func (state *State) actionUpdateMailbox(ctx context.Context, tx db.Transaction, mboxID imap.MailboxID, newName string) error {
	if err := state.user.GetRemote().UpdateMailbox(
		ctx,
		mboxID,
		strings.Split(newName, state.delimiter),
	); err != nil {
		return err
	}

	return tx.RenameMailboxWithRemoteID(ctx, mboxID, newName)
}

func (state *State) actionCreateMessage(
	ctx context.Context,
	tx db.Transaction,
	mboxID db.MailboxIDPair,
	literal []byte,
	flags imap.FlagSet,
	date time.Time,
	isSelectedMailbox bool,
	cameFromDrafts bool,
) ([]Update, imap.UID, error) {
	internalID, res, newLiteral, err := state.user.GetRemote().CreateMessage(ctx, mboxID.RemoteID, literal, flags, date)
	if err != nil {
		return nil, 0, err
	}

	{
		// Handle the case where duplicate messages can return the same remote ID.
		knownInternalID, knownErr := tx.GetMessageIDFromRemoteID(ctx, res.ID)
		if knownErr != nil && !db.IsErrNotFound(knownErr) {
			return nil, 0, knownErr
		}
		if knownErr == nil {
			// Try to collect the original message date.
			var existingMessageDate time.Time
			if existingMessage, msgErr := tx.GetMessageNoEdges(ctx, internalID); msgErr == nil {
				existingMessageDate = existingMessage.Date
			}

			if cameFromDrafts {
				reporter.ExceptionWithContext(ctx, "Append to drafts must not return an existing RemoteID", reporter.Context{
					"remote-id":     res.ID,
					"new-date":      res.Date,
					"original-date": existingMessageDate,
					"mailbox":       mboxID.RemoteID,
				})

				logrus.Errorf("Append to drafts must not return an existing RemoteID (Remote=%v, Internal=%v)", res.ID, knownInternalID)

				return nil, 0, fmt.Errorf("append to drafts returned an existing remote ID")
			}

			logrus.Debugf("Deduped message detected, adding existing %v message to mailbox instead.", knownInternalID.ShortID())

			updates, result, err := state.actionAddMessagesToMailbox(ctx,
				tx,
				[]db.MessageIDPair{{InternalID: knownInternalID, RemoteID: res.ID}},
				mboxID,
				isSelectedMailbox,
			)
			if err != nil {
				return nil, 0, err
			}

			return updates, result[0].UID, nil
		}
	}

	parsedMessage, err := imap.NewParsedMessage(newLiteral)
	if err != nil {
		return nil, 0, err
	}

	literalWithHeader, literalSize, err := rfc822.SetHeaderValueNoMemCopy(newLiteral, ids.InternalIDKey, internalID.String())
	if err != nil {
		return nil, 0, fmt.Errorf("failed to set internal ID: %w", err)
	}

	if err := state.user.GetStore().SetUnchecked(internalID, literalWithHeader); err != nil {
		return nil, 0, fmt.Errorf("failed to store message literal: %w", err)
	}

	req := db.CreateMessageReq{
		Message:     res,
		LiteralSize: literalSize,
		Body:        parsedMessage.Body,
		Structure:   parsedMessage.Structure,
		Envelope:    parsedMessage.Envelope,
		InternalID:  internalID,
	}

	messageUID, flagSet, err := tx.CreateMessageAndAddToMailbox(ctx, mboxID.InternalID, &req)
	if err != nil {
		return nil, 0, err
	}

	// We can append to non-selected mailboxes.
	var st *State

	if isSelectedMailbox {
		st = state
	}

	updates := []Update{newExistsStateUpdateWithExists(
		mboxID.InternalID,
		[]*exists{newExists(db.MessageIDPair{InternalID: internalID, RemoteID: res.ID}, messageUID, flagSet)},
		st,
	),
	}

	return updates, messageUID, nil
}

func (state *State) actionCreateRecoveredMessage(
	ctx context.Context,
	tx db.Transaction,
	literal []byte,
	flags imap.FlagSet,
	date time.Time,
) ([]Update, bool, error) {
	internalID := imap.NewInternalMessageID()
	remoteID := ids.NewRecoveredRemoteMessageID(internalID)

	parsedMessage, err := imap.NewParsedMessage(literal)
	if err != nil {
		return nil, false, err
	}

	alreadyKnown, err := state.user.GetRecoveredMessageHashesMap().Insert(internalID, literal)
	if err == nil && alreadyKnown {
		// Message is already known to us, so we ignore it.
		return nil, true, nil
	}

	if err := state.user.GetStore().SetUnchecked(internalID, bytes.NewReader(literal)); err != nil {
		return nil, false, fmt.Errorf("failed to store message literal: %w", err)
	}

	req := db.CreateMessageReq{
		Message: imap.Message{
			ID:    remoteID,
			Flags: flags,
			Date:  date,
		},
		LiteralSize: len(literal),
		Body:        parsedMessage.Body,
		Structure:   parsedMessage.Structure,
		Envelope:    parsedMessage.Envelope,
		InternalID:  internalID,
	}

	recoveryMBoxID := state.user.GetRecoveryMailboxID()

	messageUID, flagSet, err := tx.CreateMessageAndAddToMailbox(ctx, recoveryMBoxID.InternalID, &req)
	if err != nil {
		return nil, false, err
	}

	var updates = []Update{newExistsStateUpdateWithExists(
		recoveryMBoxID.InternalID,
		[]*exists{newExists(db.MessageIDPair{InternalID: internalID, RemoteID: remoteID}, messageUID, flagSet)},
		nil,
	),
	}

	return updates, false, nil
}

func (state *State) actionAddMessagesToMailbox(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
	isMailboxSelected bool,
) ([]Update, []db.UIDWithFlags, error) {
	var allUpdates []Update

	{
		haveMessageIDs, err := tx.MailboxFilterContains(ctx, mboxID.InternalID, messageIDs)
		if err != nil {
			return nil, nil, err
		}

		if remMessageIDs := xslices.Filter(messageIDs, func(messageID db.MessageIDPair) bool {
			return slices.Contains(haveMessageIDs, messageID.InternalID)
		}); len(remMessageIDs) > 0 {
			updates, err := state.actionRemoveMessagesFromMailboxUnchecked(ctx, tx, remMessageIDs, mboxID)
			if err != nil {
				return nil, nil, err
			}

			allUpdates = append(allUpdates, updates...)
		}
	}

	remoteIDs := xslices.Map(messageIDs, func(id db.MessageIDPair) imap.MessageID {
		return id.RemoteID
	})

	if err := state.user.GetRemote().AddMessagesToMailbox(ctx, remoteIDs, mboxID.RemoteID); err != nil {
		return nil, nil, err
	}

	// Messages can be added to a mailbox that is not selected.
	var st *State
	if isMailboxSelected {
		st = state
	}

	messageUIDs, update, err := AddMessagesToMailbox(ctx, tx, mboxID.InternalID, messageIDs, st, state.imapLimits)
	if err != nil {
		return nil, nil, err
	}

	allUpdates = append(allUpdates, update)

	return allUpdates, messageUIDs, nil
}

func (state *State) actionAddRecoveredMessagesToMailbox(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
) ([]db.UIDWithFlags, Update, error) {
	filter, err := tx.MailboxFilterContains(ctx, mboxID.InternalID, messageIDs)
	if err != nil {
		return nil, nil, err
	}

	toAdd := xslices.Filter(messageIDs, func(t db.MessageIDPair) bool {
		return !slices.Contains(filter, t.InternalID)
	})

	remoteIDs := xslices.Map(toAdd, func(id db.MessageIDPair) imap.MessageID {
		return id.RemoteID
	})

	if err := state.user.GetRemote().AddMessagesToMailbox(ctx, remoteIDs, mboxID.RemoteID); err != nil {
		return nil, nil, err
	}

	return AddMessagesToMailbox(ctx, tx, mboxID.InternalID, toAdd, state, state.imapLimits)
}

func (state *State) actionImportRecoveredMessage(
	ctx context.Context,
	tx db.Transaction,
	id imap.InternalMessageID,
	mboxID imap.MailboxID,
) (db.MessageIDPair, bool, error) {
	message, err := tx.GetImportedMessageData(ctx, id)
	if err != nil {
		return db.MessageIDPair{}, false, err
	}

	literal, err := state.user.GetStore().Get(id)
	if err != nil {
		return db.MessageIDPair{}, false, err
	}

	internalID, res, newLiteral, err := state.user.GetRemote().CreateMessage(ctx, mboxID, literal, message.Flags, message.Date)
	if err != nil {
		return db.MessageIDPair{}, false, err
	}

	{
		// Handle the unlikely case where duplicate messages can return the same remote ID.
		internalID, err := tx.GetMessageIDFromRemoteID(ctx, res.ID)
		if err != nil && !db.IsErrNotFound(err) {
			return db.MessageIDPair{}, false, err
		}

		if err == nil {
			return db.MessageIDPair{
				InternalID: internalID,
				RemoteID:   res.ID,
			}, true, nil
		}
	}

	parsedMessage, err := imap.NewParsedMessage(newLiteral)
	if err != nil {
		return db.MessageIDPair{}, false, err
	}

	literalReader, literalSize, err := rfc822.SetHeaderValueNoMemCopy(newLiteral, ids.InternalIDKey, internalID.String())
	if err != nil {
		return db.MessageIDPair{}, false, fmt.Errorf("failed to set internal ID: %w", err)
	}

	if err := state.user.GetStore().SetUnchecked(internalID, literalReader); err != nil {
		return db.MessageIDPair{}, false, fmt.Errorf("failed to store message literal: %w", err)
	}

	req := db.CreateMessageReq{
		Message:     res,
		LiteralSize: literalSize,
		Body:        parsedMessage.Body,
		Structure:   parsedMessage.Structure,
		Envelope:    parsedMessage.Envelope,
		InternalID:  internalID,
	}

	if err := tx.CreateMessages(ctx, &req); err != nil {
		return db.MessageIDPair{}, false, err
	}

	return db.MessageIDPair{
		InternalID: internalID,
		RemoteID:   res.ID,
	}, false, nil
}

func (state *State) actionCopyMessagesOutOfRecoveryMailbox(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
) ([]Update, []db.UIDWithFlags, error) {
	ids := make([]db.MessageIDPair, 0, len(messageIDs))

	// Import messages to remote.
	for _, id := range messageIDs {
		id, _, err := state.actionImportRecoveredMessage(ctx, tx, id.InternalID, mboxID.RemoteID)
		if err != nil {
			return nil, nil, err
		}

		ids = append(ids, id)
	}

	// Label messages in destination.
	uidWithFlags, update, err := state.actionAddRecoveredMessagesToMailbox(ctx, tx, ids, mboxID)
	if err != nil {
		return nil, nil, err
	}

	return []Update{update}, uidWithFlags, nil
}

func (state *State) actionMoveMessagesOutOfRecoveryMailbox(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
) ([]Update, []db.UIDWithFlags, error) {
	ids := make([]db.MessageIDPair, 0, len(messageIDs))
	oldInternalIDs := make([]imap.InternalMessageID, 0, len(messageIDs))

	// Import messages to remote.
	for _, id := range messageIDs {
		newID, deduped, err := state.actionImportRecoveredMessage(ctx, tx, id.InternalID, mboxID.RemoteID)
		if err != nil {
			return nil, nil, err
		}

		if !deduped {
			if err := tx.MarkMessageAsDeleted(ctx, id.InternalID); err != nil {
				return nil, nil, err
			}
		}

		ids = append(ids, newID)
		oldInternalIDs = append(oldInternalIDs, id.InternalID)
	}

	// Expunge messages
	var updates []Update
	{
		removeUpdates, err := RemoveMessagesFromMailbox(ctx, tx, state.user.GetRecoveryMailboxID().InternalID, oldInternalIDs)
		if err != nil {
			return nil, nil, err
		}

		state.user.GetRecoveredMessageHashesMap().Erase(oldInternalIDs...)

		updates = append(updates, removeUpdates...)
	}

	// Label messages in destination.
	uidWithFlags, update, err := state.actionAddRecoveredMessagesToMailbox(ctx, tx, ids, mboxID)
	if err != nil {
		return nil, nil, err
	}

	updates = append(updates, update)

	return updates, uidWithFlags, nil
}

// actionRemoveMessagesFromMailboxUnchecked is similar to actionRemoveMessagesFromMailbox, but it does not validate
// the input for whether messages actually exist in the database or if the message set is empty. use this when you
// have already validated the input beforehand (e.g.: actionAddMessagesToMailbox and actionRemoveMessagesFromMailbox).
func (state *State) actionRemoveMessagesFromMailboxUnchecked(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
) ([]Update, error) {
	internalIDs, remoteIDs := db.SplitMessageIDPairSlice(messageIDs)

	if mboxID.InternalID != state.user.GetRecoveryMailboxID().InternalID {
		if err := state.user.GetRemote().RemoveMessagesFromMailbox(ctx, remoteIDs, mboxID.RemoteID); err != nil {
			return nil, err
		}
	} else {
		state.user.GetRecoveredMessageHashesMap().Erase(internalIDs...)
	}

	return RemoveMessagesFromMailbox(ctx, tx, mboxID.InternalID, internalIDs)
}

func (state *State) actionRemoveMessagesFromMailbox(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxID db.MailboxIDPair,
) ([]Update, error) {
	haveMessageIDs, err := tx.MailboxFilterContains(ctx, mboxID.InternalID, messageIDs)
	if err != nil {
		return nil, err
	}

	messageIDs = xslices.Filter(messageIDs, func(messageID db.MessageIDPair) bool {
		return slices.Contains(haveMessageIDs, messageID.InternalID)
	})

	if len(messageIDs) == 0 {
		return nil, nil
	}

	return state.actionRemoveMessagesFromMailboxUnchecked(ctx, tx, messageIDs, mboxID)
}

func (state *State) actionMoveMessages(
	ctx context.Context,
	tx db.Transaction,
	messageIDs []db.MessageIDPair,
	mboxFromID, mboxToID db.MailboxIDPair,
) ([]Update, []db.UIDWithFlags, error) {
	var allUpdates []Update

	if mboxFromID.InternalID == mboxToID.InternalID {
		updates, err := state.actionRemoveMessagesFromMailboxUnchecked(ctx, tx, messageIDs, mboxToID)
		if err != nil {
			return nil, nil, err
		}

		allUpdates = append(allUpdates, updates...)

		updates, uid, err := state.actionAddMessagesToMailbox(ctx, tx, messageIDs, mboxToID, false)
		if err != nil {
			return nil, nil, err
		}

		allUpdates = append(allUpdates, updates...)

		return allUpdates, uid, err
	}

	{
		messageIDsToAdd, err := tx.MailboxFilterContains(ctx, mboxToID.InternalID, messageIDs)
		if err != nil {
			return nil, nil, err
		}

		if remMessageIDs := xslices.Filter(messageIDs, func(messageID db.MessageIDPair) bool {
			return slices.Contains(messageIDsToAdd, messageID.InternalID)
		}); len(remMessageIDs) > 0 {
			updates, err := state.actionRemoveMessagesFromMailboxUnchecked(ctx, tx, remMessageIDs, mboxToID)
			if err != nil {
				return nil, nil, err
			}

			allUpdates = append(allUpdates, updates...)
		}
	}

	messageInFromMBox, err := tx.MailboxFilterContains(ctx, mboxFromID.InternalID, messageIDs)
	if err != nil {
		return nil, nil, err
	}

	messagesIDsToMove := xslices.Filter(messageIDs, func(messageID db.MessageIDPair) bool {
		return slices.Contains(messageInFromMBox, messageID.InternalID)
	})

	internalIDs, remoteIDs := db.SplitMessageIDPairSlice(messagesIDsToMove)

	shouldRemoveOldMessages, err := state.user.GetRemote().MoveMessagesFromMailbox(ctx, remoteIDs, mboxFromID.RemoteID, mboxToID.RemoteID)
	if err != nil {
		return nil, nil, err
	}

	messageUIDs, updates, err := MoveMessagesFromMailbox(
		ctx,
		tx,
		mboxFromID.InternalID,
		mboxToID.InternalID,
		messagesIDsToMove,
		internalIDs,
		state,
		state.imapLimits,
		shouldRemoveOldMessages,
	)
	if err != nil {
		return nil, nil, err
	}

	allUpdates = append(allUpdates, updates...)

	return allUpdates, messageUIDs, nil
}

func (state *State) actionAddMessageFlags(
	ctx context.Context,
	tx db.Transaction,
	messages []snapMsgWithSeq,
	addFlags imap.FlagSet,
) ([]Update, error) {
	internalMessageIDs := xslices.Map(messages, func(sm snapMsgWithSeq) imap.InternalMessageID {
		return sm.ID.InternalID
	})

	return state.applyMessageFlagsAdded(ctx, tx, internalMessageIDs, addFlags)
}

func (state *State) actionRemoveMessageFlags(
	ctx context.Context,
	tx db.Transaction,
	messages []snapMsgWithSeq,
	remFlags imap.FlagSet,
) ([]Update, error) {
	internalMessageIDs := xslices.Map(messages, func(sm snapMsgWithSeq) imap.InternalMessageID {
		return sm.ID.InternalID
	})

	return state.applyMessageFlagsRemoved(ctx, tx, internalMessageIDs, remFlags)
}

func (state *State) actionSetMessageFlags(ctx context.Context,
	tx db.Transaction,
	messages []snapMsgWithSeq,
	setFlags imap.FlagSet) ([]Update, error) {
	if setFlags.ContainsUnchecked(imap.FlagRecentLowerCase) {
		return nil, fmt.Errorf("recent flag is read-only")
	}

	internalMessageIDs := xslices.Map(messages, func(sm snapMsgWithSeq) imap.InternalMessageID {
		return sm.ID.InternalID
	})

	return state.applyMessageFlagsSet(ctx, tx, internalMessageIDs, setFlags)
}
