// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/db/ent/mailbox"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxattr"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxflag"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxpermflag"
	"github.com/ProtonMail/gluon/internal/db/ent/predicate"
	"github.com/ProtonMail/gluon/internal/db/ent/uid"
)

// MailboxUpdate is the builder for updating Mailbox entities.
type MailboxUpdate struct {
	config
	hooks    []Hook
	mutation *MailboxMutation
}

// Where appends a list predicates to the MailboxUpdate builder.
func (mu *MailboxUpdate) Where(ps ...predicate.Mailbox) *MailboxUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetRemoteID sets the "RemoteID" field.
func (mu *MailboxUpdate) SetRemoteID(ii imap.LabelID) *MailboxUpdate {
	mu.mutation.SetRemoteID(ii)
	return mu
}

// SetNillableRemoteID sets the "RemoteID" field if the given value is not nil.
func (mu *MailboxUpdate) SetNillableRemoteID(ii *imap.LabelID) *MailboxUpdate {
	if ii != nil {
		mu.SetRemoteID(*ii)
	}
	return mu
}

// ClearRemoteID clears the value of the "RemoteID" field.
func (mu *MailboxUpdate) ClearRemoteID() *MailboxUpdate {
	mu.mutation.ClearRemoteID()
	return mu
}

// SetName sets the "Name" field.
func (mu *MailboxUpdate) SetName(s string) *MailboxUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetUIDNext sets the "UIDNext" field.
func (mu *MailboxUpdate) SetUIDNext(i imap.UID) *MailboxUpdate {
	mu.mutation.ResetUIDNext()
	mu.mutation.SetUIDNext(i)
	return mu
}

// SetNillableUIDNext sets the "UIDNext" field if the given value is not nil.
func (mu *MailboxUpdate) SetNillableUIDNext(i *imap.UID) *MailboxUpdate {
	if i != nil {
		mu.SetUIDNext(*i)
	}
	return mu
}

// AddUIDNext adds i to the "UIDNext" field.
func (mu *MailboxUpdate) AddUIDNext(i imap.UID) *MailboxUpdate {
	mu.mutation.AddUIDNext(i)
	return mu
}

// SetUIDValidity sets the "UIDValidity" field.
func (mu *MailboxUpdate) SetUIDValidity(i imap.UID) *MailboxUpdate {
	mu.mutation.ResetUIDValidity()
	mu.mutation.SetUIDValidity(i)
	return mu
}

// SetNillableUIDValidity sets the "UIDValidity" field if the given value is not nil.
func (mu *MailboxUpdate) SetNillableUIDValidity(i *imap.UID) *MailboxUpdate {
	if i != nil {
		mu.SetUIDValidity(*i)
	}
	return mu
}

// AddUIDValidity adds i to the "UIDValidity" field.
func (mu *MailboxUpdate) AddUIDValidity(i imap.UID) *MailboxUpdate {
	mu.mutation.AddUIDValidity(i)
	return mu
}

// SetSubscribed sets the "Subscribed" field.
func (mu *MailboxUpdate) SetSubscribed(b bool) *MailboxUpdate {
	mu.mutation.SetSubscribed(b)
	return mu
}

// SetNillableSubscribed sets the "Subscribed" field if the given value is not nil.
func (mu *MailboxUpdate) SetNillableSubscribed(b *bool) *MailboxUpdate {
	if b != nil {
		mu.SetSubscribed(*b)
	}
	return mu
}

// AddUIDIDs adds the "UIDs" edge to the UID entity by IDs.
func (mu *MailboxUpdate) AddUIDIDs(ids ...int) *MailboxUpdate {
	mu.mutation.AddUIDIDs(ids...)
	return mu
}

// AddUIDs adds the "UIDs" edges to the UID entity.
func (mu *MailboxUpdate) AddUIDs(u ...*UID) *MailboxUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.AddUIDIDs(ids...)
}

// AddFlagIDs adds the "flags" edge to the MailboxFlag entity by IDs.
func (mu *MailboxUpdate) AddFlagIDs(ids ...int) *MailboxUpdate {
	mu.mutation.AddFlagIDs(ids...)
	return mu
}

// AddFlags adds the "flags" edges to the MailboxFlag entity.
func (mu *MailboxUpdate) AddFlags(m ...*MailboxFlag) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddFlagIDs(ids...)
}

// AddPermanentFlagIDs adds the "permanent_flags" edge to the MailboxPermFlag entity by IDs.
func (mu *MailboxUpdate) AddPermanentFlagIDs(ids ...int) *MailboxUpdate {
	mu.mutation.AddPermanentFlagIDs(ids...)
	return mu
}

// AddPermanentFlags adds the "permanent_flags" edges to the MailboxPermFlag entity.
func (mu *MailboxUpdate) AddPermanentFlags(m ...*MailboxPermFlag) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddPermanentFlagIDs(ids...)
}

// AddAttributeIDs adds the "attributes" edge to the MailboxAttr entity by IDs.
func (mu *MailboxUpdate) AddAttributeIDs(ids ...int) *MailboxUpdate {
	mu.mutation.AddAttributeIDs(ids...)
	return mu
}

// AddAttributes adds the "attributes" edges to the MailboxAttr entity.
func (mu *MailboxUpdate) AddAttributes(m ...*MailboxAttr) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddAttributeIDs(ids...)
}

// Mutation returns the MailboxMutation object of the builder.
func (mu *MailboxUpdate) Mutation() *MailboxMutation {
	return mu.mutation
}

// ClearUIDs clears all "UIDs" edges to the UID entity.
func (mu *MailboxUpdate) ClearUIDs() *MailboxUpdate {
	mu.mutation.ClearUIDs()
	return mu
}

// RemoveUIDIDs removes the "UIDs" edge to UID entities by IDs.
func (mu *MailboxUpdate) RemoveUIDIDs(ids ...int) *MailboxUpdate {
	mu.mutation.RemoveUIDIDs(ids...)
	return mu
}

// RemoveUIDs removes "UIDs" edges to UID entities.
func (mu *MailboxUpdate) RemoveUIDs(u ...*UID) *MailboxUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.RemoveUIDIDs(ids...)
}

// ClearFlags clears all "flags" edges to the MailboxFlag entity.
func (mu *MailboxUpdate) ClearFlags() *MailboxUpdate {
	mu.mutation.ClearFlags()
	return mu
}

// RemoveFlagIDs removes the "flags" edge to MailboxFlag entities by IDs.
func (mu *MailboxUpdate) RemoveFlagIDs(ids ...int) *MailboxUpdate {
	mu.mutation.RemoveFlagIDs(ids...)
	return mu
}

// RemoveFlags removes "flags" edges to MailboxFlag entities.
func (mu *MailboxUpdate) RemoveFlags(m ...*MailboxFlag) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveFlagIDs(ids...)
}

// ClearPermanentFlags clears all "permanent_flags" edges to the MailboxPermFlag entity.
func (mu *MailboxUpdate) ClearPermanentFlags() *MailboxUpdate {
	mu.mutation.ClearPermanentFlags()
	return mu
}

// RemovePermanentFlagIDs removes the "permanent_flags" edge to MailboxPermFlag entities by IDs.
func (mu *MailboxUpdate) RemovePermanentFlagIDs(ids ...int) *MailboxUpdate {
	mu.mutation.RemovePermanentFlagIDs(ids...)
	return mu
}

// RemovePermanentFlags removes "permanent_flags" edges to MailboxPermFlag entities.
func (mu *MailboxUpdate) RemovePermanentFlags(m ...*MailboxPermFlag) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemovePermanentFlagIDs(ids...)
}

// ClearAttributes clears all "attributes" edges to the MailboxAttr entity.
func (mu *MailboxUpdate) ClearAttributes() *MailboxUpdate {
	mu.mutation.ClearAttributes()
	return mu
}

// RemoveAttributeIDs removes the "attributes" edge to MailboxAttr entities by IDs.
func (mu *MailboxUpdate) RemoveAttributeIDs(ids ...int) *MailboxUpdate {
	mu.mutation.RemoveAttributeIDs(ids...)
	return mu
}

// RemoveAttributes removes "attributes" edges to MailboxAttr entities.
func (mu *MailboxUpdate) RemoveAttributes(m ...*MailboxAttr) *MailboxUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveAttributeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MailboxUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MailboxUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MailboxUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MailboxUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MailboxUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailbox.Table,
			Columns: mailbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mailbox.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.RemoteID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldRemoteID,
		})
	}
	if mu.mutation.RemoteIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: mailbox.FieldRemoteID,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldName,
		})
	}
	if value, ok := mu.mutation.UIDNext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDNext,
		})
	}
	if value, ok := mu.mutation.AddedUIDNext(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDNext,
		})
	}
	if value, ok := mu.mutation.UIDValidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDValidity,
		})
	}
	if value, ok := mu.mutation.AddedUIDValidity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDValidity,
		})
	}
	if value, ok := mu.mutation.Subscribed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mailbox.FieldSubscribed,
		})
	}
	if mu.mutation.UIDsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedUIDsIDs(); len(nodes) > 0 && !mu.mutation.UIDsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UIDsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.FlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedFlagsIDs(); len(nodes) > 0 && !mu.mutation.FlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.FlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.PermanentFlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPermanentFlagsIDs(); len(nodes) > 0 && !mu.mutation.PermanentFlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PermanentFlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !mu.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MailboxUpdateOne is the builder for updating a single Mailbox entity.
type MailboxUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MailboxMutation
}

// SetRemoteID sets the "RemoteID" field.
func (muo *MailboxUpdateOne) SetRemoteID(ii imap.LabelID) *MailboxUpdateOne {
	muo.mutation.SetRemoteID(ii)
	return muo
}

// SetNillableRemoteID sets the "RemoteID" field if the given value is not nil.
func (muo *MailboxUpdateOne) SetNillableRemoteID(ii *imap.LabelID) *MailboxUpdateOne {
	if ii != nil {
		muo.SetRemoteID(*ii)
	}
	return muo
}

// ClearRemoteID clears the value of the "RemoteID" field.
func (muo *MailboxUpdateOne) ClearRemoteID() *MailboxUpdateOne {
	muo.mutation.ClearRemoteID()
	return muo
}

// SetName sets the "Name" field.
func (muo *MailboxUpdateOne) SetName(s string) *MailboxUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetUIDNext sets the "UIDNext" field.
func (muo *MailboxUpdateOne) SetUIDNext(i imap.UID) *MailboxUpdateOne {
	muo.mutation.ResetUIDNext()
	muo.mutation.SetUIDNext(i)
	return muo
}

// SetNillableUIDNext sets the "UIDNext" field if the given value is not nil.
func (muo *MailboxUpdateOne) SetNillableUIDNext(i *imap.UID) *MailboxUpdateOne {
	if i != nil {
		muo.SetUIDNext(*i)
	}
	return muo
}

// AddUIDNext adds i to the "UIDNext" field.
func (muo *MailboxUpdateOne) AddUIDNext(i imap.UID) *MailboxUpdateOne {
	muo.mutation.AddUIDNext(i)
	return muo
}

// SetUIDValidity sets the "UIDValidity" field.
func (muo *MailboxUpdateOne) SetUIDValidity(i imap.UID) *MailboxUpdateOne {
	muo.mutation.ResetUIDValidity()
	muo.mutation.SetUIDValidity(i)
	return muo
}

// SetNillableUIDValidity sets the "UIDValidity" field if the given value is not nil.
func (muo *MailboxUpdateOne) SetNillableUIDValidity(i *imap.UID) *MailboxUpdateOne {
	if i != nil {
		muo.SetUIDValidity(*i)
	}
	return muo
}

// AddUIDValidity adds i to the "UIDValidity" field.
func (muo *MailboxUpdateOne) AddUIDValidity(i imap.UID) *MailboxUpdateOne {
	muo.mutation.AddUIDValidity(i)
	return muo
}

// SetSubscribed sets the "Subscribed" field.
func (muo *MailboxUpdateOne) SetSubscribed(b bool) *MailboxUpdateOne {
	muo.mutation.SetSubscribed(b)
	return muo
}

// SetNillableSubscribed sets the "Subscribed" field if the given value is not nil.
func (muo *MailboxUpdateOne) SetNillableSubscribed(b *bool) *MailboxUpdateOne {
	if b != nil {
		muo.SetSubscribed(*b)
	}
	return muo
}

// AddUIDIDs adds the "UIDs" edge to the UID entity by IDs.
func (muo *MailboxUpdateOne) AddUIDIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.AddUIDIDs(ids...)
	return muo
}

// AddUIDs adds the "UIDs" edges to the UID entity.
func (muo *MailboxUpdateOne) AddUIDs(u ...*UID) *MailboxUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.AddUIDIDs(ids...)
}

// AddFlagIDs adds the "flags" edge to the MailboxFlag entity by IDs.
func (muo *MailboxUpdateOne) AddFlagIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.AddFlagIDs(ids...)
	return muo
}

// AddFlags adds the "flags" edges to the MailboxFlag entity.
func (muo *MailboxUpdateOne) AddFlags(m ...*MailboxFlag) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddFlagIDs(ids...)
}

// AddPermanentFlagIDs adds the "permanent_flags" edge to the MailboxPermFlag entity by IDs.
func (muo *MailboxUpdateOne) AddPermanentFlagIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.AddPermanentFlagIDs(ids...)
	return muo
}

// AddPermanentFlags adds the "permanent_flags" edges to the MailboxPermFlag entity.
func (muo *MailboxUpdateOne) AddPermanentFlags(m ...*MailboxPermFlag) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddPermanentFlagIDs(ids...)
}

// AddAttributeIDs adds the "attributes" edge to the MailboxAttr entity by IDs.
func (muo *MailboxUpdateOne) AddAttributeIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.AddAttributeIDs(ids...)
	return muo
}

// AddAttributes adds the "attributes" edges to the MailboxAttr entity.
func (muo *MailboxUpdateOne) AddAttributes(m ...*MailboxAttr) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddAttributeIDs(ids...)
}

// Mutation returns the MailboxMutation object of the builder.
func (muo *MailboxUpdateOne) Mutation() *MailboxMutation {
	return muo.mutation
}

// ClearUIDs clears all "UIDs" edges to the UID entity.
func (muo *MailboxUpdateOne) ClearUIDs() *MailboxUpdateOne {
	muo.mutation.ClearUIDs()
	return muo
}

// RemoveUIDIDs removes the "UIDs" edge to UID entities by IDs.
func (muo *MailboxUpdateOne) RemoveUIDIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.RemoveUIDIDs(ids...)
	return muo
}

// RemoveUIDs removes "UIDs" edges to UID entities.
func (muo *MailboxUpdateOne) RemoveUIDs(u ...*UID) *MailboxUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.RemoveUIDIDs(ids...)
}

// ClearFlags clears all "flags" edges to the MailboxFlag entity.
func (muo *MailboxUpdateOne) ClearFlags() *MailboxUpdateOne {
	muo.mutation.ClearFlags()
	return muo
}

// RemoveFlagIDs removes the "flags" edge to MailboxFlag entities by IDs.
func (muo *MailboxUpdateOne) RemoveFlagIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.RemoveFlagIDs(ids...)
	return muo
}

// RemoveFlags removes "flags" edges to MailboxFlag entities.
func (muo *MailboxUpdateOne) RemoveFlags(m ...*MailboxFlag) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveFlagIDs(ids...)
}

// ClearPermanentFlags clears all "permanent_flags" edges to the MailboxPermFlag entity.
func (muo *MailboxUpdateOne) ClearPermanentFlags() *MailboxUpdateOne {
	muo.mutation.ClearPermanentFlags()
	return muo
}

// RemovePermanentFlagIDs removes the "permanent_flags" edge to MailboxPermFlag entities by IDs.
func (muo *MailboxUpdateOne) RemovePermanentFlagIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.RemovePermanentFlagIDs(ids...)
	return muo
}

// RemovePermanentFlags removes "permanent_flags" edges to MailboxPermFlag entities.
func (muo *MailboxUpdateOne) RemovePermanentFlags(m ...*MailboxPermFlag) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemovePermanentFlagIDs(ids...)
}

// ClearAttributes clears all "attributes" edges to the MailboxAttr entity.
func (muo *MailboxUpdateOne) ClearAttributes() *MailboxUpdateOne {
	muo.mutation.ClearAttributes()
	return muo
}

// RemoveAttributeIDs removes the "attributes" edge to MailboxAttr entities by IDs.
func (muo *MailboxUpdateOne) RemoveAttributeIDs(ids ...int) *MailboxUpdateOne {
	muo.mutation.RemoveAttributeIDs(ids...)
	return muo
}

// RemoveAttributes removes "attributes" edges to MailboxAttr entities.
func (muo *MailboxUpdateOne) RemoveAttributes(m ...*MailboxAttr) *MailboxUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveAttributeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MailboxUpdateOne) Select(field string, fields ...string) *MailboxUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mailbox entity.
func (muo *MailboxUpdateOne) Save(ctx context.Context) (*Mailbox, error) {
	var (
		err  error
		node *Mailbox
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Mailbox)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MailboxMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MailboxUpdateOne) SaveX(ctx context.Context) *Mailbox {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MailboxUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MailboxUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MailboxUpdateOne) sqlSave(ctx context.Context) (_node *Mailbox, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mailbox.Table,
			Columns: mailbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: mailbox.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mailbox.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailbox.FieldID)
		for _, f := range fields {
			if !mailbox.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mailbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.RemoteID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldRemoteID,
		})
	}
	if muo.mutation.RemoteIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: mailbox.FieldRemoteID,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailbox.FieldName,
		})
	}
	if value, ok := muo.mutation.UIDNext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDNext,
		})
	}
	if value, ok := muo.mutation.AddedUIDNext(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDNext,
		})
	}
	if value, ok := muo.mutation.UIDValidity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDValidity,
		})
	}
	if value, ok := muo.mutation.AddedUIDValidity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: mailbox.FieldUIDValidity,
		})
	}
	if value, ok := muo.mutation.Subscribed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: mailbox.FieldSubscribed,
		})
	}
	if muo.mutation.UIDsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedUIDsIDs(); len(nodes) > 0 && !muo.mutation.UIDsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UIDsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.UIDsTable,
			Columns: []string{mailbox.UIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: uid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.FlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedFlagsIDs(); len(nodes) > 0 && !muo.mutation.FlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.FlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.FlagsTable,
			Columns: []string{mailbox.FlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.PermanentFlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPermanentFlagsIDs(); len(nodes) > 0 && !muo.mutation.PermanentFlagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PermanentFlagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.PermanentFlagsTable,
			Columns: []string{mailbox.PermanentFlagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxpermflag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !muo.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mailbox.AttributesTable,
			Columns: []string{mailbox.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mailboxattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mailbox{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
