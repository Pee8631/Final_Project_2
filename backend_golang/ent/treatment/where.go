// Code generated by entc, DO NOT EDIT.

package treatment

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TreatmentRecord applies equality check predicate on the "treatmentRecord" field. It's identical to TreatmentRecordEQ.
func TreatmentRecord(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTreatmentRecord), v))
	})
}

// DateTime applies equality check predicate on the "dateTime" field. It's identical to DateTimeEQ.
func DateTime(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTime), v))
	})
}

// TakeTime applies equality check predicate on the "takeTime" field. It's identical to TakeTimeEQ.
func TakeTime(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTakeTime), v))
	})
}

// TreatmentRecordEQ applies the EQ predicate on the "treatmentRecord" field.
func TreatmentRecordEQ(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordNEQ applies the NEQ predicate on the "treatmentRecord" field.
func TreatmentRecordNEQ(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordIn applies the In predicate on the "treatmentRecord" field.
func TreatmentRecordIn(vs ...string) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTreatmentRecord), v...))
	})
}

// TreatmentRecordNotIn applies the NotIn predicate on the "treatmentRecord" field.
func TreatmentRecordNotIn(vs ...string) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTreatmentRecord), v...))
	})
}

// TreatmentRecordGT applies the GT predicate on the "treatmentRecord" field.
func TreatmentRecordGT(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordGTE applies the GTE predicate on the "treatmentRecord" field.
func TreatmentRecordGTE(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordLT applies the LT predicate on the "treatmentRecord" field.
func TreatmentRecordLT(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordLTE applies the LTE predicate on the "treatmentRecord" field.
func TreatmentRecordLTE(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordContains applies the Contains predicate on the "treatmentRecord" field.
func TreatmentRecordContains(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordHasPrefix applies the HasPrefix predicate on the "treatmentRecord" field.
func TreatmentRecordHasPrefix(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordHasSuffix applies the HasSuffix predicate on the "treatmentRecord" field.
func TreatmentRecordHasSuffix(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordEqualFold applies the EqualFold predicate on the "treatmentRecord" field.
func TreatmentRecordEqualFold(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTreatmentRecord), v))
	})
}

// TreatmentRecordContainsFold applies the ContainsFold predicate on the "treatmentRecord" field.
func TreatmentRecordContainsFold(v string) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTreatmentRecord), v))
	})
}

// DateTimeEQ applies the EQ predicate on the "dateTime" field.
func DateTimeEQ(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTime), v))
	})
}

// DateTimeNEQ applies the NEQ predicate on the "dateTime" field.
func DateTimeNEQ(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateTime), v))
	})
}

// DateTimeIn applies the In predicate on the "dateTime" field.
func DateTimeIn(vs ...time.Time) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateTime), v...))
	})
}

// DateTimeNotIn applies the NotIn predicate on the "dateTime" field.
func DateTimeNotIn(vs ...time.Time) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateTime), v...))
	})
}

// DateTimeGT applies the GT predicate on the "dateTime" field.
func DateTimeGT(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateTime), v))
	})
}

// DateTimeGTE applies the GTE predicate on the "dateTime" field.
func DateTimeGTE(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateTime), v))
	})
}

// DateTimeLT applies the LT predicate on the "dateTime" field.
func DateTimeLT(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateTime), v))
	})
}

// DateTimeLTE applies the LTE predicate on the "dateTime" field.
func DateTimeLTE(v time.Time) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateTime), v))
	})
}

// DateTimeIsNil applies the IsNil predicate on the "dateTime" field.
func DateTimeIsNil() predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDateTime)))
	})
}

// DateTimeNotNil applies the NotNil predicate on the "dateTime" field.
func DateTimeNotNil() predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDateTime)))
	})
}

// TakeTimeEQ applies the EQ predicate on the "takeTime" field.
func TakeTimeEQ(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTakeTime), v))
	})
}

// TakeTimeNEQ applies the NEQ predicate on the "takeTime" field.
func TakeTimeNEQ(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTakeTime), v))
	})
}

// TakeTimeIn applies the In predicate on the "takeTime" field.
func TakeTimeIn(vs ...float64) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTakeTime), v...))
	})
}

// TakeTimeNotIn applies the NotIn predicate on the "takeTime" field.
func TakeTimeNotIn(vs ...float64) predicate.Treatment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Treatment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTakeTime), v...))
	})
}

// TakeTimeGT applies the GT predicate on the "takeTime" field.
func TakeTimeGT(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTakeTime), v))
	})
}

// TakeTimeGTE applies the GTE predicate on the "takeTime" field.
func TakeTimeGTE(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTakeTime), v))
	})
}

// TakeTimeLT applies the LT predicate on the "takeTime" field.
func TakeTimeLT(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTakeTime), v))
	})
}

// TakeTimeLTE applies the LTE predicate on the "takeTime" field.
func TakeTimeLTE(v float64) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTakeTime), v))
	})
}

// HasTreatmentWasRecordedByDoctor applies the HasEdge predicate on the "treatment_was_recorded_by_doctor" edge.
func HasTreatmentWasRecordedByDoctor() predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentWasRecordedByDoctorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TreatmentWasRecordedByDoctorTable, TreatmentWasRecordedByDoctorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTreatmentWasRecordedByDoctorWith applies the HasEdge predicate on the "treatment_was_recorded_by_doctor" edge with a given conditions (other predicates).
func HasTreatmentWasRecordedByDoctorWith(preds ...predicate.User) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentWasRecordedByDoctorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TreatmentWasRecordedByDoctorTable, TreatmentWasRecordedByDoctorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserIsTheTreatmentOfRecord applies the HasEdge predicate on the "user_is_the_treatment_of_record" edge.
func HasUserIsTheTreatmentOfRecord() predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIsTheTreatmentOfRecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIsTheTreatmentOfRecordTable, UserIsTheTreatmentOfRecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIsTheTreatmentOfRecordWith applies the HasEdge predicate on the "user_is_the_treatment_of_record" edge with a given conditions (other predicates).
func HasUserIsTheTreatmentOfRecordWith(preds ...predicate.User) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIsTheTreatmentOfRecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIsTheTreatmentOfRecordTable, UserIsTheTreatmentOfRecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Treatment) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Treatment) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
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
func Not(p predicate.Treatment) predicate.Treatment {
	return predicate.Treatment(func(s *sql.Selector) {
		p(s.Not())
	})
}
