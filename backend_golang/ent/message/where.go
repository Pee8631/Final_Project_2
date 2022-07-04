// Code generated by entc, DO NOT EDIT.

package message

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// MessageText applies equality check predicate on the "message_text" field. It's identical to MessageTextEQ.
func MessageText(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageText), v))
	})
}

// SentDateTime applies equality check predicate on the "sent_dateTime" field. It's identical to SentDateTimeEQ.
func SentDateTime(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentDateTime), v))
	})
}

// MessageTextEQ applies the EQ predicate on the "message_text" field.
func MessageTextEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageText), v))
	})
}

// MessageTextNEQ applies the NEQ predicate on the "message_text" field.
func MessageTextNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessageText), v))
	})
}

// MessageTextIn applies the In predicate on the "message_text" field.
func MessageTextIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessageText), v...))
	})
}

// MessageTextNotIn applies the NotIn predicate on the "message_text" field.
func MessageTextNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessageText), v...))
	})
}

// MessageTextGT applies the GT predicate on the "message_text" field.
func MessageTextGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessageText), v))
	})
}

// MessageTextGTE applies the GTE predicate on the "message_text" field.
func MessageTextGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessageText), v))
	})
}

// MessageTextLT applies the LT predicate on the "message_text" field.
func MessageTextLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessageText), v))
	})
}

// MessageTextLTE applies the LTE predicate on the "message_text" field.
func MessageTextLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessageText), v))
	})
}

// MessageTextContains applies the Contains predicate on the "message_text" field.
func MessageTextContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessageText), v))
	})
}

// MessageTextHasPrefix applies the HasPrefix predicate on the "message_text" field.
func MessageTextHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessageText), v))
	})
}

// MessageTextHasSuffix applies the HasSuffix predicate on the "message_text" field.
func MessageTextHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessageText), v))
	})
}

// MessageTextEqualFold applies the EqualFold predicate on the "message_text" field.
func MessageTextEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessageText), v))
	})
}

// MessageTextContainsFold applies the ContainsFold predicate on the "message_text" field.
func MessageTextContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessageText), v))
	})
}

// SentDateTimeEQ applies the EQ predicate on the "sent_dateTime" field.
func SentDateTimeEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeNEQ applies the NEQ predicate on the "sent_dateTime" field.
func SentDateTimeNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeIn applies the In predicate on the "sent_dateTime" field.
func SentDateTimeIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSentDateTime), v...))
	})
}

// SentDateTimeNotIn applies the NotIn predicate on the "sent_dateTime" field.
func SentDateTimeNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSentDateTime), v...))
	})
}

// SentDateTimeGT applies the GT predicate on the "sent_dateTime" field.
func SentDateTimeGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeGTE applies the GTE predicate on the "sent_dateTime" field.
func SentDateTimeGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeLT applies the LT predicate on the "sent_dateTime" field.
func SentDateTimeLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeLTE applies the LTE predicate on the "sent_dateTime" field.
func SentDateTimeLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSentDateTime), v))
	})
}

// SentDateTimeIsNil applies the IsNil predicate on the "sent_dateTime" field.
func SentDateTimeIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSentDateTime)))
	})
}

// SentDateTimeNotNil applies the NotNil predicate on the "sent_dateTime" field.
func SentDateTimeNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSentDateTime)))
	})
}

// HasWhatMessagesAreInThisChat applies the HasEdge predicate on the "What_messages_are_in_this_chat" edge.
func HasWhatMessagesAreInThisChat() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhatMessagesAreInThisChatTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhatMessagesAreInThisChatTable, WhatMessagesAreInThisChatColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhatMessagesAreInThisChatWith applies the HasEdge predicate on the "What_messages_are_in_this_chat" edge with a given conditions (other predicates).
func HasWhatMessagesAreInThisChatWith(preds ...predicate.Chat) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhatMessagesAreInThisChatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhatMessagesAreInThisChatTable, WhatMessagesAreInThisChatColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWhoSendMessages applies the HasEdge predicate on the "Who_send_messages" edge.
func HasWhoSendMessages() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoSendMessagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhoSendMessagesTable, WhoSendMessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhoSendMessagesWith applies the HasEdge predicate on the "Who_send_messages" edge with a given conditions (other predicates).
func HasWhoSendMessagesWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoSendMessagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhoSendMessagesTable, WhoSendMessagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}