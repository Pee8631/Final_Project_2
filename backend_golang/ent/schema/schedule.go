package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("activity"),
		field.Text("detail"),
		field.String("status"),
		field.Time("startTime").Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("endTime").Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("schedule_appointment", Appointment.Type),
		edge.From("schedule_doctor", User.Type).
			Ref("doctor_schedule").Unique(),
			
	}
}
