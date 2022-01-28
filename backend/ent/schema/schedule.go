package schema

import (
	"entgo.io/ent"
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
		field.String("status").NotEmpty(),


	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("time_schedule", ScheduleTime.Type),
		edge.From("who_is_the_owner_of_this_schedule", User.Type).
		Ref("doctor_has_schedule").
		Unique(),
	}
}
