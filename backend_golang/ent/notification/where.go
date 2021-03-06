// Code generated by entc, DO NOT EDIT.

package notification

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Message applies equality check predicate on the "Message" field. It's identical to MessageEQ.
func Message(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// CreatedDate applies equality check predicate on the "CreatedDate" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// RecipientId applies equality check predicate on the "RecipientId" field. It's identical to RecipientIdEQ.
func RecipientId(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipientId), v))
	})
}

// SenderId applies equality check predicate on the "SenderId" field. It's identical to SenderIdEQ.
func SenderId(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSenderId), v))
	})
}

// AppointmentId applies equality check predicate on the "AppointmentId" field. It's identical to AppointmentIdEQ.
func AppointmentId(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointmentId), v))
	})
}

// MessageEQ applies the EQ predicate on the "Message" field.
func MessageEQ(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "Message" field.
func MessageNEQ(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "Message" field.
func MessageIn(vs ...string) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "Message" field.
func MessageNotIn(vs ...string) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "Message" field.
func MessageGT(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "Message" field.
func MessageGTE(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "Message" field.
func MessageLT(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "Message" field.
func MessageLTE(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "Message" field.
func MessageContains(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "Message" field.
func MessageHasPrefix(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "Message" field.
func MessageHasSuffix(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "Message" field.
func MessageEqualFold(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "Message" field.
func MessageContainsFold(v string) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// CreatedDateEQ applies the EQ predicate on the "CreatedDate" field.
func CreatedDateEQ(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateNEQ applies the NEQ predicate on the "CreatedDate" field.
func CreatedDateNEQ(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateIn applies the In predicate on the "CreatedDate" field.
func CreatedDateIn(vs ...time.Time) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateNotIn applies the NotIn predicate on the "CreatedDate" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedDate), v...))
	})
}

// CreatedDateGT applies the GT predicate on the "CreatedDate" field.
func CreatedDateGT(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateGTE applies the GTE predicate on the "CreatedDate" field.
func CreatedDateGTE(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLT applies the LT predicate on the "CreatedDate" field.
func CreatedDateLT(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateLTE applies the LTE predicate on the "CreatedDate" field.
func CreatedDateLTE(v time.Time) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedDate), v))
	})
}

// CreatedDateIsNil applies the IsNil predicate on the "CreatedDate" field.
func CreatedDateIsNil() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedDate)))
	})
}

// CreatedDateNotNil applies the NotNil predicate on the "CreatedDate" field.
func CreatedDateNotNil() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedDate)))
	})
}

// RecipientIdEQ applies the EQ predicate on the "RecipientId" field.
func RecipientIdEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipientId), v))
	})
}

// RecipientIdNEQ applies the NEQ predicate on the "RecipientId" field.
func RecipientIdNEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecipientId), v))
	})
}

// RecipientIdIn applies the In predicate on the "RecipientId" field.
func RecipientIdIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRecipientId), v...))
	})
}

// RecipientIdNotIn applies the NotIn predicate on the "RecipientId" field.
func RecipientIdNotIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRecipientId), v...))
	})
}

// RecipientIdGT applies the GT predicate on the "RecipientId" field.
func RecipientIdGT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecipientId), v))
	})
}

// RecipientIdGTE applies the GTE predicate on the "RecipientId" field.
func RecipientIdGTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecipientId), v))
	})
}

// RecipientIdLT applies the LT predicate on the "RecipientId" field.
func RecipientIdLT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecipientId), v))
	})
}

// RecipientIdLTE applies the LTE predicate on the "RecipientId" field.
func RecipientIdLTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecipientId), v))
	})
}

// SenderIdEQ applies the EQ predicate on the "SenderId" field.
func SenderIdEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSenderId), v))
	})
}

// SenderIdNEQ applies the NEQ predicate on the "SenderId" field.
func SenderIdNEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSenderId), v))
	})
}

// SenderIdIn applies the In predicate on the "SenderId" field.
func SenderIdIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSenderId), v...))
	})
}

// SenderIdNotIn applies the NotIn predicate on the "SenderId" field.
func SenderIdNotIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSenderId), v...))
	})
}

// SenderIdGT applies the GT predicate on the "SenderId" field.
func SenderIdGT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSenderId), v))
	})
}

// SenderIdGTE applies the GTE predicate on the "SenderId" field.
func SenderIdGTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSenderId), v))
	})
}

// SenderIdLT applies the LT predicate on the "SenderId" field.
func SenderIdLT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSenderId), v))
	})
}

// SenderIdLTE applies the LTE predicate on the "SenderId" field.
func SenderIdLTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSenderId), v))
	})
}

// AppointmentIdEQ applies the EQ predicate on the "AppointmentId" field.
func AppointmentIdEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointmentId), v))
	})
}

// AppointmentIdNEQ applies the NEQ predicate on the "AppointmentId" field.
func AppointmentIdNEQ(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppointmentId), v))
	})
}

// AppointmentIdIn applies the In predicate on the "AppointmentId" field.
func AppointmentIdIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppointmentId), v...))
	})
}

// AppointmentIdNotIn applies the NotIn predicate on the "AppointmentId" field.
func AppointmentIdNotIn(vs ...int) predicate.Notification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppointmentId), v...))
	})
}

// AppointmentIdGT applies the GT predicate on the "AppointmentId" field.
func AppointmentIdGT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppointmentId), v))
	})
}

// AppointmentIdGTE applies the GTE predicate on the "AppointmentId" field.
func AppointmentIdGTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppointmentId), v))
	})
}

// AppointmentIdLT applies the LT predicate on the "AppointmentId" field.
func AppointmentIdLT(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppointmentId), v))
	})
}

// AppointmentIdLTE applies the LTE predicate on the "AppointmentId" field.
func AppointmentIdLTE(v int) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppointmentId), v))
	})
}

// HasUserNotification applies the HasEdge predicate on the "user_notification" edge.
func HasUserNotification() predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserNotificationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserNotificationTable, UserNotificationPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserNotificationWith applies the HasEdge predicate on the "user_notification" edge with a given conditions (other predicates).
func HasUserNotificationWith(preds ...predicate.User) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserNotificationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserNotificationTable, UserNotificationPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
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
func Not(p predicate.Notification) predicate.Notification {
	return predicate.Notification(func(s *sql.Selector) {
		p(s.Not())
	})
}
