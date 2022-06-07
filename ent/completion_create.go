// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"habitscape/backend/ent/completion"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompletionCreate is the builder for creating a Completion entity.
type CompletionCreate struct {
	config
	mutation *CompletionMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (cc *CompletionCreate) SetText(s string) *CompletionCreate {
	cc.mutation.SetText(s)
	return cc
}

// SetDateCompleted sets the "date_completed" field.
func (cc *CompletionCreate) SetDateCompleted(t time.Time) *CompletionCreate {
	cc.mutation.SetDateCompleted(t)
	return cc
}

// SetNillableDateCompleted sets the "date_completed" field if the given value is not nil.
func (cc *CompletionCreate) SetNillableDateCompleted(t *time.Time) *CompletionCreate {
	if t != nil {
		cc.SetDateCompleted(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CompletionCreate) SetStatus(c completion.Status) *CompletionCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetHabitID sets the "habit" edge to the Habit entity by ID.
func (cc *CompletionCreate) SetHabitID(id int) *CompletionCreate {
	cc.mutation.SetHabitID(id)
	return cc
}

// SetHabit sets the "habit" edge to the Habit entity.
func (cc *CompletionCreate) SetHabit(h *Habit) *CompletionCreate {
	return cc.SetHabitID(h.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *CompletionCreate) SetUserID(id int) *CompletionCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CompletionCreate) SetUser(u *User) *CompletionCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CompletionMutation object of the builder.
func (cc *CompletionCreate) Mutation() *CompletionMutation {
	return cc.mutation
}

// Save creates the Completion in the database.
func (cc *CompletionCreate) Save(ctx context.Context) (*Completion, error) {
	var (
		err  error
		node *Completion
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompletionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompletionCreate) SaveX(ctx context.Context) *Completion {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompletionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompletionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompletionCreate) defaults() {
	if _, ok := cc.mutation.DateCompleted(); !ok {
		v := completion.DefaultDateCompleted()
		cc.mutation.SetDateCompleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompletionCreate) check() error {
	if _, ok := cc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Completion.text"`)}
	}
	if v, ok := cc.mutation.Text(); ok {
		if err := completion.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Completion.text": %w`, err)}
		}
	}
	if _, ok := cc.mutation.DateCompleted(); !ok {
		return &ValidationError{Name: "date_completed", err: errors.New(`ent: missing required field "Completion.date_completed"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Completion.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := completion.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Completion.status": %w`, err)}
		}
	}
	if _, ok := cc.mutation.HabitID(); !ok {
		return &ValidationError{Name: "habit", err: errors.New(`ent: missing required edge "Completion.habit"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Completion.user"`)}
	}
	return nil
}

func (cc *CompletionCreate) sqlSave(ctx context.Context) (*Completion, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CompletionCreate) createSpec() (*Completion, *sqlgraph.CreateSpec) {
	var (
		_node = &Completion{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: completion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: completion.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: completion.FieldText,
		})
		_node.Text = value
	}
	if value, ok := cc.mutation.DateCompleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: completion.FieldDateCompleted,
		})
		_node.DateCompleted = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: completion.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := cc.mutation.HabitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   completion.HabitTable,
			Columns: []string{completion.HabitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: habit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.habit_completions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   completion.UserTable,
			Columns: []string{completion.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_completions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompletionCreateBulk is the builder for creating many Completion entities in bulk.
type CompletionCreateBulk struct {
	config
	builders []*CompletionCreate
}

// Save creates the Completion entities in the database.
func (ccb *CompletionCreateBulk) Save(ctx context.Context) ([]*Completion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Completion, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompletionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompletionCreateBulk) SaveX(ctx context.Context) []*Completion {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompletionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompletionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}