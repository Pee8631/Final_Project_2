// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/predicate"
	"FinalProject/ent/scheduletime"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleTimeDelete is the builder for deleting a ScheduleTime entity.
type ScheduleTimeDelete struct {
	config
	hooks    []Hook
	mutation *ScheduleTimeMutation
}

// Where appends a list predicates to the ScheduleTimeDelete builder.
func (std *ScheduleTimeDelete) Where(ps ...predicate.ScheduleTime) *ScheduleTimeDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *ScheduleTimeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(std.hooks) == 0 {
		affected, err = std.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleTimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			std.mutation = mutation
			affected, err = std.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(std.hooks) - 1; i >= 0; i-- {
			if std.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = std.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, std.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (std *ScheduleTimeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *ScheduleTimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: scheduletime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scheduletime.FieldID,
			},
		},
	}
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, std.driver, _spec)
}

// ScheduleTimeDeleteOne is the builder for deleting a single ScheduleTime entity.
type ScheduleTimeDeleteOne struct {
	std *ScheduleTimeDelete
}

// Exec executes the deletion query.
func (stdo *ScheduleTimeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scheduletime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *ScheduleTimeDeleteOne) ExecX(ctx context.Context) {
	stdo.std.ExecX(ctx)
}