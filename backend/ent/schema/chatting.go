package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chating holds the schema definition for the Chating entity.
type Chatting struct {
	ent.Schema
}

// Fields of the Chating.
func (Chatting) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("message"),
		field.Time("dateTime").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Chating.
func (Chatting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chatting_with_whom", User.Type).
			Ref("user_chatting_with_whom").
			Unique(),
		edge.From("whose_is_this_msg", User.Type).
			Ref("who_is_owner_this_msg").
			Unique(),
	}
}
