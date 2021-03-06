// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/hospital"
	"FinalProject/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HospitalCreate is the builder for creating a Hospital entity.
type HospitalCreate struct {
	config
	mutation *HospitalMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (hc *HospitalCreate) SetName(s string) *HospitalCreate {
	hc.mutation.SetName(s)
	return hc
}

// AddHospitalHasDoctorIDs adds the "hospital_has_doctor" edge to the User entity by IDs.
func (hc *HospitalCreate) AddHospitalHasDoctorIDs(ids ...int) *HospitalCreate {
	hc.mutation.AddHospitalHasDoctorIDs(ids...)
	return hc
}

// AddHospitalHasDoctor adds the "hospital_has_doctor" edges to the User entity.
func (hc *HospitalCreate) AddHospitalHasDoctor(u ...*User) *HospitalCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hc.AddHospitalHasDoctorIDs(ids...)
}

// Mutation returns the HospitalMutation object of the builder.
func (hc *HospitalCreate) Mutation() *HospitalMutation {
	return hc.mutation
}

// Save creates the Hospital in the database.
func (hc *HospitalCreate) Save(ctx context.Context) (*Hospital, error) {
	var (
		err  error
		node *Hospital
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HospitalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HospitalCreate) SaveX(ctx context.Context) *Hospital {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HospitalCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HospitalCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HospitalCreate) check() error {
	if _, ok := hc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Hospital.name"`)}
	}
	if v, ok := hc.mutation.Name(); ok {
		if err := hospital.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Hospital.name": %w`, err)}
		}
	}
	return nil
}

func (hc *HospitalCreate) sqlSave(ctx context.Context) (*Hospital, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HospitalCreate) createSpec() (*Hospital, *sqlgraph.CreateSpec) {
	var (
		_node = &Hospital{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hospital.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hospital.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hospital.FieldName,
		})
		_node.Name = value
	}
	if nodes := hc.mutation.HospitalHasDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   hospital.HospitalHasDoctorTable,
			Columns: []string{hospital.HospitalHasDoctorColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HospitalCreateBulk is the builder for creating many Hospital entities in bulk.
type HospitalCreateBulk struct {
	config
	builders []*HospitalCreate
}

// Save creates the Hospital entities in the database.
func (hcb *HospitalCreateBulk) Save(ctx context.Context) ([]*Hospital, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hospital, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HospitalMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HospitalCreateBulk) SaveX(ctx context.Context) []*Hospital {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HospitalCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HospitalCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
