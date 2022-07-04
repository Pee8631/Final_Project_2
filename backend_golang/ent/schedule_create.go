// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/appointment"
	"FinalProject/ent/schedule"
	"FinalProject/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleCreate is the builder for creating a Schedule entity.
type ScheduleCreate struct {
	config
	mutation *ScheduleMutation
	hooks    []Hook
}

// SetActivity sets the "activity" field.
func (sc *ScheduleCreate) SetActivity(s string) *ScheduleCreate {
	sc.mutation.SetActivity(s)
	return sc
}

// SetDetail sets the "detail" field.
func (sc *ScheduleCreate) SetDetail(s string) *ScheduleCreate {
	sc.mutation.SetDetail(s)
	return sc
}

// SetStatus sets the "status" field.
func (sc *ScheduleCreate) SetStatus(s string) *ScheduleCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetStartTime sets the "startTime" field.
func (sc *ScheduleCreate) SetStartTime(t time.Time) *ScheduleCreate {
	sc.mutation.SetStartTime(t)
	return sc
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableStartTime(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetStartTime(*t)
	}
	return sc
}

// SetEndTime sets the "endTime" field.
func (sc *ScheduleCreate) SetEndTime(t time.Time) *ScheduleCreate {
	sc.mutation.SetEndTime(t)
	return sc
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableEndTime(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetEndTime(*t)
	}
	return sc
}

// AddScheduleAppointmentIDs adds the "schedule_appointment" edge to the Appointment entity by IDs.
func (sc *ScheduleCreate) AddScheduleAppointmentIDs(ids ...int) *ScheduleCreate {
	sc.mutation.AddScheduleAppointmentIDs(ids...)
	return sc
}

// AddScheduleAppointment adds the "schedule_appointment" edges to the Appointment entity.
func (sc *ScheduleCreate) AddScheduleAppointment(a ...*Appointment) *ScheduleCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddScheduleAppointmentIDs(ids...)
}

// SetScheduleDoctorID sets the "schedule_doctor" edge to the User entity by ID.
func (sc *ScheduleCreate) SetScheduleDoctorID(id int) *ScheduleCreate {
	sc.mutation.SetScheduleDoctorID(id)
	return sc
}

// SetNillableScheduleDoctorID sets the "schedule_doctor" edge to the User entity by ID if the given value is not nil.
func (sc *ScheduleCreate) SetNillableScheduleDoctorID(id *int) *ScheduleCreate {
	if id != nil {
		sc = sc.SetScheduleDoctorID(*id)
	}
	return sc
}

// SetScheduleDoctor sets the "schedule_doctor" edge to the User entity.
func (sc *ScheduleCreate) SetScheduleDoctor(u *User) *ScheduleCreate {
	return sc.SetScheduleDoctorID(u.ID)
}

// Mutation returns the ScheduleMutation object of the builder.
func (sc *ScheduleCreate) Mutation() *ScheduleMutation {
	return sc.mutation
}

// Save creates the Schedule in the database.
func (sc *ScheduleCreate) Save(ctx context.Context) (*Schedule, error) {
	var (
		err  error
		node *Schedule
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScheduleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScheduleCreate) SaveX(ctx context.Context) *Schedule {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScheduleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScheduleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScheduleCreate) check() error {
	if _, ok := sc.mutation.Activity(); !ok {
		return &ValidationError{Name: "activity", err: errors.New(`ent: missing required field "Schedule.activity"`)}
	}
	if _, ok := sc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`ent: missing required field "Schedule.detail"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Schedule.status"`)}
	}
	return nil
}

func (sc *ScheduleCreate) sqlSave(ctx context.Context) (*Schedule, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ScheduleCreate) createSpec() (*Schedule, *sqlgraph.CreateSpec) {
	var (
		_node = &Schedule{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: schedule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: schedule.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Activity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldActivity,
		})
		_node.Activity = value
	}
	if value, ok := sc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: schedule.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := sc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: schedule.FieldEndTime,
		})
		_node.EndTime = value
	}
	if nodes := sc.mutation.ScheduleAppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.ScheduleAppointmentTable,
			Columns: []string{schedule.ScheduleAppointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: appointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ScheduleDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedule.ScheduleDoctorTable,
			Columns: []string{schedule.ScheduleDoctorColumn},
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
		_node.user_doctor_schedule = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScheduleCreateBulk is the builder for creating many Schedule entities in bulk.
type ScheduleCreateBulk struct {
	config
	builders []*ScheduleCreate
}

// Save creates the Schedule entities in the database.
func (scb *ScheduleCreateBulk) Save(ctx context.Context) ([]*Schedule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Schedule, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScheduleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScheduleCreateBulk) SaveX(ctx context.Context) []*Schedule {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScheduleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScheduleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}