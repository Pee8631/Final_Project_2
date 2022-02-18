package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ScheduleTime holds the schema definition for the ScheduleTime entity.
type ScheduleTime struct {
	ent.Schema
}

// Fields of the ScheduleTime.
func (ScheduleTime) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.Time("startTime").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("stopTime").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the ScheduleTime.
func (ScheduleTime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("what_time_is_the_schedule", Schedule.Type).
		Ref("time_schedule").
		Unique(),
	}
}
