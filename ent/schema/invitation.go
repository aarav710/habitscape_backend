package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Invitation holds the schema definition for the Invitation entity.
type Invitation struct {
	ent.Schema
}

// Fields of the Invitation.
func (Invitation) Fields() []ent.Field {
	return nil
}

// Edges of the Invitation.
func (Invitation) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("habit", Habit.Type).Ref("invitations").Required().Unique(),
		edge.From("admin", Admin.Type).Ref("invitations").Required().Unique(),
    edge.From("user", User.Type).Ref("invitations").Required().Unique(),
	}
}
