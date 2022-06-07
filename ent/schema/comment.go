package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
	"entgo.io/ent/schema/edge"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").MaxLen(255).NotEmpty(),
		field.Time("date_created").Default(time.Now().UTC),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("comments").Unique().Required(),
		edge.From("user", User.Type).Ref("comments").Unique().Required(),
	}
}
