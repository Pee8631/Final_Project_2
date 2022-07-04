package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chating holds the schema definition for the Chating entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chating.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("Chat_room_name"),
		field.Bool("IsLockChat").Default(false),
	}
}

// Edges of the Chating.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat_user", User.Type),
		edge.To("chat_message", Message.Type),
		edge.To("chat_appointment", Appointment.Type),
	}
}
