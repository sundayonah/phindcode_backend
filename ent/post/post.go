// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeShares holds the string denoting the shares edge name in mutations.
	EdgeShares = "shares"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "post_likes"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "post_comments"
	// SharesTable is the table that holds the shares relation/edge.
	SharesTable = "shares"
	// SharesInverseTable is the table name for the Share entity.
	// It exists in this package in order to avoid circular dependency with the "share" package.
	SharesInverseTable = "shares"
	// SharesColumn is the table column denoting the shares relation/edge.
	SharesColumn = "post_shares"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldImage,
	FieldCategory,
	FieldCode,
	FieldUserID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	CategoryValidator func(string) error
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Post queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLikesCount orders the results by likes count.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikesStep(), opts...)
	}
}

// ByLikes orders the results by likes terms.
func ByLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySharesCount orders the results by shares count.
func BySharesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSharesStep(), opts...)
	}
}

// ByShares orders the results by shares terms.
func ByShares(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSharesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
	)
}
func newSharesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SharesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SharesTable, SharesColumn),
	)
}
