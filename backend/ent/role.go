// Code generated by entc, DO NOT EDIT.

package ent

import (
	"FinalProject/ent/role"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// RoleUser holds the value of the role_user edge.
	RoleUser []*User `json:"role_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleUserOrErr returns the RoleUser value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RoleUserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.RoleUser, nil
	}
	return nil, &NotLoadedError{edge: "role_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			values[i] = new(sql.NullInt64)
		case role.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		}
	}
	return nil
}

// QueryRoleUser queries the "role_user" edge of the Role entity.
func (r *Role) QueryRoleUser() *UserQuery {
	return (&RoleClient{config: r.config}).QueryRoleUser(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
