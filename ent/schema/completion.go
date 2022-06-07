package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Completion holds the schema definition for the Completion entity.
type Completion struct {
	ent.Schema
}

// Fields of the Completion.
func (Completion) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").
		  NotEmpty().
			MaxLen(255),
    field.Time("date_completed").
		  Default(time.Now().UTC),
		field.Enum("status").
		  Values("completed", "missed", "skipped"),
	}
}

// Edges of the Completion.
func (Completion) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("habit", Habit.Type).Ref("completions").Unique().Required(),
		edge.From("user", User.Type).Ref("completions").Unique().Required(),
	}
}
