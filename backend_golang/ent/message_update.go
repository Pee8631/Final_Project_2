// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/chat"
	"FinalProject/ent/message"
	"FinalProject/ent/predicate"
	"FinalProject/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetMessageText sets the "message_text" field.
func (mu *MessageUpdate) SetMessageText(s string) *MessageUpdate {
	mu.mutation.SetMessageText(s)
	return mu
}

// SetSentDateTime sets the "sent_dateTime" field.
func (mu *MessageUpdate) SetSentDateTime(t time.Time) *MessageUpdate {
	mu.mutation.SetSentDateTime(t)
	return mu
}

// SetNillableSentDateTime sets the "sent_dateTime" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableSentDateTime(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetSentDateTime(*t)
	}
	return mu
}

// ClearSentDateTime clears the value of the "sent_dateTime" field.
func (mu *MessageUpdate) ClearSentDateTime() *MessageUpdate {
	mu.mutation.ClearSentDateTime()
	return mu
}

// SetWhatMessagesAreInThisChatID sets the "What_messages_are_in_this_chat" edge to the Chat entity by ID.
func (mu *MessageUpdate) SetWhatMessagesAreInThisChatID(id int) *MessageUpdate {
	mu.mutation.SetWhatMessagesAreInThisChatID(id)
	return mu
}

// SetNillableWhatMessagesAreInThisChatID sets the "What_messages_are_in_this_chat" edge to the Chat entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableWhatMessagesAreInThisChatID(id *int) *MessageUpdate {
	if id != nil {
		mu = mu.SetWhatMessagesAreInThisChatID(*id)
	}
	return mu
}

// SetWhatMessagesAreInThisChat sets the "What_messages_are_in_this_chat" edge to the Chat entity.
func (mu *MessageUpdate) SetWhatMessagesAreInThisChat(c *Chat) *MessageUpdate {
	return mu.SetWhatMessagesAreInThisChatID(c.ID)
}

// SetWhoSendMessagesID sets the "Who_send_messages" edge to the User entity by ID.
func (mu *MessageUpdate) SetWhoSendMessagesID(id int) *MessageUpdate {
	mu.mutation.SetWhoSendMessagesID(id)
	return mu
}

// SetNillableWhoSendMessagesID sets the "Who_send_messages" edge to the User entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableWhoSendMessagesID(id *int) *MessageUpdate {
	if id != nil {
		mu = mu.SetWhoSendMessagesID(*id)
	}
	return mu
}

// SetWhoSendMessages sets the "Who_send_messages" edge to the User entity.
func (mu *MessageUpdate) SetWhoSendMessages(u *User) *MessageUpdate {
	return mu.SetWhoSendMessagesID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearWhatMessagesAreInThisChat clears the "What_messages_are_in_this_chat" edge to the Chat entity.
func (mu *MessageUpdate) ClearWhatMessagesAreInThisChat() *MessageUpdate {
	mu.mutation.ClearWhatMessagesAreInThisChat()
	return mu
}

// ClearWhoSendMessages clears the "Who_send_messages" edge to the User entity.
func (mu *MessageUpdate) ClearWhoSendMessages() *MessageUpdate {
	mu.mutation.ClearWhoSendMessages()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
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
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
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
	if value, ok := mu.mutation.MessageText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageText,
		})
	}
	if value, ok := mu.mutation.SentDateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldSentDateTime,
		})
	}
	if mu.mutation.SentDateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: message.FieldSentDateTime,
		})
	}
	if mu.mutation.WhatMessagesAreInThisChatCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhatMessagesAreInThisChatTable,
			Columns: []string{message.WhatMessagesAreInThisChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.WhatMessagesAreInThisChatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhatMessagesAreInThisChatTable,
			Columns: []string{message.WhatMessagesAreInThisChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.WhoSendMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhoSendMessagesTable,
			Columns: []string{message.WhoSendMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.WhoSendMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhoSendMessagesTable,
			Columns: []string{message.WhoSendMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
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
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetMessageText sets the "message_text" field.
func (muo *MessageUpdateOne) SetMessageText(s string) *MessageUpdateOne {
	muo.mutation.SetMessageText(s)
	return muo
}

// SetSentDateTime sets the "sent_dateTime" field.
func (muo *MessageUpdateOne) SetSentDateTime(t time.Time) *MessageUpdateOne {
	muo.mutation.SetSentDateTime(t)
	return muo
}

// SetNillableSentDateTime sets the "sent_dateTime" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableSentDateTime(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetSentDateTime(*t)
	}
	return muo
}

// ClearSentDateTime clears the value of the "sent_dateTime" field.
func (muo *MessageUpdateOne) ClearSentDateTime() *MessageUpdateOne {
	muo.mutation.ClearSentDateTime()
	return muo
}

// SetWhatMessagesAreInThisChatID sets the "What_messages_are_in_this_chat" edge to the Chat entity by ID.
func (muo *MessageUpdateOne) SetWhatMessagesAreInThisChatID(id int) *MessageUpdateOne {
	muo.mutation.SetWhatMessagesAreInThisChatID(id)
	return muo
}

// SetNillableWhatMessagesAreInThisChatID sets the "What_messages_are_in_this_chat" edge to the Chat entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableWhatMessagesAreInThisChatID(id *int) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetWhatMessagesAreInThisChatID(*id)
	}
	return muo
}

// SetWhatMessagesAreInThisChat sets the "What_messages_are_in_this_chat" edge to the Chat entity.
func (muo *MessageUpdateOne) SetWhatMessagesAreInThisChat(c *Chat) *MessageUpdateOne {
	return muo.SetWhatMessagesAreInThisChatID(c.ID)
}

// SetWhoSendMessagesID sets the "Who_send_messages" edge to the User entity by ID.
func (muo *MessageUpdateOne) SetWhoSendMessagesID(id int) *MessageUpdateOne {
	muo.mutation.SetWhoSendMessagesID(id)
	return muo
}

// SetNillableWhoSendMessagesID sets the "Who_send_messages" edge to the User entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableWhoSendMessagesID(id *int) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetWhoSendMessagesID(*id)
	}
	return muo
}

// SetWhoSendMessages sets the "Who_send_messages" edge to the User entity.
func (muo *MessageUpdateOne) SetWhoSendMessages(u *User) *MessageUpdateOne {
	return muo.SetWhoSendMessagesID(u.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearWhatMessagesAreInThisChat clears the "What_messages_are_in_this_chat" edge to the Chat entity.
func (muo *MessageUpdateOne) ClearWhatMessagesAreInThisChat() *MessageUpdateOne {
	muo.mutation.ClearWhatMessagesAreInThisChat()
	return muo
}

// ClearWhoSendMessages clears the "Who_send_messages" edge to the User entity.
func (muo *MessageUpdateOne) ClearWhoSendMessages() *MessageUpdateOne {
	muo.mutation.ClearWhoSendMessages()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
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
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
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
	if value, ok := muo.mutation.MessageText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageText,
		})
	}
	if value, ok := muo.mutation.SentDateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldSentDateTime,
		})
	}
	if muo.mutation.SentDateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: message.FieldSentDateTime,
		})
	}
	if muo.mutation.WhatMessagesAreInThisChatCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhatMessagesAreInThisChatTable,
			Columns: []string{message.WhatMessagesAreInThisChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.WhatMessagesAreInThisChatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhatMessagesAreInThisChatTable,
			Columns: []string{message.WhatMessagesAreInThisChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.WhoSendMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhoSendMessagesTable,
			Columns: []string{message.WhoSendMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.WhoSendMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.WhoSendMessagesTable,
			Columns: []string{message.WhoSendMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
