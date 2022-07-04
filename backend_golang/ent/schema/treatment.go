package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Treatment holds the schema definition for the Treatment entity.
type Treatment struct {
	ent.Schema
}

// Fields of the Treatment.
func (Treatment) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.Text("treatmentRecord"),
		field.Time("dateTime").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Float("takeTime"),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("treatment_was_recorded_by_doctor", User.Type).
		Ref("doctor_record_treatment").
		Unique(),
		edge.From("user_is_the_treatment_of_record", User.Type).
		Ref("user_have_treatment").
		Unique(),
	}
}
