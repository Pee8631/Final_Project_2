package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Telecom holds the schema definition for the Telecom entity.
type Telecom struct {
	ent.Schema
}

// Fields of the Telecom.
func (Telecom) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("username").NotEmpty(),
		field.String("platform").NotEmpty(),
	}
}

// Edges of the Telecom.
func (Telecom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("who_is_the_owner_of_this_telecom", User.Type).
		Ref("user_have_telecoms").
		Unique(),
	}
}
