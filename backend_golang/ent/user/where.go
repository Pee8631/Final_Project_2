// Code generated by entc, DO NOT EDIT.

package user

import (
	"FinalProject/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// HasDoctorHasCertification applies the HasEdge predicate on the "doctor_has_certification" edge.
func HasDoctorHasCertification() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorHasCertificationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorHasCertificationTable, DoctorHasCertificationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorHasCertificationWith applies the HasEdge predicate on the "doctor_has_certification" edge with a given conditions (other predicates).
func HasDoctorHasCertificationWith(preds ...predicate.Certification) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorHasCertificationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorHasCertificationTable, DoctorHasCertificationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHasPInfo applies the HasEdge predicate on the "user_has_PInfo" edge.
func HasUserHasPInfo() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHasPInfoTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHasPInfoTable, UserHasPInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHasPInfoWith applies the HasEdge predicate on the "user_has_PInfo" edge with a given conditions (other predicates).
func HasUserHasPInfoWith(preds ...predicate.PInfo) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHasPInfoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHasPInfoTable, UserHasPInfoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDoctorSchedule applies the HasEdge predicate on the "doctor_schedule" edge.
func HasDoctorSchedule() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorScheduleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorScheduleTable, DoctorScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorScheduleWith applies the HasEdge predicate on the "doctor_schedule" edge with a given conditions (other predicates).
func HasDoctorScheduleWith(preds ...predicate.Schedule) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorScheduleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorScheduleTable, DoctorScheduleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHaveTelecoms applies the HasEdge predicate on the "user_have_telecoms" edge.
func HasUserHaveTelecoms() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTelecomsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTelecomsTable, UserHaveTelecomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHaveTelecomsWith applies the HasEdge predicate on the "user_have_telecoms" edge with a given conditions (other predicates).
func HasUserHaveTelecomsWith(preds ...predicate.Telecom) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTelecomsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTelecomsTable, UserHaveTelecomsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDoctorRecordTreatment applies the HasEdge predicate on the "doctor_record_treatment" edge.
func HasDoctorRecordTreatment() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorRecordTreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorRecordTreatmentTable, DoctorRecordTreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorRecordTreatmentWith applies the HasEdge predicate on the "doctor_record_treatment" edge with a given conditions (other predicates).
func HasDoctorRecordTreatmentWith(preds ...predicate.Treatment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorRecordTreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorRecordTreatmentTable, DoctorRecordTreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHaveTreatment applies the HasEdge predicate on the "user_have_treatment" edge.
func HasUserHaveTreatment() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTreatmentTable, UserHaveTreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHaveTreatmentWith applies the HasEdge predicate on the "user_have_treatment" edge with a given conditions (other predicates).
func HasUserHaveTreatmentWith(preds ...predicate.Treatment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTreatmentTable, UserHaveTreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHaveToken applies the HasEdge predicate on the "user_have_token" edge.
func HasUserHaveToken() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTokenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTokenTable, UserHaveTokenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHaveTokenWith applies the HasEdge predicate on the "user_have_token" edge with a given conditions (other predicates).
func HasUserHaveTokenWith(preds ...predicate.Token) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveTokenInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserHaveTokenTable, UserHaveTokenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserSendMessage applies the HasEdge predicate on the "user_send_message" edge.
func HasUserSendMessage() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserSendMessageTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserSendMessageTable, UserSendMessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserSendMessageWith applies the HasEdge predicate on the "user_send_message" edge with a given conditions (other predicates).
func HasUserSendMessageWith(preds ...predicate.Message) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserSendMessageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserSendMessageTable, UserSendMessageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserSendNotification applies the HasEdge predicate on the "user_send_notification" edge.
func HasUserSendNotification() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserSendNotificationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserSendNotificationTable, UserSendNotificationPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserSendNotificationWith applies the HasEdge predicate on the "user_send_notification" edge with a given conditions (other predicates).
func HasUserSendNotificationWith(preds ...predicate.Notification) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserSendNotificationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserSendNotificationTable, UserSendNotificationPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHasDepartment applies the HasEdge predicate on the "has_department" edge.
func HasHasDepartment() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HasDepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasDepartmentTable, HasDepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHasDepartmentWith applies the HasEdge predicate on the "has_department" edge with a given conditions (other predicates).
func HasHasDepartmentWith(preds ...predicate.Department) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HasDepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HasDepartmentTable, HasDepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFromHospital applies the HasEdge predicate on the "from_hospital" edge.
func HasFromHospital() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromHospitalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromHospitalTable, FromHospitalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromHospitalWith applies the HasEdge predicate on the "from_hospital" edge with a given conditions (other predicates).
func HasFromHospitalWith(preds ...predicate.Hospital) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromHospitalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromHospitalTable, FromHospitalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHaveDisease applies the HasEdge predicate on the "user_have_disease" edge.
func HasUserHaveDisease() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveDiseaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserHaveDiseaseTable, UserHaveDiseasePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHaveDiseaseWith applies the HasEdge predicate on the "user_have_disease" edge with a given conditions (other predicates).
func HasUserHaveDiseaseWith(preds ...predicate.Disease) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveDiseaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserHaveDiseaseTable, UserHaveDiseasePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserHaveRole applies the HasEdge predicate on the "user_have_role" edge.
func HasUserHaveRole() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveRoleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserHaveRoleTable, UserHaveRolePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserHaveRoleWith applies the HasEdge predicate on the "user_have_role" edge with a given conditions (other predicates).
func HasUserHaveRoleWith(preds ...predicate.Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserHaveRoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserHaveRoleTable, UserHaveRolePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWhoIsInThisChat applies the HasEdge predicate on the "who_is_in_this_chat" edge.
func HasWhoIsInThisChat() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoIsInThisChatTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WhoIsInThisChatTable, WhoIsInThisChatPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhoIsInThisChatWith applies the HasEdge predicate on the "who_is_in_this_chat" edge with a given conditions (other predicates).
func HasWhoIsInThisChatWith(preds ...predicate.Chat) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoIsInThisChatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WhoIsInThisChatTable, WhoIsInThisChatPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
