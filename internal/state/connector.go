package state

import (
	"context"
	"time"

	"github.com/ProtonMail/gluon/imap"
)

// Connector interface for State differs slightly from the connector.Connector interface as it needs the ability
// to generate internal IDs for each request as well as track local metadata associated with each state. The local
// metadata (e.g.: IMAP ID extension info) should be injected into the context before each call to ensure the
// connector.Connector can receive this information.
// Sadly, due to Go's cyclic dependencies, this needs to be an interface. The implementation of this interface
// is available in the backend package.
type Connector interface {
	// SetConnMetadataValue sets a metadata value associated with the current connector.
	SetConnMetadataValue(key string, value any)

	// ClearConnMetadataValue clears a metadata value associated with the current connector.
	ClearConnMetadataValue(key string)

	// ClearAllConnMetadata clears all metadata values associated with the current connector.
	ClearAllConnMetadata()

	// CreateMailbox creates a new mailbox with the given name.
	CreateMailbox(ctx context.Context, name []string) (imap.InternalMailboxID, imap.Mailbox, error)

	// UpdateMailbox sets the name of the mailbox with the given ID to the given new name.
	UpdateMailbox(ctx context.Context, mboxID imap.LabelID, oldName, newName []string) error

	// DeleteMailbox deletes the mailbox with the given ID and name.
	DeleteMailbox(ctx context.Context, mboxID imap.LabelID) error

	// CreateMessage appends a message literal to the mailbox with the given ID.
	CreateMessage(
		ctx context.Context,
		mboxID imap.LabelID,
		literal []byte,
		message *imap.ParsedMessage,
		flags imap.FlagSet,
		date time.Time,
	) (imap.InternalMessageID, imap.Message, error)

	// AddMessagesToMailbox adds the message with the given ID to the mailbox with the given ID.
	AddMessagesToMailbox(
		ctx context.Context,
		messageIDs []imap.MessageID,
		mboxID imap.LabelID,
	) error

	// RemoveMessagesFromMailbox removes the message with the given ID from the mailbox with the given ID.
	RemoveMessagesFromMailbox(
		ctx context.Context,
		messageIDs []imap.MessageID,
		mboxID imap.LabelID,
	) error

	// MoveMessagesFromMailbox removes the message with the given ID from the mailbox with the given ID.
	MoveMessagesFromMailbox(
		ctx context.Context,
		messageIDs []imap.MessageID,
		mboxFromID imap.LabelID,
		mboxToID imap.LabelID,
	) error

	// SetMessagesSeen marks the message with the given ID as seen or unseen.
	SetMessagesSeen(ctx context.Context, messageIDs []imap.MessageID, seen bool) error

	// SetMessagesFlagged marks the message with the given ID as seen or unseen.
	SetMessagesFlagged(ctx context.Context, messageIDs []imap.MessageID, flagged bool) error
}
