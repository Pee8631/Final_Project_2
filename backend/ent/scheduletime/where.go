// Code generated by entc, DO NOT EDIT.

package scheduletime

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StopTime applies equality check predicate on the "stopTime" field. It's identical to StopTimeEQ.
func StopTime(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStopTime), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.ScheduleTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScheduleTime(func(s *sql.Selector) {
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
func StartTimeNotIn(vs ...time.Time) predicate.ScheduleTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScheduleTime(func(s *sql.Selector) {
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
func StartTimeGT(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// StopTimeEQ applies the EQ predicate on the "stopTime" field.
func StopTimeEQ(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStopTime), v))
	})
}

// StopTimeNEQ applies the NEQ predicate on the "stopTime" field.
func StopTimeNEQ(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStopTime), v))
	})
}

// StopTimeIn applies the In predicate on the "stopTime" field.
func StopTimeIn(vs ...time.Time) predicate.ScheduleTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScheduleTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStopTime), v...))
	})
}

// StopTimeNotIn applies the NotIn predicate on the "stopTime" field.
func StopTimeNotIn(vs ...time.Time) predicate.ScheduleTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScheduleTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStopTime), v...))
	})
}

// StopTimeGT applies the GT predicate on the "stopTime" field.
func StopTimeGT(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStopTime), v))
	})
}

// StopTimeGTE applies the GTE predicate on the "stopTime" field.
func StopTimeGTE(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStopTime), v))
	})
}

// StopTimeLT applies the LT predicate on the "stopTime" field.
func StopTimeLT(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStopTime), v))
	})
}

// StopTimeLTE applies the LTE predicate on the "stopTime" field.
func StopTimeLTE(v time.Time) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStopTime), v))
	})
}

// HasWhatTimeIsTheSchedule applies the HasEdge predicate on the "what_time_is_the_schedule" edge.
func HasWhatTimeIsTheSchedule() predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhatTimeIsTheScheduleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhatTimeIsTheScheduleTable, WhatTimeIsTheScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhatTimeIsTheScheduleWith applies the HasEdge predicate on the "what_time_is_the_schedule" edge with a given conditions (other predicates).
func HasWhatTimeIsTheScheduleWith(preds ...predicate.Schedule) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhatTimeIsTheScheduleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhatTimeIsTheScheduleTable, WhatTimeIsTheScheduleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScheduleTime) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScheduleTime) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
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
func Not(p predicate.ScheduleTime) predicate.ScheduleTime {
	return predicate.ScheduleTime(func(s *sql.Selector) {
		p(s.Not())
	})
}
