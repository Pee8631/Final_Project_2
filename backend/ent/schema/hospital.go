package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Hospital.
func (Hospital) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hospital_has_doctor", User.Type),
	}
}
