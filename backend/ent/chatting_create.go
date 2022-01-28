// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/chatting"
	"FinalProject/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChattingCreate is the builder for creating a Chatting entity.
type ChattingCreate struct {
	config
	mutation *ChattingMutation
	hooks    []Hook
}

// SetMessage sets the "message" field.
func (cc *ChattingCreate) SetMessage(s string) *ChattingCreate {
	cc.mutation.SetMessage(s)
	return cc
}

// SetDateTime sets the "dateTime" field.
func (cc *ChattingCreate) SetDateTime(t time.Time) *ChattingCreate {
	cc.mutation.SetDateTime(t)
	return cc
}

// SetChattingWithWhomID sets the "chatting_with_whom" edge to the User entity by ID.
func (cc *ChattingCreate) SetChattingWithWhomID(id int) *ChattingCreate {
	cc.mutation.SetChattingWithWhomID(id)
	return cc
}

// SetNillableChattingWithWhomID sets the "chatting_with_whom" edge to the User entity by ID if the given value is not nil.
func (cc *ChattingCreate) SetNillableChattingWithWhomID(id *int) *ChattingCreate {
	if id != nil {
		cc = cc.SetChattingWithWhomID(*id)
	}
	return cc
}

// SetChattingWithWhom sets the "chatting_with_whom" edge to the User entity.
func (cc *ChattingCreate) SetChattingWithWhom(u *User) *ChattingCreate {
	return cc.SetChattingWithWhomID(u.ID)
}

// SetWhoseIsThisMsgID sets the "whose_is_this_msg" edge to the User entity by ID.
func (cc *ChattingCreate) SetWhoseIsThisMsgID(id int) *ChattingCreate {
	cc.mutation.SetWhoseIsThisMsgID(id)
	return cc
}

// SetNillableWhoseIsThisMsgID sets the "whose_is_this_msg" edge to the User entity by ID if the given value is not nil.
func (cc *ChattingCreate) SetNillableWhoseIsThisMsgID(id *int) *ChattingCreate {
	if id != nil {
		cc = cc.SetWhoseIsThisMsgID(*id)
	}
	return cc
}

// SetWhoseIsThisMsg sets the "whose_is_this_msg" edge to the User entity.
func (cc *ChattingCreate) SetWhoseIsThisMsg(u *User) *ChattingCreate {
	return cc.SetWhoseIsThisMsgID(u.ID)
}

// Mutation returns the ChattingMutation object of the builder.
func (cc *ChattingCreate) Mutation() *ChattingMutation {
	return cc.mutation
}

// Save creates the Chatting in the database.
func (cc *ChattingCreate) Save(ctx context.Context) (*Chatting, error) {
	var (
		err  error
		node *Chatting
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChattingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChattingCreate) SaveX(ctx context.Context) *Chatting {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChattingCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChattingCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChattingCreate) check() error {
	if _, ok := cc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := cc.mutation.DateTime(); !ok {
		return &ValidationError{Name: "dateTime", err: errors.New(`ent: missing required field "dateTime"`)}
	}
	return nil
}

func (cc *ChattingCreate) sqlSave(ctx context.Context) (*Chatting, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ChattingCreate) createSpec() (*Chatting, *sqlgraph.CreateSpec) {
	var (
		_node = &Chatting{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chatting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chatting.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatting.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := cc.mutation.DateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatting.FieldDateTime,
		})
		_node.DateTime = value
	}
	if nodes := cc.mutation.ChattingWithWhomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatting.ChattingWithWhomTable,
			Columns: []string{chatting.ChattingWithWhomColumn},
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
		_node.user_user_chatting_with_whom = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.WhoseIsThisMsgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatting.WhoseIsThisMsgTable,
			Columns: []string{chatting.WhoseIsThisMsgColumn},
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
		_node.user_who_is_owner_this_msg = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChattingCreateBulk is the builder for creating many Chatting entities in bulk.
type ChattingCreateBulk struct {
	config
	builders []*ChattingCreate
}

// Save creates the Chatting entities in the database.
func (ccb *ChattingCreateBulk) Save(ctx context.Context) ([]*Chatting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chatting, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChattingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChattingCreateBulk) SaveX(ctx context.Context) []*Chatting {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChattingCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChattingCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
