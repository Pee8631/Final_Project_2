package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.String("username").Unique().NotEmpty(),
		field.String("password"),
		field.String("email"),
		field.String("telephone"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor_has_certification", Certification.Type),
		edge.To("user_chatting_with_whom", Chatting.Type),
		edge.To("who_is_owner_this_msg", Chatting.Type),
		edge.To("user_has_data", Data.Type),
		edge.To("doctor_has_schedule", Schedule.Type),
		edge.To("user_have_telecoms", Telecom.Type),
		edge.To("doctor_record_treatment", Treatment.Type),
		edge.To("user_have_treatment", Treatment.Type),
		edge.From("has_department", Department.Type).
			Ref("department_has_doctor").
			Unique(),
		edge.From("from_hospital", Hospital.Type).
			Ref("hospital_has_doctor").
			Unique(),
		edge.From("user_have_disease", Disease.Type).
			Ref("disease_user"),
		edge.From("user_have_role", Role.Type).
			Ref("role_user"),

	}
}
