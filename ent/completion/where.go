// Code generated by entc, DO NOT EDIT.

package completion

import (
	"habitscape/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// DateCompleted applies equality check predicate on the "date_completed" field. It's identical to DateCompletedEQ.
func DateCompleted(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCompleted), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// DateCompletedEQ applies the EQ predicate on the "date_completed" field.
func DateCompletedEQ(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCompleted), v))
	})
}

// DateCompletedNEQ applies the NEQ predicate on the "date_completed" field.
func DateCompletedNEQ(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateCompleted), v))
	})
}

// DateCompletedIn applies the In predicate on the "date_completed" field.
func DateCompletedIn(vs ...time.Time) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateCompleted), v...))
	})
}

// DateCompletedNotIn applies the NotIn predicate on the "date_completed" field.
func DateCompletedNotIn(vs ...time.Time) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateCompleted), v...))
	})
}

// DateCompletedGT applies the GT predicate on the "date_completed" field.
func DateCompletedGT(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateCompleted), v))
	})
}

// DateCompletedGTE applies the GTE predicate on the "date_completed" field.
func DateCompletedGTE(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateCompleted), v))
	})
}

// DateCompletedLT applies the LT predicate on the "date_completed" field.
func DateCompletedLT(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateCompleted), v))
	})
}

// DateCompletedLTE applies the LTE predicate on the "date_completed" field.
func DateCompletedLTE(v time.Time) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateCompleted), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
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
func StatusNotIn(vs ...Status) predicate.Completion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Completion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// HasHabit applies the HasEdge predicate on the "habit" edge.
func HasHabit() predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HabitTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HabitTable, HabitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHabitWith applies the HasEdge predicate on the "habit" edge with a given conditions (other predicates).
func HasHabitWith(preds ...predicate.Habit) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HabitInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HabitTable, HabitColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Completion) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Completion) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
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
func Not(p predicate.Completion) predicate.Completion {
	return predicate.Completion(func(s *sql.Selector) {
		p(s.Not())
	})
}
