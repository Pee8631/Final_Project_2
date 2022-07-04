// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/pinfo"
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

// PInfoUpdate is the builder for updating PInfo entities.
type PInfoUpdate struct {
	config
	hooks    []Hook
	mutation *PInfoMutation
}

// Where appends a list predicates to the PInfoUpdate builder.
func (pu *PInfoUpdate) Where(ps ...predicate.PInfo) *PInfoUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetIdCardNumber sets the "idCardNumber" field.
func (pu *PInfoUpdate) SetIdCardNumber(s string) *PInfoUpdate {
	pu.mutation.SetIdCardNumber(s)
	return pu
}

// SetFirstName sets the "firstName" field.
func (pu *PInfoUpdate) SetFirstName(s string) *PInfoUpdate {
	pu.mutation.SetFirstName(s)
	return pu
}

// SetLastName sets the "lastName" field.
func (pu *PInfoUpdate) SetLastName(s string) *PInfoUpdate {
	pu.mutation.SetLastName(s)
	return pu
}

// SetGender sets the "gender" field.
func (pu *PInfoUpdate) SetGender(i int) *PInfoUpdate {
	pu.mutation.ResetGender()
	pu.mutation.SetGender(i)
	return pu
}

// AddGender adds i to the "gender" field.
func (pu *PInfoUpdate) AddGender(i int) *PInfoUpdate {
	pu.mutation.AddGender(i)
	return pu
}

// SetBrithDate sets the "brithDate" field.
func (pu *PInfoUpdate) SetBrithDate(t time.Time) *PInfoUpdate {
	pu.mutation.SetBrithDate(t)
	return pu
}

// SetNillableBrithDate sets the "brithDate" field if the given value is not nil.
func (pu *PInfoUpdate) SetNillableBrithDate(t *time.Time) *PInfoUpdate {
	if t != nil {
		pu.SetBrithDate(*t)
	}
	return pu
}

// ClearBrithDate clears the value of the "brithDate" field.
func (pu *PInfoUpdate) ClearBrithDate() *PInfoUpdate {
	pu.mutation.ClearBrithDate()
	return pu
}

// SetBloodGroup sets the "bloodGroup" field.
func (pu *PInfoUpdate) SetBloodGroup(s string) *PInfoUpdate {
	pu.mutation.SetBloodGroup(s)
	return pu
}

// SetAddress sets the "address" field.
func (pu *PInfoUpdate) SetAddress(s string) *PInfoUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID.
func (pu *PInfoUpdate) SetWhoIsTheOwnerOfThisPInfoID(id int) *PInfoUpdate {
	pu.mutation.SetWhoIsTheOwnerOfThisPInfoID(id)
	return pu
}

// SetNillableWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID if the given value is not nil.
func (pu *PInfoUpdate) SetNillableWhoIsTheOwnerOfThisPInfoID(id *int) *PInfoUpdate {
	if id != nil {
		pu = pu.SetWhoIsTheOwnerOfThisPInfoID(*id)
	}
	return pu
}

// SetWhoIsTheOwnerOfThisPInfo sets the "who_is_the_owner_of_this_PInfo" edge to the User entity.
func (pu *PInfoUpdate) SetWhoIsTheOwnerOfThisPInfo(u *User) *PInfoUpdate {
	return pu.SetWhoIsTheOwnerOfThisPInfoID(u.ID)
}

// Mutation returns the PInfoMutation object of the builder.
func (pu *PInfoUpdate) Mutation() *PInfoMutation {
	return pu.mutation
}

// ClearWhoIsTheOwnerOfThisPInfo clears the "who_is_the_owner_of_this_PInfo" edge to the User entity.
func (pu *PInfoUpdate) ClearWhoIsTheOwnerOfThisPInfo() *PInfoUpdate {
	pu.mutation.ClearWhoIsTheOwnerOfThisPInfo()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PInfoUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PInfoUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pinfo.Table,
			Columns: pinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pinfo.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.IdCardNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldIdCardNumber,
		})
	}
	if value, ok := pu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldFirstName,
		})
	}
	if value, ok := pu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldLastName,
		})
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pinfo.FieldGender,
		})
	}
	if value, ok := pu.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pinfo.FieldGender,
		})
	}
	if value, ok := pu.mutation.BrithDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pinfo.FieldBrithDate,
		})
	}
	if pu.mutation.BrithDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pinfo.FieldBrithDate,
		})
	}
	if value, ok := pu.mutation.BloodGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldBloodGroup,
		})
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldAddress,
		})
	}
	if pu.mutation.WhoIsTheOwnerOfThisPInfoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.WhoIsTheOwnerOfThisPInfoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PInfoUpdateOne is the builder for updating a single PInfo entity.
type PInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PInfoMutation
}

// SetIdCardNumber sets the "idCardNumber" field.
func (puo *PInfoUpdateOne) SetIdCardNumber(s string) *PInfoUpdateOne {
	puo.mutation.SetIdCardNumber(s)
	return puo
}

// SetFirstName sets the "firstName" field.
func (puo *PInfoUpdateOne) SetFirstName(s string) *PInfoUpdateOne {
	puo.mutation.SetFirstName(s)
	return puo
}

// SetLastName sets the "lastName" field.
func (puo *PInfoUpdateOne) SetLastName(s string) *PInfoUpdateOne {
	puo.mutation.SetLastName(s)
	return puo
}

// SetGender sets the "gender" field.
func (puo *PInfoUpdateOne) SetGender(i int) *PInfoUpdateOne {
	puo.mutation.ResetGender()
	puo.mutation.SetGender(i)
	return puo
}

// AddGender adds i to the "gender" field.
func (puo *PInfoUpdateOne) AddGender(i int) *PInfoUpdateOne {
	puo.mutation.AddGender(i)
	return puo
}

// SetBrithDate sets the "brithDate" field.
func (puo *PInfoUpdateOne) SetBrithDate(t time.Time) *PInfoUpdateOne {
	puo.mutation.SetBrithDate(t)
	return puo
}

// SetNillableBrithDate sets the "brithDate" field if the given value is not nil.
func (puo *PInfoUpdateOne) SetNillableBrithDate(t *time.Time) *PInfoUpdateOne {
	if t != nil {
		puo.SetBrithDate(*t)
	}
	return puo
}

// ClearBrithDate clears the value of the "brithDate" field.
func (puo *PInfoUpdateOne) ClearBrithDate() *PInfoUpdateOne {
	puo.mutation.ClearBrithDate()
	return puo
}

// SetBloodGroup sets the "bloodGroup" field.
func (puo *PInfoUpdateOne) SetBloodGroup(s string) *PInfoUpdateOne {
	puo.mutation.SetBloodGroup(s)
	return puo
}

// SetAddress sets the "address" field.
func (puo *PInfoUpdateOne) SetAddress(s string) *PInfoUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID.
func (puo *PInfoUpdateOne) SetWhoIsTheOwnerOfThisPInfoID(id int) *PInfoUpdateOne {
	puo.mutation.SetWhoIsTheOwnerOfThisPInfoID(id)
	return puo
}

// SetNillableWhoIsTheOwnerOfThisPInfoID sets the "who_is_the_owner_of_this_PInfo" edge to the User entity by ID if the given value is not nil.
func (puo *PInfoUpdateOne) SetNillableWhoIsTheOwnerOfThisPInfoID(id *int) *PInfoUpdateOne {
	if id != nil {
		puo = puo.SetWhoIsTheOwnerOfThisPInfoID(*id)
	}
	return puo
}

// SetWhoIsTheOwnerOfThisPInfo sets the "who_is_the_owner_of_this_PInfo" edge to the User entity.
func (puo *PInfoUpdateOne) SetWhoIsTheOwnerOfThisPInfo(u *User) *PInfoUpdateOne {
	return puo.SetWhoIsTheOwnerOfThisPInfoID(u.ID)
}

// Mutation returns the PInfoMutation object of the builder.
func (puo *PInfoUpdateOne) Mutation() *PInfoMutation {
	return puo.mutation
}

// ClearWhoIsTheOwnerOfThisPInfo clears the "who_is_the_owner_of_this_PInfo" edge to the User entity.
func (puo *PInfoUpdateOne) ClearWhoIsTheOwnerOfThisPInfo() *PInfoUpdateOne {
	puo.mutation.ClearWhoIsTheOwnerOfThisPInfo()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PInfoUpdateOne) Select(field string, fields ...string) *PInfoUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated PInfo entity.
func (puo *PInfoUpdateOne) Save(ctx context.Context) (*PInfo, error) {
	var (
		err  error
		node *PInfo
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PInfoUpdateOne) SaveX(ctx context.Context) *PInfo {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PInfoUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PInfoUpdateOne) sqlSave(ctx context.Context) (_node *PInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pinfo.Table,
			Columns: pinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pinfo.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pinfo.FieldID)
		for _, f := range fields {
			if !pinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.IdCardNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldIdCardNumber,
		})
	}
	if value, ok := puo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldFirstName,
		})
	}
	if value, ok := puo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldLastName,
		})
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pinfo.FieldGender,
		})
	}
	if value, ok := puo.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pinfo.FieldGender,
		})
	}
	if value, ok := puo.mutation.BrithDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pinfo.FieldBrithDate,
		})
	}
	if puo.mutation.BrithDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pinfo.FieldBrithDate,
		})
	}
	if value, ok := puo.mutation.BloodGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldBloodGroup,
		})
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pinfo.FieldAddress,
		})
	}
	if puo.mutation.WhoIsTheOwnerOfThisPInfoCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.WhoIsTheOwnerOfThisPInfoIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PInfo{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
