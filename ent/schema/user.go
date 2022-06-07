package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("bio").
		    MaxLen(255).Nillable().Optional(),
		field.String("email").
		    Unique(),
		field.String("username").
				Unique(),
		field.Time("created_at").
				Default(time.Now().UTC),
    field.String("photo_url"),	
  }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("habits", Habit.Type),
	  edge.To("admins", Admin.Type),
		edge.To("comments", Comment.Type),
		edge.To("posts", Post.Type),
		edge.To("completions", Completion.Type),
		edge.To("likes", Like.Type),
		edge.To("invitations", Invitation.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
			// unique index.
			index.Fields("email", "username").
					Unique(),
	}
}