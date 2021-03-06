// Code generated by entc, DO NOT EDIT.

package telecom

import (
	"FinalProject/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Platform applies equality check predicate on the "platform" field. It's identical to PlatformEQ.
func Platform(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatform), v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephone), v))
	})
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephone), v))
	})
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelephone), v...))
	})
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelephone), v...))
	})
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephone), v))
	})
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephone), v))
	})
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephone), v))
	})
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephone), v))
	})
}

// TelephoneContains applies the Contains predicate on the "telephone" field.
func TelephoneContains(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTelephone), v))
	})
}

// TelephoneHasPrefix applies the HasPrefix predicate on the "telephone" field.
func TelephoneHasPrefix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTelephone), v))
	})
}

// TelephoneHasSuffix applies the HasSuffix predicate on the "telephone" field.
func TelephoneHasSuffix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTelephone), v))
	})
}

// TelephoneEqualFold applies the EqualFold predicate on the "telephone" field.
func TelephoneEqualFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTelephone), v))
	})
}

// TelephoneContainsFold applies the ContainsFold predicate on the "telephone" field.
func TelephoneContainsFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTelephone), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
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
func UsernameNotIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
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
func UsernameGT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatform), v))
	})
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlatform), v))
	})
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlatform), v...))
	})
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...string) predicate.Telecom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Telecom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlatform), v...))
	})
}

// PlatformGT applies the GT predicate on the "platform" field.
func PlatformGT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlatform), v))
	})
}

// PlatformGTE applies the GTE predicate on the "platform" field.
func PlatformGTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlatform), v))
	})
}

// PlatformLT applies the LT predicate on the "platform" field.
func PlatformLT(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlatform), v))
	})
}

// PlatformLTE applies the LTE predicate on the "platform" field.
func PlatformLTE(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlatform), v))
	})
}

// PlatformContains applies the Contains predicate on the "platform" field.
func PlatformContains(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlatform), v))
	})
}

// PlatformHasPrefix applies the HasPrefix predicate on the "platform" field.
func PlatformHasPrefix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlatform), v))
	})
}

// PlatformHasSuffix applies the HasSuffix predicate on the "platform" field.
func PlatformHasSuffix(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlatform), v))
	})
}

// PlatformEqualFold applies the EqualFold predicate on the "platform" field.
func PlatformEqualFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlatform), v))
	})
}

// PlatformContainsFold applies the ContainsFold predicate on the "platform" field.
func PlatformContainsFold(v string) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlatform), v))
	})
}

// HasWhoIsTheOwnerOfThisTelecom applies the HasEdge predicate on the "who_is_the_owner_of_this_telecom" edge.
func HasWhoIsTheOwnerOfThisTelecom() predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoIsTheOwnerOfThisTelecomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhoIsTheOwnerOfThisTelecomTable, WhoIsTheOwnerOfThisTelecomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWhoIsTheOwnerOfThisTelecomWith applies the HasEdge predicate on the "who_is_the_owner_of_this_telecom" edge with a given conditions (other predicates).
func HasWhoIsTheOwnerOfThisTelecomWith(preds ...predicate.User) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WhoIsTheOwnerOfThisTelecomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WhoIsTheOwnerOfThisTelecomTable, WhoIsTheOwnerOfThisTelecomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Telecom) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Telecom) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
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
func Not(p predicate.Telecom) predicate.Telecom {
	return predicate.Telecom(func(s *sql.Selector) {
		p(s.Not())
	})
}
