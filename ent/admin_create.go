// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"habitscape/backend/ent/admin"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/invitation"
	"habitscape/backend/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AdminCreate) SetUserID(id int) *AdminCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AdminCreate) SetUser(u *User) *AdminCreate {
	return ac.SetUserID(u.ID)
}

// SetHabitID sets the "habit" edge to the Habit entity by ID.
func (ac *AdminCreate) SetHabitID(id int) *AdminCreate {
	ac.mutation.SetHabitID(id)
	return ac
}

// SetHabit sets the "habit" edge to the Habit entity.
func (ac *AdminCreate) SetHabit(h *Habit) *AdminCreate {
	return ac.SetHabitID(h.ID)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (ac *AdminCreate) AddInvitationIDs(ids ...int) *AdminCreate {
	ac.mutation.AddInvitationIDs(ids...)
	return ac
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (ac *AdminCreate) AddInvitations(i ...*Invitation) *AdminCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ac.AddInvitationIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Admin.user"`)}
	}
	if _, ok := ac.mutation.HabitID(); !ok {
		return &ValidationError{Name: "habit", err: errors.New(`ent: missing required edge "Admin.habit"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: admin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		}
	)
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.UserTable,
			Columns: []string{admin.UserColumn},
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
		_node.user_admins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.HabitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   admin.HabitTable,
			Columns: []string{admin.HabitColumn},
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
		_node.habit_admins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.InvitationsTable,
			Columns: []string{admin.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	builders []*AdminCreate
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
