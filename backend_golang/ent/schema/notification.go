package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chating holds the schema definition for the Chating entity.
type Notification struct {
	ent.Schema
}

// Fields of the Chating.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("Message"),
		field.Time("CreatedDate").Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Int("RecipientId"),
		field.Int("SenderId"),
		field.Int("AppointmentId"),
	}
}

// Edges of the Chating.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_notification", User.Type).
			Ref("user_send_notification"),
	}
}
