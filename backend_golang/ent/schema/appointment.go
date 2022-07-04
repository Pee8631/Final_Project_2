package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Appointmenting holds the schema definition for the Appointmenting entity.
type Appointment struct {
	ent.Schema
}

// Fields of the Appointmenting.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("reasonForAppointment"),
		field.Text("detail"),
		field.Time("startTime").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("endTime").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.String("status"),
		field.Int("DoctorId"),
		field.Int("UserId"),
	}
}

// Edges of the Appointmenting.
func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appointment_schedule", Schedule.Type).
			Ref("schedule_appointment").
			Unique(),
		edge.From("appointment_chat", Chat.Type).
			Ref("chat_appointment").
			Unique(),
	}
}
