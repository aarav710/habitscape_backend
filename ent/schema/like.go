package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return nil
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user", User.Type).Ref("likes").Unique().Required(),
		edge.From("post", Post.Type).Ref("likes").Unique().Required(),
	}
}
