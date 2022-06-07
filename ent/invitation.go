// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"habitscape/backend/ent/admin"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/invitation"
	"habitscape/backend/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Invitation is the model entity for the Invitation schema.
type Invitation struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InvitationQuery when eager-loading is set.
	Edges             InvitationEdges `json:"edges"`
	admin_invitations *int
	habit_invitations *int
	user_invitations  *int
}

// InvitationEdges holds the relations/edges for other nodes in the graph.
type InvitationEdges struct {
	// Habit holds the value of the habit edge.
	Habit *Habit `json:"habit,omitempty"`
	// Admin holds the value of the admin edge.
	Admin *Admin `json:"admin,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// HabitOrErr returns the Habit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) HabitOrErr() (*Habit, error) {
	if e.loadedTypes[0] {
		if e.Habit == nil {
			// The edge habit was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: habit.Label}
		}
		return e.Habit, nil
	}
	return nil, &NotLoadedError{edge: "habit"}
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) AdminOrErr() (*Admin, error) {
	if e.loadedTypes[1] {
		if e.Admin == nil {
			// The edge admin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: admin.Label}
		}
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InvitationEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Invitation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case invitation.FieldID:
			values[i] = new(sql.NullInt64)
		case invitation.ForeignKeys[0]: // admin_invitations
			values[i] = new(sql.NullInt64)
		case invitation.ForeignKeys[1]: // habit_invitations
			values[i] = new(sql.NullInt64)
		case invitation.ForeignKeys[2]: // user_invitations
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Invitation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Invitation fields.
func (i *Invitation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case invitation.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case invitation.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field admin_invitations", value)
			} else if value.Valid {
				i.admin_invitations = new(int)
				*i.admin_invitations = int(value.Int64)
			}
		case invitation.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field habit_invitations", value)
			} else if value.Valid {
				i.habit_invitations = new(int)
				*i.habit_invitations = int(value.Int64)
			}
		case invitation.ForeignKeys[2]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_invitations", value)
			} else if value.Valid {
				i.user_invitations = new(int)
				*i.user_invitations = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryHabit queries the "habit" edge of the Invitation entity.
func (i *Invitation) QueryHabit() *HabitQuery {
	return (&InvitationClient{config: i.config}).QueryHabit(i)
}

// QueryAdmin queries the "admin" edge of the Invitation entity.
func (i *Invitation) QueryAdmin() *AdminQuery {
	return (&InvitationClient{config: i.config}).QueryAdmin(i)
}

// QueryUser queries the "user" edge of the Invitation entity.
func (i *Invitation) QueryUser() *UserQuery {
	return (&InvitationClient{config: i.config}).QueryUser(i)
}

// Update returns a builder for updating this Invitation.
// Note that you need to call Invitation.Unwrap() before calling this method if this Invitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Invitation) Update() *InvitationUpdateOne {
	return (&InvitationClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Invitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Invitation) Unwrap() *Invitation {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Invitation is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Invitation) String() string {
	var builder strings.Builder
	builder.WriteString("Invitation(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Invitations is a parsable slice of Invitation.
type Invitations []*Invitation

func (i Invitations) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}