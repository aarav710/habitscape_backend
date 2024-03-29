// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"habitscape/backend/ent/completion"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Completion is the model entity for the Completion schema.
type Completion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// DateCompleted holds the value of the "date_completed" field.
	DateCompleted time.Time `json:"date_completed,omitempty"`
	// Status holds the value of the "status" field.
	Status completion.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompletionQuery when eager-loading is set.
	Edges             CompletionEdges `json:"edges"`
	habit_completions *int
	user_completions  *int
}

// CompletionEdges holds the relations/edges for other nodes in the graph.
type CompletionEdges struct {
	// Habit holds the value of the habit edge.
	Habit *Habit `json:"habit,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HabitOrErr returns the Habit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompletionEdges) HabitOrErr() (*Habit, error) {
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

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompletionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
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
func (*Completion) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case completion.FieldID:
			values[i] = new(sql.NullInt64)
		case completion.FieldText, completion.FieldStatus:
			values[i] = new(sql.NullString)
		case completion.FieldDateCompleted:
			values[i] = new(sql.NullTime)
		case completion.ForeignKeys[0]: // habit_completions
			values[i] = new(sql.NullInt64)
		case completion.ForeignKeys[1]: // user_completions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Completion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Completion fields.
func (c *Completion) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case completion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case completion.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case completion.FieldDateCompleted:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_completed", values[i])
			} else if value.Valid {
				c.DateCompleted = value.Time
			}
		case completion.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = completion.Status(value.String)
			}
		case completion.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field habit_completions", value)
			} else if value.Valid {
				c.habit_completions = new(int)
				*c.habit_completions = int(value.Int64)
			}
		case completion.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_completions", value)
			} else if value.Valid {
				c.user_completions = new(int)
				*c.user_completions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryHabit queries the "habit" edge of the Completion entity.
func (c *Completion) QueryHabit() *HabitQuery {
	return (&CompletionClient{config: c.config}).QueryHabit(c)
}

// QueryUser queries the "user" edge of the Completion entity.
func (c *Completion) QueryUser() *UserQuery {
	return (&CompletionClient{config: c.config}).QueryUser(c)
}

// Update returns a builder for updating this Completion.
// Note that you need to call Completion.Unwrap() before calling this method if this Completion
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Completion) Update() *CompletionUpdateOne {
	return (&CompletionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Completion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Completion) Unwrap() *Completion {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Completion is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Completion) String() string {
	var builder strings.Builder
	builder.WriteString("Completion(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", text=")
	builder.WriteString(c.Text)
	builder.WriteString(", date_completed=")
	builder.WriteString(c.DateCompleted.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Completions is a parsable slice of Completion.
type Completions []*Completion

func (c Completions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
