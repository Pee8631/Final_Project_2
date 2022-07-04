// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/pinfo"
	"FinalProject/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PInfoCreate is the builder for creating a PInfo entity.
type PInfoCreate struct {
	config
	mutation *PInfoMutation
	hooks    []Hook
}

// SetIdCardNumber sets the "idCardNumber" field.
func (pc *PInfoCreate) SetIdCardNumber(s string) *PInfoCreate {
	pc.mutation.SetIdCardNumber(s)
	return pc
}

// SetFirstName sets the "firstName" field.
func (pc *PInfoCreate) SetFirstName(s string) *PInfoCreate {
	pc.mutation.SetFirstName(s)
	return pc
}

// SetLastName sets the "lastName" field.
func (pc *PInfoCreate) SetLastName(s string) *PInfoCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetGender sets the "gender" field.
func (pc *PInfoCreate) SetGender(i int) *PInfoCreate {
	pc.mutation.SetGender(i)
	return pc
}

// SetBrithDate sets the "brithDate" field.
func (pc *PInfoCreate) SetBrithDate(t time.Time) *PInfoCreate {
	pc.mutation.SetBrithDate(t)
	return pc
}

// SetNillableBrithDate sets the "brithDate" field if the given value is not nil.
func (pc *PInfoCreate) SetNillableBrithDate(t *time.Time) *PInfoCreate {
	if t != nil {
		pc.SetBrithDate(*t)
	}
	return pc
}

// SetBloodGroup sets the "bloodGroup" field.
func (pc *PInfoCreate) SetBloodGroup(s string) *PInfoCreate {
	pc.mutation.SetBloodGroup(s)
	return pc
}

// SetAddress sets the "address" field.
func (pc *PInfoCreate) SetAddress(s string) *PInfoCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID.
func (pc *PInfoCreate) SetWhoIsTheOwnerOfThisPInfoID(id int) *PInfoCreate {
	pc.mutation.SetWhoIsTheOwnerOfThisPInfoID(id)
	return pc
}

// SetNillableWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID if the given value is not nil.
func (pc *PInfoCreate) SetNillableWhoIsTheOwnerOfThisPInfoID(id *int) *PInfoCreate {
	if id != nil {
		pc = pc.SetWhoIsTheOwnerOfThisPInfoID(*id)
	}
	return pc
}

// SetWhoIsTheOwnerOfThisPInfo sets the "who_is_the_owner_of_this_PInfo" edge to the User entity.
func (pc *PInfoCreate) SetWhoIsTheOwnerOfThisPInfo(u *User) *PInfoCreate {
	return pc.SetWhoIsTheOwnerOfThisPInfoID(u.ID)
}

// Mutation returns the PInfoMutation object of the builder.
func (pc *PInfoCreate) Mutation() *PInfoMutation {
	return pc.mutation
}

// Save creates the PInfo in the database.
func (pc *PInfoCreate) Save(ctx context.Context) (*PInfo, error) {
	var (
		err  error
		node *PInfo
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PInfoCreate) SaveX(ctx context.Context) *PInfo {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PInfoCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PInfoCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PInfoCreate) check() error {
	if _, ok := pc.mutation.IdCardNumber(); !ok {
		return &ValidationError{Name: "idCardNumber", err: errors.New(`ent: missing required field "PInfo.idCardNumber"`)}
	}
	if _, ok := pc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "firstName", err: errors.New(`ent: missing required field "PInfo.firstName"`)}
	}
	if _, ok := pc.mutation.LastName(); !ok {
		return &ValidationError{Name: "lastName", err: errors.New(`ent: missing required field "PInfo.lastName"`)}
	}
	if _, ok := pc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "PInfo.gender"`)}
	}
	if _, ok := pc.mutation.BloodGroup(); !ok {
		return &ValidationError{Name: "bloodGroup", err: errors.New(`ent: missing required field "PInfo.bloodGroup"`)}
	}
	if _, ok := pc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "PInfo.address"`)}
	}
	return nil
}

func (pc *PInfoCreate) sqlSave(ctx context.Context) (*PInfo, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PInfoCreate) createSpec() (*PInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &PInfo{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pinfo.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.IdCardNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldIdCardNumber,
		})
		_node.IdCardNumber = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pinfo.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := pc.mutation.BrithDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pinfo.FieldBrithDate,
		})
		_node.BrithDate = value
	}
	if value, ok := pc.mutation.BloodGroup(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldBloodGroup,
		})
		_node.BloodGroup = value
	}
	if value, ok := pc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldAddress,
		})
		_node.Address = value
	}
	if nodes := pc.mutation.WhoIsTheOwnerOfThisPInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pinfo.WhoIsTheOwnerOfThisPInfoTable,
			Columns: []string{pinfo.WhoIsTheOwnerOfThisPInfoColumn},
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
		_node.user_user_has_p_info = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PInfoCreateBulk is the builder for creating many PInfo entities in bulk.
type PInfoCreateBulk struct {
	config
	builders []*PInfoCreate
}

// Save creates the PInfo entities in the database.
func (pcb *PInfoCreateBulk) Save(ctx context.Context) ([]*PInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*PInfo, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PInfoCreateBulk) SaveX(ctx context.Context) []*PInfo {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PInfoCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
