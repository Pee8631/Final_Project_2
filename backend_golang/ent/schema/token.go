package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Token struct {
	ent.Schema
}

// Fields of the User.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("AuthToken"),
		field.Time("GeneratedAt").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("ExpiresAt").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the User.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("authentication_token", User.Type).
			Ref("user_have_token").
			Unique(),
	}
}
