// Code generated by ent, DO NOT EDIT.

package share

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sundayonah/phindcode_backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Share {
	return predicate.Share(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Share {
	return predicate.Share(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Share {
	return predicate.Share(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Share {
	return predicate.Share(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Share {
	return predicate.Share(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Share {
	return predicate.Share(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Share {
	return predicate.Share(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldUserID, v))
}

// ShareTo applies equality check predicate on the "share_to" field. It's identical to ShareToEQ.
func ShareTo(v string) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldShareTo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Share {
	return predicate.Share(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Share {
	return predicate.Share(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Share {
	return predicate.Share(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Share {
	return predicate.Share(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Share {
	return predicate.Share(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Share {
	return predicate.Share(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Share {
	return predicate.Share(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Share {
	return predicate.Share(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Share {
	return predicate.Share(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Share {
	return predicate.Share(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Share {
	return predicate.Share(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Share {
	return predicate.Share(sql.FieldContainsFold(FieldUserID, v))
}

// ShareToEQ applies the EQ predicate on the "share_to" field.
func ShareToEQ(v string) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldShareTo, v))
}

// ShareToNEQ applies the NEQ predicate on the "share_to" field.
func ShareToNEQ(v string) predicate.Share {
	return predicate.Share(sql.FieldNEQ(FieldShareTo, v))
}

// ShareToIn applies the In predicate on the "share_to" field.
func ShareToIn(vs ...string) predicate.Share {
	return predicate.Share(sql.FieldIn(FieldShareTo, vs...))
}

// ShareToNotIn applies the NotIn predicate on the "share_to" field.
func ShareToNotIn(vs ...string) predicate.Share {
	return predicate.Share(sql.FieldNotIn(FieldShareTo, vs...))
}

// ShareToGT applies the GT predicate on the "share_to" field.
func ShareToGT(v string) predicate.Share {
	return predicate.Share(sql.FieldGT(FieldShareTo, v))
}

// ShareToGTE applies the GTE predicate on the "share_to" field.
func ShareToGTE(v string) predicate.Share {
	return predicate.Share(sql.FieldGTE(FieldShareTo, v))
}

// ShareToLT applies the LT predicate on the "share_to" field.
func ShareToLT(v string) predicate.Share {
	return predicate.Share(sql.FieldLT(FieldShareTo, v))
}

// ShareToLTE applies the LTE predicate on the "share_to" field.
func ShareToLTE(v string) predicate.Share {
	return predicate.Share(sql.FieldLTE(FieldShareTo, v))
}

// ShareToContains applies the Contains predicate on the "share_to" field.
func ShareToContains(v string) predicate.Share {
	return predicate.Share(sql.FieldContains(FieldShareTo, v))
}

// ShareToHasPrefix applies the HasPrefix predicate on the "share_to" field.
func ShareToHasPrefix(v string) predicate.Share {
	return predicate.Share(sql.FieldHasPrefix(FieldShareTo, v))
}

// ShareToHasSuffix applies the HasSuffix predicate on the "share_to" field.
func ShareToHasSuffix(v string) predicate.Share {
	return predicate.Share(sql.FieldHasSuffix(FieldShareTo, v))
}

// ShareToEqualFold applies the EqualFold predicate on the "share_to" field.
func ShareToEqualFold(v string) predicate.Share {
	return predicate.Share(sql.FieldEqualFold(FieldShareTo, v))
}

// ShareToContainsFold applies the ContainsFold predicate on the "share_to" field.
func ShareToContainsFold(v string) predicate.Share {
	return predicate.Share(sql.FieldContainsFold(FieldShareTo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Share {
	return predicate.Share(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Share {
	return predicate.Share(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Share {
	return predicate.Share(sql.FieldLTE(FieldCreatedAt, v))
}

// HasPost applies the HasEdge predicate on the "post" edge.
func HasPost() predicate.Share {
	return predicate.Share(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.Share {
	return predicate.Share(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Share) predicate.Share {
	return predicate.Share(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Share) predicate.Share {
	return predicate.Share(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Share) predicate.Share {
	return predicate.Share(sql.NotPredicates(p))
}