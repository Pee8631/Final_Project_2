package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Certification holds the schema definition for the Certification entity.
type Certification struct {
	ent.Schema
}

// Fields of the Certification.
func (Certification) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("code").NotEmpty(),
		field.String("diloma").NotEmpty(),
		field.Time("dateOfIssuing").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("dateOfExp").
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.String("Issuer"),

	}
}

// Edges of the Certification.
func (Certification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor_owner", User.Type).
			Ref("doctor_has_certification").
			Unique(),
	}
}
