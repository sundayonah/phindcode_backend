// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sundayonah/phindcode_backend/ent/comment"
	"github.com/sundayonah/phindcode_backend/ent/like"
	"github.com/sundayonah/phindcode_backend/ent/post"
	"github.com/sundayonah/phindcode_backend/ent/predicate"
	"github.com/sundayonah/phindcode_backend/ent/share"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PostUpdate) SetDescription(s string) *PostUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PostUpdate) SetNillableDescription(s *string) *PostUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetImage sets the "image" field.
func (pu *PostUpdate) SetImage(s string) *PostUpdate {
	pu.mutation.SetImage(s)
	return pu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (pu *PostUpdate) SetNillableImage(s *string) *PostUpdate {
	if s != nil {
		pu.SetImage(*s)
	}
	return pu
}

// ClearImage clears the value of the "image" field.
func (pu *PostUpdate) ClearImage() *PostUpdate {
	pu.mutation.ClearImage()
	return pu
}

// SetCategory sets the "category" field.
func (pu *PostUpdate) SetCategory(s string) *PostUpdate {
	pu.mutation.SetCategory(s)
	return pu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCategory(s *string) *PostUpdate {
	if s != nil {
		pu.SetCategory(*s)
	}
	return pu
}

// SetCode sets the "code" field.
func (pu *PostUpdate) SetCode(s string) *PostUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCode(s *string) *PostUpdate {
	if s != nil {
		pu.SetCode(*s)
	}
	return pu
}

// ClearCode clears the value of the "code" field.
func (pu *PostUpdate) ClearCode() *PostUpdate {
	pu.mutation.ClearCode()
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(s string) *PostUpdate {
	pu.mutation.SetUserID(s)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableUserID(s *string) *PostUpdate {
	if s != nil {
		pu.SetUserID(*s)
	}
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pu *PostUpdate) AddLikeIDs(ids ...int) *PostUpdate {
	pu.mutation.AddLikeIDs(ids...)
	return pu
}

// AddLikes adds the "likes" edges to the Like entity.
func (pu *PostUpdate) AddLikes(l ...*Like) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.AddLikeIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PostUpdate) AddCommentIDs(ids ...int) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// AddShareIDs adds the "shares" edge to the Share entity by IDs.
func (pu *PostUpdate) AddShareIDs(ids ...int) *PostUpdate {
	pu.mutation.AddShareIDs(ids...)
	return pu
}

// AddShares adds the "shares" edges to the Share entity.
func (pu *PostUpdate) AddShares(s ...*Share) *PostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddShareIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearLikes clears all "likes" edges to the Like entity.
func (pu *PostUpdate) ClearLikes() *PostUpdate {
	pu.mutation.ClearLikes()
	return pu
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (pu *PostUpdate) RemoveLikeIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveLikeIDs(ids...)
	return pu
}

// RemoveLikes removes "likes" edges to Like entities.
func (pu *PostUpdate) RemoveLikes(l ...*Like) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.RemoveLikeIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PostUpdate) ClearComments() *PostUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PostUpdate) RemoveCommentIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// ClearShares clears all "shares" edges to the Share entity.
func (pu *PostUpdate) ClearShares() *PostUpdate {
	pu.mutation.ClearShares()
	return pu
}

// RemoveShareIDs removes the "shares" edge to Share entities by IDs.
func (pu *PostUpdate) RemoveShareIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveShareIDs(ids...)
	return pu
}

// RemoveShares removes "shares" edges to Share entities.
func (pu *PostUpdate) RemoveShares(s ...*Share) *PostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveShareIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Description(); ok {
		if err := post.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Post.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Category(); ok {
		if err := post.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Post.category": %w`, err)}
		}
	}
	if v, ok := pu.mutation.UserID(); ok {
		if err := post.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Post.user_id": %w`, err)}
		}
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Image(); ok {
		_spec.SetField(post.FieldImage, field.TypeString, value)
	}
	if pu.mutation.ImageCleared() {
		_spec.ClearField(post.FieldImage, field.TypeString)
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.SetField(post.FieldCategory, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(post.FieldCode, field.TypeString, value)
	}
	if pu.mutation.CodeCleared() {
		_spec.ClearField(post.FieldCode, field.TypeString)
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLikesIDs(); len(nodes) > 0 && !pu.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSharesIDs(); len(nodes) > 0 && !pu.mutation.SharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetDescription sets the "description" field.
func (puo *PostUpdateOne) SetDescription(s string) *PostUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableDescription(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetImage sets the "image" field.
func (puo *PostUpdateOne) SetImage(s string) *PostUpdateOne {
	puo.mutation.SetImage(s)
	return puo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableImage(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetImage(*s)
	}
	return puo
}

// ClearImage clears the value of the "image" field.
func (puo *PostUpdateOne) ClearImage() *PostUpdateOne {
	puo.mutation.ClearImage()
	return puo
}

// SetCategory sets the "category" field.
func (puo *PostUpdateOne) SetCategory(s string) *PostUpdateOne {
	puo.mutation.SetCategory(s)
	return puo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCategory(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetCategory(*s)
	}
	return puo
}

// SetCode sets the "code" field.
func (puo *PostUpdateOne) SetCode(s string) *PostUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCode(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetCode(*s)
	}
	return puo
}

// ClearCode clears the value of the "code" field.
func (puo *PostUpdateOne) ClearCode() *PostUpdateOne {
	puo.mutation.ClearCode()
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(s string) *PostUpdateOne {
	puo.mutation.SetUserID(s)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUserID(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetUserID(*s)
	}
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (puo *PostUpdateOne) AddLikeIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddLikeIDs(ids...)
	return puo
}

// AddLikes adds the "likes" edges to the Like entity.
func (puo *PostUpdateOne) AddLikes(l ...*Like) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.AddLikeIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PostUpdateOne) AddCommentIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// AddShareIDs adds the "shares" edge to the Share entity by IDs.
func (puo *PostUpdateOne) AddShareIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddShareIDs(ids...)
	return puo
}

// AddShares adds the "shares" edges to the Share entity.
func (puo *PostUpdateOne) AddShares(s ...*Share) *PostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddShareIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearLikes clears all "likes" edges to the Like entity.
func (puo *PostUpdateOne) ClearLikes() *PostUpdateOne {
	puo.mutation.ClearLikes()
	return puo
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (puo *PostUpdateOne) RemoveLikeIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveLikeIDs(ids...)
	return puo
}

// RemoveLikes removes "likes" edges to Like entities.
func (puo *PostUpdateOne) RemoveLikes(l ...*Like) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.RemoveLikeIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PostUpdateOne) ClearComments() *PostUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// ClearShares clears all "shares" edges to the Share entity.
func (puo *PostUpdateOne) ClearShares() *PostUpdateOne {
	puo.mutation.ClearShares()
	return puo
}

// RemoveShareIDs removes the "shares" edge to Share entities by IDs.
func (puo *PostUpdateOne) RemoveShareIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveShareIDs(ids...)
	return puo
}

// RemoveShares removes "shares" edges to Share entities.
func (puo *PostUpdateOne) RemoveShares(s ...*Share) *PostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveShareIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Description(); ok {
		if err := post.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Post.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Category(); ok {
		if err := post.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Post.category": %w`, err)}
		}
	}
	if v, ok := puo.mutation.UserID(); ok {
		if err := post.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Post.user_id": %w`, err)}
		}
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Image(); ok {
		_spec.SetField(post.FieldImage, field.TypeString, value)
	}
	if puo.mutation.ImageCleared() {
		_spec.ClearField(post.FieldImage, field.TypeString)
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.SetField(post.FieldCategory, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(post.FieldCode, field.TypeString, value)
	}
	if puo.mutation.CodeCleared() {
		_spec.ClearField(post.FieldCode, field.TypeString)
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(post.FieldUserID, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLikesIDs(); len(nodes) > 0 && !puo.mutation.LikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSharesIDs(); len(nodes) > 0 && !puo.mutation.SharesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SharesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.SharesTable,
			Columns: []string{post.SharesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(share.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
