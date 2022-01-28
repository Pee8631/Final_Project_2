// Code generated by entc, DO NOT EDIT.

package certification

import (
	"FinalProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// Diloma applies equality check predicate on the "diloma" field. It's identical to DilomaEQ.
func Diloma(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiloma), v))
	})
}

// DateOfIssuing applies equality check predicate on the "dateOfIssuing" field. It's identical to DateOfIssuingEQ.
func DateOfIssuing(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfExp applies equality check predicate on the "dateOfExp" field. It's identical to DateOfExpEQ.
func DateOfExp(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateOfExp), v))
	})
}

// Issuer applies equality check predicate on the "Issuer" field. It's identical to IssuerEQ.
func Issuer(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIssuer), v))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// DilomaEQ applies the EQ predicate on the "diloma" field.
func DilomaEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiloma), v))
	})
}

// DilomaNEQ applies the NEQ predicate on the "diloma" field.
func DilomaNEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiloma), v))
	})
}

// DilomaIn applies the In predicate on the "diloma" field.
func DilomaIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiloma), v...))
	})
}

// DilomaNotIn applies the NotIn predicate on the "diloma" field.
func DilomaNotIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiloma), v...))
	})
}

// DilomaGT applies the GT predicate on the "diloma" field.
func DilomaGT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiloma), v))
	})
}

// DilomaGTE applies the GTE predicate on the "diloma" field.
func DilomaGTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiloma), v))
	})
}

// DilomaLT applies the LT predicate on the "diloma" field.
func DilomaLT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiloma), v))
	})
}

// DilomaLTE applies the LTE predicate on the "diloma" field.
func DilomaLTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiloma), v))
	})
}

// DilomaContains applies the Contains predicate on the "diloma" field.
func DilomaContains(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiloma), v))
	})
}

// DilomaHasPrefix applies the HasPrefix predicate on the "diloma" field.
func DilomaHasPrefix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiloma), v))
	})
}

// DilomaHasSuffix applies the HasSuffix predicate on the "diloma" field.
func DilomaHasSuffix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiloma), v))
	})
}

// DilomaEqualFold applies the EqualFold predicate on the "diloma" field.
func DilomaEqualFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiloma), v))
	})
}

// DilomaContainsFold applies the ContainsFold predicate on the "diloma" field.
func DilomaContainsFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiloma), v))
	})
}

// DateOfIssuingEQ applies the EQ predicate on the "dateOfIssuing" field.
func DateOfIssuingEQ(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfIssuingNEQ applies the NEQ predicate on the "dateOfIssuing" field.
func DateOfIssuingNEQ(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfIssuingIn applies the In predicate on the "dateOfIssuing" field.
func DateOfIssuingIn(vs ...time.Time) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateOfIssuing), v...))
	})
}

// DateOfIssuingNotIn applies the NotIn predicate on the "dateOfIssuing" field.
func DateOfIssuingNotIn(vs ...time.Time) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateOfIssuing), v...))
	})
}

// DateOfIssuingGT applies the GT predicate on the "dateOfIssuing" field.
func DateOfIssuingGT(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfIssuingGTE applies the GTE predicate on the "dateOfIssuing" field.
func DateOfIssuingGTE(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfIssuingLT applies the LT predicate on the "dateOfIssuing" field.
func DateOfIssuingLT(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfIssuingLTE applies the LTE predicate on the "dateOfIssuing" field.
func DateOfIssuingLTE(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateOfIssuing), v))
	})
}

// DateOfExpEQ applies the EQ predicate on the "dateOfExp" field.
func DateOfExpEQ(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateOfExp), v))
	})
}

// DateOfExpNEQ applies the NEQ predicate on the "dateOfExp" field.
func DateOfExpNEQ(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateOfExp), v))
	})
}

// DateOfExpIn applies the In predicate on the "dateOfExp" field.
func DateOfExpIn(vs ...time.Time) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateOfExp), v...))
	})
}

// DateOfExpNotIn applies the NotIn predicate on the "dateOfExp" field.
func DateOfExpNotIn(vs ...time.Time) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateOfExp), v...))
	})
}

// DateOfExpGT applies the GT predicate on the "dateOfExp" field.
func DateOfExpGT(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateOfExp), v))
	})
}

// DateOfExpGTE applies the GTE predicate on the "dateOfExp" field.
func DateOfExpGTE(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateOfExp), v))
	})
}

// DateOfExpLT applies the LT predicate on the "dateOfExp" field.
func DateOfExpLT(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateOfExp), v))
	})
}

// DateOfExpLTE applies the LTE predicate on the "dateOfExp" field.
func DateOfExpLTE(v time.Time) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateOfExp), v))
	})
}

// IssuerEQ applies the EQ predicate on the "Issuer" field.
func IssuerEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIssuer), v))
	})
}

// IssuerNEQ applies the NEQ predicate on the "Issuer" field.
func IssuerNEQ(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIssuer), v))
	})
}

// IssuerIn applies the In predicate on the "Issuer" field.
func IssuerIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIssuer), v...))
	})
}

// IssuerNotIn applies the NotIn predicate on the "Issuer" field.
func IssuerNotIn(vs ...string) predicate.Certification {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certification(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIssuer), v...))
	})
}

// IssuerGT applies the GT predicate on the "Issuer" field.
func IssuerGT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIssuer), v))
	})
}

// IssuerGTE applies the GTE predicate on the "Issuer" field.
func IssuerGTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIssuer), v))
	})
}

// IssuerLT applies the LT predicate on the "Issuer" field.
func IssuerLT(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIssuer), v))
	})
}

// IssuerLTE applies the LTE predicate on the "Issuer" field.
func IssuerLTE(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIssuer), v))
	})
}

// IssuerContains applies the Contains predicate on the "Issuer" field.
func IssuerContains(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIssuer), v))
	})
}

// IssuerHasPrefix applies the HasPrefix predicate on the "Issuer" field.
func IssuerHasPrefix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIssuer), v))
	})
}

// IssuerHasSuffix applies the HasSuffix predicate on the "Issuer" field.
func IssuerHasSuffix(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIssuer), v))
	})
}

// IssuerEqualFold applies the EqualFold predicate on the "Issuer" field.
func IssuerEqualFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIssuer), v))
	})
}

// IssuerContainsFold applies the ContainsFold predicate on the "Issuer" field.
func IssuerContainsFold(v string) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIssuer), v))
	})
}

// HasDoctorOwner applies the HasEdge predicate on the "doctor_owner" edge.
func HasDoctorOwner() predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorOwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorOwnerTable, DoctorOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorOwnerWith applies the HasEdge predicate on the "doctor_owner" edge with a given conditions (other predicates).
func HasDoctorOwnerWith(preds ...predicate.User) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorOwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorOwnerTable, DoctorOwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Certification) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Certification) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
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
func Not(p predicate.Certification) predicate.Certification {
	return predicate.Certification(func(s *sql.Selector) {
		p(s.Not())
	})
}
