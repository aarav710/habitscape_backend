// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"habitscape/backend/ent/comment"
	"habitscape/backend/ent/post"
	"habitscape/backend/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges         CommentEdges `json:"edges"`
	post_comments *int
	user_comments *int
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Post == nil {
			// The edge post was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) UserOrErr() (*User, error) {
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
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			values[i] = new(sql.NullInt64)
		case comment.FieldText:
			values[i] = new(sql.NullString)
		case comment.FieldDateCreated:
			values[i] = new(sql.NullTime)
		case comment.ForeignKeys[0]: // post_comments
			values[i] = new(sql.NullInt64)
		case comment.ForeignKeys[1]: // user_comments
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case comment.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				c.DateCreated = value.Time
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field post_comments", value)
			} else if value.Valid {
				c.post_comments = new(int)
				*c.post_comments = int(value.Int64)
			}
		case comment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_comments", value)
			} else if value.Valid {
				c.user_comments = new(int)
				*c.user_comments = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPost queries the "post" edge of the Comment entity.
func (c *Comment) QueryPost() *PostQuery {
	return (&CommentClient{config: c.config}).QueryPost(c)
}

// QueryUser queries the "user" edge of the Comment entity.
func (c *Comment) QueryUser() *UserQuery {
	return (&CommentClient{config: c.config}).QueryUser(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", text=")
	builder.WriteString(c.Text)
	builder.WriteString(", date_created=")
	builder.WriteString(c.DateCreated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}