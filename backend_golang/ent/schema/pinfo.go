package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PInfo holds the schema definition for the PInfo entity.
type PInfo struct {
	ent.Schema
}

// Fields of the PInfo.
func (PInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("profile"),
		field.String("idCardNumber").Unique(),
		field.String("prefix"),
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
		field.Text("about"),
	}
}

// Edges of the PInfo.
func (PInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("who_is_the_owner_of_this_PInfo", User.Type).
			Ref("user_has_PInfo").
			Unique(),
	}
}
