// Code generated by entc, DO NOT EDIT.

package appointment

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ReasonForAppointment applies equality check predicate on the "reasonForAppointment" field. It's identical to ReasonForAppointmentEQ.
func ReasonForAppointment(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReasonForAppointment), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// DoctorId applies equality check predicate on the "DoctorId" field. It's identical to DoctorIdEQ.
func DoctorId(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorId), v))
	})
}

// UserId applies equality check predicate on the "UserId" field. It's identical to UserIdEQ.
func UserId(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// ReasonForAppointmentEQ applies the EQ predicate on the "reasonForAppointment" field.
func ReasonForAppointmentEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentNEQ applies the NEQ predicate on the "reasonForAppointment" field.
func ReasonForAppointmentNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentIn applies the In predicate on the "reasonForAppointment" field.
func ReasonForAppointmentIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReasonForAppointment), v...))
	})
}

// ReasonForAppointmentNotIn applies the NotIn predicate on the "reasonForAppointment" field.
func ReasonForAppointmentNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReasonForAppointment), v...))
	})
}

// ReasonForAppointmentGT applies the GT predicate on the "reasonForAppointment" field.
func ReasonForAppointmentGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentGTE applies the GTE predicate on the "reasonForAppointment" field.
func ReasonForAppointmentGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentLT applies the LT predicate on the "reasonForAppointment" field.
func ReasonForAppointmentLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentLTE applies the LTE predicate on the "reasonForAppointment" field.
func ReasonForAppointmentLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentContains applies the Contains predicate on the "reasonForAppointment" field.
func ReasonForAppointmentContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentHasPrefix applies the HasPrefix predicate on the "reasonForAppointment" field.
func ReasonForAppointmentHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentHasSuffix applies the HasSuffix predicate on the "reasonForAppointment" field.
func ReasonForAppointmentHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentEqualFold applies the EqualFold predicate on the "reasonForAppointment" field.
func ReasonForAppointmentEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReasonForAppointment), v))
	})
}

// ReasonForAppointmentContainsFold applies the ContainsFold predicate on the "reasonForAppointment" field.
func ReasonForAppointmentContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReasonForAppointment), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...time.Time) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// DoctorIdEQ applies the EQ predicate on the "DoctorId" field.
func DoctorIdEQ(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorId), v))
	})
}

// DoctorIdNEQ applies the NEQ predicate on the "DoctorId" field.
func DoctorIdNEQ(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorId), v))
	})
}

// DoctorIdIn applies the In predicate on the "DoctorId" field.
func DoctorIdIn(vs ...int) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorId), v...))
	})
}

// DoctorIdNotIn applies the NotIn predicate on the "DoctorId" field.
func DoctorIdNotIn(vs ...int) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorId), v...))
	})
}

// DoctorIdGT applies the GT predicate on the "DoctorId" field.
func DoctorIdGT(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorId), v))
	})
}

// DoctorIdGTE applies the GTE predicate on the "DoctorId" field.
func DoctorIdGTE(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorId), v))
	})
}

// DoctorIdLT applies the LT predicate on the "DoctorId" field.
func DoctorIdLT(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorId), v))
	})
}

// DoctorIdLTE applies the LTE predicate on the "DoctorId" field.
func DoctorIdLTE(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorId), v))
	})
}

// UserIdEQ applies the EQ predicate on the "UserId" field.
func UserIdEQ(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// UserIdNEQ applies the NEQ predicate on the "UserId" field.
func UserIdNEQ(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserId), v))
	})
}

// UserIdIn applies the In predicate on the "UserId" field.
func UserIdIn(vs ...int) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserId), v...))
	})
}

// UserIdNotIn applies the NotIn predicate on the "UserId" field.
func UserIdNotIn(vs ...int) predicate.Appointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Appointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserId), v...))
	})
}

// UserIdGT applies the GT predicate on the "UserId" field.
func UserIdGT(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserId), v))
	})
}

// UserIdGTE applies the GTE predicate on the "UserId" field.
func UserIdGTE(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserId), v))
	})
}

// UserIdLT applies the LT predicate on the "UserId" field.
func UserIdLT(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserId), v))
	})
}

// UserIdLTE applies the LTE predicate on the "UserId" field.
func UserIdLTE(v int) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserId), v))
	})
}

// HasAppointmentSchedule applies the HasEdge predicate on the "appointment_schedule" edge.
func HasAppointmentSchedule() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppointmentScheduleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppointmentScheduleTable, AppointmentScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppointmentScheduleWith applies the HasEdge predicate on the "appointment_schedule" edge with a given conditions (other predicates).
func HasAppointmentScheduleWith(preds ...predicate.Schedule) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppointmentScheduleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppointmentScheduleTable, AppointmentScheduleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppointmentChat applies the HasEdge predicate on the "appointment_chat" edge.
func HasAppointmentChat() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppointmentChatTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppointmentChatTable, AppointmentChatColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppointmentChatWith applies the HasEdge predicate on the "appointment_chat" edge with a given conditions (other predicates).
func HasAppointmentChatWith(preds ...predicate.Chat) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppointmentChatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppointmentChatTable, AppointmentChatColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
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
func Not(p predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		p(s.Not())
	})
}