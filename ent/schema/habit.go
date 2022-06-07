package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Habit holds the schema definition for the Habit entity.
type Habit struct {
	ent.Schema
}

// Fields of the Habit.
func (Habit) Fields() []ent.Field {
	return []ent.Field{
    field.String("description").
		  MaxLen(255),
		field.Int("frequency").
		  Range(1,7),
		field.String("title").
		  NotEmpty(),
		field.String("photo_url"),
		field.Time("date_created").
		  Default(time.Now().UTC),
	}
}

// Edges of the Habit.
func (Habit) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("users", User.Type).Ref("habits"),
		edge.To("admins", Admin.Type),
		edge.To("posts", Post.Type),
		edge.To("completions", Completion.Type),
		edge.To("invitations", Invitation.Type),
	}
}
