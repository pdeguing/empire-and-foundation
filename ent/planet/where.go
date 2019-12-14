// Code generated by entc, DO NOT EDIT.

package planet

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pdeguing/empire-and-foundation/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
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
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
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
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreatedAt), v))
		},
	)
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// MetalStock applies equality check predicate on the "metal_stock" field. It's identical to MetalStockEQ.
func MetalStock(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMetalStock), v))
		},
	)
}

// MetalMine applies equality check predicate on the "metal_mine" field. It's identical to MetalMineEQ.
func MetalMine(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMetalMine), v))
		},
	)
}

// LastMetalUpdate applies equality check predicate on the "last_metal_update" field. It's identical to LastMetalUpdateEQ.
func LastMetalUpdate(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldCreatedAt), v...))
		},
	)
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
		},
	)
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldCreatedAt), v))
		},
	)
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldCreatedAt), v))
		},
	)
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUpdatedAt), v...))
		},
	)
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
		},
	)
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUpdatedAt), v))
		},
	)
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
		},
	)
}

// MetalStockEQ applies the EQ predicate on the "metal_stock" field.
func MetalStockEQ(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMetalStock), v))
		},
	)
}

// MetalStockNEQ applies the NEQ predicate on the "metal_stock" field.
func MetalStockNEQ(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldMetalStock), v))
		},
	)
}

// MetalStockIn applies the In predicate on the "metal_stock" field.
func MetalStockIn(vs ...int64) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldMetalStock), v...))
		},
	)
}

// MetalStockNotIn applies the NotIn predicate on the "metal_stock" field.
func MetalStockNotIn(vs ...int64) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldMetalStock), v...))
		},
	)
}

// MetalStockGT applies the GT predicate on the "metal_stock" field.
func MetalStockGT(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldMetalStock), v))
		},
	)
}

// MetalStockGTE applies the GTE predicate on the "metal_stock" field.
func MetalStockGTE(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldMetalStock), v))
		},
	)
}

// MetalStockLT applies the LT predicate on the "metal_stock" field.
func MetalStockLT(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldMetalStock), v))
		},
	)
}

// MetalStockLTE applies the LTE predicate on the "metal_stock" field.
func MetalStockLTE(v int64) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldMetalStock), v))
		},
	)
}

// MetalMineEQ applies the EQ predicate on the "metal_mine" field.
func MetalMineEQ(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMetalMine), v))
		},
	)
}

// MetalMineNEQ applies the NEQ predicate on the "metal_mine" field.
func MetalMineNEQ(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldMetalMine), v))
		},
	)
}

// MetalMineIn applies the In predicate on the "metal_mine" field.
func MetalMineIn(vs ...int) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldMetalMine), v...))
		},
	)
}

// MetalMineNotIn applies the NotIn predicate on the "metal_mine" field.
func MetalMineNotIn(vs ...int) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldMetalMine), v...))
		},
	)
}

// MetalMineGT applies the GT predicate on the "metal_mine" field.
func MetalMineGT(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldMetalMine), v))
		},
	)
}

// MetalMineGTE applies the GTE predicate on the "metal_mine" field.
func MetalMineGTE(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldMetalMine), v))
		},
	)
}

// MetalMineLT applies the LT predicate on the "metal_mine" field.
func MetalMineLT(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldMetalMine), v))
		},
	)
}

// MetalMineLTE applies the LTE predicate on the "metal_mine" field.
func MetalMineLTE(v int) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldMetalMine), v))
		},
	)
}

// LastMetalUpdateEQ applies the EQ predicate on the "last_metal_update" field.
func LastMetalUpdateEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// LastMetalUpdateNEQ applies the NEQ predicate on the "last_metal_update" field.
func LastMetalUpdateNEQ(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// LastMetalUpdateIn applies the In predicate on the "last_metal_update" field.
func LastMetalUpdateIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldLastMetalUpdate), v...))
		},
	)
}

// LastMetalUpdateNotIn applies the NotIn predicate on the "last_metal_update" field.
func LastMetalUpdateNotIn(vs ...time.Time) predicate.Planet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Planet(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldLastMetalUpdate), v...))
		},
	)
}

// LastMetalUpdateGT applies the GT predicate on the "last_metal_update" field.
func LastMetalUpdateGT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// LastMetalUpdateGTE applies the GTE predicate on the "last_metal_update" field.
func LastMetalUpdateGTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// LastMetalUpdateLT applies the LT predicate on the "last_metal_update" field.
func LastMetalUpdateLT(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// LastMetalUpdateLTE applies the LTE predicate on the "last_metal_update" field.
func LastMetalUpdateLTE(v time.Time) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldLastMetalUpdate), v))
		},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(OwnerTable, FieldID),
				sql.Edge(sql.M2O, true, OwnerTable, OwnerColumn),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(OwnerInverseTable, FieldID),
				sql.Edge(sql.M2O, true, OwnerTable, OwnerColumn),
			)
			sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
				for _, p := range preds {
					p(s)
				}
			})
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Planet) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Planet) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Planet) predicate.Planet {
	return predicate.Planet(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
