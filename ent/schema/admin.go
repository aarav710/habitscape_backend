package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return nil
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("admins").Unique().Required(),
		edge.From("habit", Habit.Type).Ref("admins").Unique().Required(),
		edge.To("invitations", Invitation.Type),
	}
}
