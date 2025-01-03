package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Share struct {
	ent.Schema
}

func (Share) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").NotEmpty(),
		field.String("share_to").NotEmpty(), // platform or user shared to
		field.Time("created_at").Default(time.Now),
	}
}

func (Share) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("shares").
			Unique().
			Required(),
	}
}
