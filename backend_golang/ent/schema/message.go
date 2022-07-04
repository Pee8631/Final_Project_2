package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chating holds the schema definition for the Chating entity.
type Message struct {
	ent.Schema
}

// Fields of the Chating.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("message_text"),
		field.Time("sent_dateTime").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the Chating.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("What_messages_are_in_this_chat", Chat.Type).
			Ref("chat_message").
			Unique(),
		edge.From("Who_send_messages", User.Type).
			Ref("user_send_message").
			Unique(),
	}
}
