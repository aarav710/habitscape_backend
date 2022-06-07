package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("photo_url"),
		field.String("caption").MaxLen(255),
    field.Time("date_created").
		  Default(time.Now().UTC),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
    edge.From("habit", Habit.Type).Ref("posts").Unique().Required(),
		edge.To("comments", Comment.Type),
		edge.From("user", User.Type).Ref("posts").Unique().Required(),
		edge.To("likes", Like.Type),
	}
}
