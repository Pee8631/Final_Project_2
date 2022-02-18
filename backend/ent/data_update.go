// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/data"
	"FinalProject/ent/predicate"
	"FinalProject/ent/user"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DataUpdate is the builder for updating Data entities.
type DataUpdate struct {
	config
	hooks    []Hook
	mutation *DataMutation
}

// Where appends a list predicates to the DataUpdate builder.
func (du *DataUpdate) Where(ps ...predicate.Data) *DataUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetIdCardNumber sets the "idCardNumber" field.
func (du *DataUpdate) SetIdCardNumber(s string) *DataUpdate {
	du.mutation.SetIdCardNumber(s)
	return du
}

// SetFirstName sets the "firstName" field.
func (du *DataUpdate) SetFirstName(s string) *DataUpdate {
	du.mutation.SetFirstName(s)
	return du
}

// SetLastName sets the "lastName" field.
func (du *DataUpdate) SetLastName(s string) *DataUpdate {
	du.mutation.SetLastName(s)
	return du
}

// SetGender sets the "gender" field.
func (du *DataUpdate) SetGender(i int) *DataUpdate {
	du.mutation.ResetGender()
	du.mutation.SetGender(i)
	return du
}

// AddGender adds i to the "gender" field.
func (du *DataUpdate) AddGender(i int) *DataUpdate {
	du.mutation.AddGender(i)
	return du
}

// SetBrithDate sets the "brithDate" field.
func (du *DataUpdate) SetBrithDate(t time.Time) *DataUpdate {
	du.mutation.SetBrithDate(t)
	return du
}

// SetNillableBrithDate sets the "brithDate" field if the given value is not nil.
func (du *DataUpdate) SetNillableBrithDate(t *time.Time) *DataUpdate {
	if t != nil {
		du.SetBrithDate(*t)
	}
	return du
}

// ClearBrithDate clears the value of the "brithDate" field.
func (du *DataUpdate) ClearBrithDate() *DataUpdate {
	du.mutation.ClearBrithDate()
	return du
}

// SetBloodGroup sets the "bloodGroup" field.
func (du *DataUpdate) SetBloodGroup(s string) *DataUpdate {
	du.mutation.SetBloodGroup(s)
	return du
}

// SetAddress sets the "address" field.
func (du *DataUpdate) SetAddress(s string) *DataUpdate {
	du.mutation.SetAddress(s)
	return du
}

// SetWhoIsTheOwnerOfThisDataID sets the "who_is_the_owner_of_this_data" edge to the User entity by ID.
func (du *DataUpdate) SetWhoIsTheOwnerOfThisDataID(id int) *DataUpdate {
	du.mutation.SetWhoIsTheOwnerOfThisDataID(id)
	return du
}

// SetNillableWhoIsTheOwnerOfThisDataID sets the "who_is_the_owner_of_this_data" edge to the User entity by ID if the given value is not nil.
func (du *DataUpdate) SetNillableWhoIsTheOwnerOfThisDataID(id *int) *DataUpdate {
	if id != nil {
		du = du.SetWhoIsTheOwnerOfThisDataID(*id)
	}
	return du
}

// SetWhoIsTheOwnerOfThisData sets the "who_is_the_owner_of_this_data" edge to the User entity.
func (du *DataUpdate) SetWhoIsTheOwnerOfThisData(u *User) *DataUpdate {
	return du.SetWhoIsTheOwnerOfThisDataID(u.ID)
}

// Mutation returns the DataMutation object of the builder.
func (du *DataUpdate) Mutation() *DataMutation {
	return du.mutation
}

// ClearWhoIsTheOwnerOfThisData clears the "who_is_the_owner_of_this_data" edge to the User entity.
func (du *DataUpdate) ClearWhoIsTheOwnerOfThisData() *DataUpdate {
	du.mutation.ClearWhoIsTheOwnerOfThisData()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DataUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DataUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DataUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   data.Table,
			Columns: data.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: data.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.IdCardNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldIdCardNumber,
		})
	}
	if value, ok := du.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldFirstName,
		})
	}
	if value, ok := du.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldLastName,
		})
	}
	if value, ok := du.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: data.FieldGender,
		})
	}
	if value, ok := du.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: data.FieldGender,
		})
	}
	if value, ok := du.mutation.BrithDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: data.FieldBrithDate,
		})
	}
	if du.mutation.BrithDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: data.FieldBrithDate,
		})
	}
	if value, ok := du.mutation.BloodGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldBloodGroup,
		})
	}
	if value, ok := du.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldAddress,
		})
	}
	if du.mutation.WhoIsTheOwnerOfThisDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   data.WhoIsTheOwnerOfThisDataTable,
			Columns: []string{data.WhoIsTheOwnerOfThisDataColumn},
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
	if nodes := du.mutation.WhoIsTheOwnerOfThisDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   data.WhoIsTheOwnerOfThisDataTable,
			Columns: []string{data.WhoIsTheOwnerOfThisDataColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{data.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DataUpdateOne is the builder for updating a single Data entity.
type DataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DataMutation
}

// SetIdCardNumber sets the "idCardNumber" field.
func (duo *DataUpdateOne) SetIdCardNumber(s string) *DataUpdateOne {
	duo.mutation.SetIdCardNumber(s)
	return duo
}

// SetFirstName sets the "firstName" field.
func (duo *DataUpdateOne) SetFirstName(s string) *DataUpdateOne {
	duo.mutation.SetFirstName(s)
	return duo
}

// SetLastName sets the "lastName" field.
func (duo *DataUpdateOne) SetLastName(s string) *DataUpdateOne {
	duo.mutation.SetLastName(s)
	return duo
}

// SetGender sets the "gender" field.
func (duo *DataUpdateOne) SetGender(i int) *DataUpdateOne {
	duo.mutation.ResetGender()
	duo.mutation.SetGender(i)
	return duo
}

// AddGender adds i to the "gender" field.
func (duo *DataUpdateOne) AddGender(i int) *DataUpdateOne {
	duo.mutation.AddGender(i)
	return duo
}

// SetBrithDate sets the "brithDate" field.
func (duo *DataUpdateOne) SetBrithDate(t time.Time) *DataUpdateOne {
	duo.mutation.SetBrithDate(t)
	return duo
}

// SetNillableBrithDate sets the "brithDate" field if the given value is not nil.
func (duo *DataUpdateOne) SetNillableBrithDate(t *time.Time) *DataUpdateOne {
	if t != nil {
		duo.SetBrithDate(*t)
	}
	return duo
}

// ClearBrithDate clears the value of the "brithDate" field.
func (duo *DataUpdateOne) ClearBrithDate() *DataUpdateOne {
	duo.mutation.ClearBrithDate()
	return duo
}

// SetBloodGroup sets the "bloodGroup" field.
func (duo *DataUpdateOne) SetBloodGroup(s string) *DataUpdateOne {
	duo.mutation.SetBloodGroup(s)
	return duo
}

// SetAddress sets the "address" field.
func (duo *DataUpdateOne) SetAddress(s string) *DataUpdateOne {
	duo.mutation.SetAddress(s)
	return duo
}

// SetWhoIsTheOwnerOfThisDataID sets the "who_is_the_owner_of_this_data" edge to the User entity by ID.
func (duo *DataUpdateOne) SetWhoIsTheOwnerOfThisDataID(id int) *DataUpdateOne {
	duo.mutation.SetWhoIsTheOwnerOfThisDataID(id)
	return duo
}

// SetNillableWhoIsTheOwnerOfThisDataID sets the "who_is_the_owner_of_this_data" edge to the User entity by ID if the given value is not nil.
func (duo *DataUpdateOne) SetNillableWhoIsTheOwnerOfThisDataID(id *int) *DataUpdateOne {
	if id != nil {
		duo = duo.SetWhoIsTheOwnerOfThisDataID(*id)
	}
	return duo
}

// SetWhoIsTheOwnerOfThisData sets the "who_is_the_owner_of_this_data" edge to the User entity.
func (duo *DataUpdateOne) SetWhoIsTheOwnerOfThisData(u *User) *DataUpdateOne {
	return duo.SetWhoIsTheOwnerOfThisDataID(u.ID)
}

// Mutation returns the DataMutation object of the builder.
func (duo *DataUpdateOne) Mutation() *DataMutation {
	return duo.mutation
}

// ClearWhoIsTheOwnerOfThisData clears the "who_is_the_owner_of_this_data" edge to the User entity.
func (duo *DataUpdateOne) ClearWhoIsTheOwnerOfThisData() *DataUpdateOne {
	duo.mutation.ClearWhoIsTheOwnerOfThisData()
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DataUpdateOne) Select(field string, fields ...string) *DataUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Data entity.
func (duo *DataUpdateOne) Save(ctx context.Context) (*Data, error) {
	var (
		err  error
		node *Data
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DataUpdateOne) SaveX(ctx context.Context) *Data {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DataUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DataUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DataUpdateOne) sqlSave(ctx context.Context) (_node *Data, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   data.Table,
			Columns: data.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: data.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Data.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, data.FieldID)
		for _, f := range fields {
			if !data.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != data.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.IdCardNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldIdCardNumber,
		})
	}
	if value, ok := duo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldFirstName,
		})
	}
	if value, ok := duo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldLastName,
		})
	}
	if value, ok := duo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: data.FieldGender,
		})
	}
	if value, ok := duo.mutation.AddedGender(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: data.FieldGender,
		})
	}
	if value, ok := duo.mutation.BrithDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: data.FieldBrithDate,
		})
	}
	if duo.mutation.BrithDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: data.FieldBrithDate,
		})
	}
	if value, ok := duo.mutation.BloodGroup(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldBloodGroup,
		})
	}
	if value, ok := duo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: data.FieldAddress,
		})
	}
	if duo.mutation.WhoIsTheOwnerOfThisDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   data.WhoIsTheOwnerOfThisDataTable,
			Columns: []string{data.WhoIsTheOwnerOfThisDataColumn},
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
	if nodes := duo.mutation.WhoIsTheOwnerOfThisDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   data.WhoIsTheOwnerOfThisDataTable,
			Columns: []string{data.WhoIsTheOwnerOfThisDataColumn},
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
	_node = &Data{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{data.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}