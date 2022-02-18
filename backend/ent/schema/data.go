package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Data holds the schema definition for the Data entity.
type Data struct {
	ent.Schema
}

// Fields of the Data.
func (Data) Fields() []ent.Field {
	return []ent.Field{
		field.String("idCardNumber").Unique(),
		field.String("firstName"),
		field.String("lastName"),
		field.Int("gender"),
		field.Time("brithDate").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.String("bloodGroup"),
		field.Text("address"),
	}
}

// Edges of the Data.
func (Data) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("who_is_the_owner_of_this_data", User.Type).
			Ref("user_has_data").
			Unique(),
	}
}
