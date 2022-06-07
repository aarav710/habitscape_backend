// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"habitscape/backend/ent/admin"
	"habitscape/backend/ent/completion"
	"habitscape/backend/ent/habit"
	"habitscape/backend/ent/invitation"
	"habitscape/backend/ent/post"
	"habitscape/backend/ent/predicate"
	"habitscape/backend/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HabitUpdate is the builder for updating Habit entities.
type HabitUpdate struct {
	config
	hooks    []Hook
	mutation *HabitMutation
}

// Where appends a list predicates to the HabitUpdate builder.
func (hu *HabitUpdate) Where(ps ...predicate.Habit) *HabitUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetDescription sets the "description" field.
func (hu *HabitUpdate) SetDescription(s string) *HabitUpdate {
	hu.mutation.SetDescription(s)
	return hu
}

// SetFrequency sets the "frequency" field.
func (hu *HabitUpdate) SetFrequency(i int) *HabitUpdate {
	hu.mutation.ResetFrequency()
	hu.mutation.SetFrequency(i)
	return hu
}

// AddFrequency adds i to the "frequency" field.
func (hu *HabitUpdate) AddFrequency(i int) *HabitUpdate {
	hu.mutation.AddFrequency(i)
	return hu
}

// SetTitle sets the "title" field.
func (hu *HabitUpdate) SetTitle(s string) *HabitUpdate {
	hu.mutation.SetTitle(s)
	return hu
}

// SetPhotoURL sets the "photo_url" field.
func (hu *HabitUpdate) SetPhotoURL(s string) *HabitUpdate {
	hu.mutation.SetPhotoURL(s)
	return hu
}

// SetDateCreated sets the "date_created" field.
func (hu *HabitUpdate) SetDateCreated(t time.Time) *HabitUpdate {
	hu.mutation.SetDateCreated(t)
	return hu
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (hu *HabitUpdate) SetNillableDateCreated(t *time.Time) *HabitUpdate {
	if t != nil {
		hu.SetDateCreated(*t)
	}
	return hu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (hu *HabitUpdate) AddUserIDs(ids ...int) *HabitUpdate {
	hu.mutation.AddUserIDs(ids...)
	return hu
}

// AddUsers adds the "users" edges to the User entity.
func (hu *HabitUpdate) AddUsers(u ...*User) *HabitUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hu.AddUserIDs(ids...)
}

// AddAdminIDs adds the "admins" edge to the Admin entity by IDs.
func (hu *HabitUpdate) AddAdminIDs(ids ...int) *HabitUpdate {
	hu.mutation.AddAdminIDs(ids...)
	return hu
}

// AddAdmins adds the "admins" edges to the Admin entity.
func (hu *HabitUpdate) AddAdmins(a ...*Admin) *HabitUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hu.AddAdminIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (hu *HabitUpdate) AddPostIDs(ids ...int) *HabitUpdate {
	hu.mutation.AddPostIDs(ids...)
	return hu
}

// AddPosts adds the "posts" edges to the Post entity.
func (hu *HabitUpdate) AddPosts(p ...*Post) *HabitUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return hu.AddPostIDs(ids...)
}

// AddCompletionIDs adds the "completions" edge to the Completion entity by IDs.
func (hu *HabitUpdate) AddCompletionIDs(ids ...int) *HabitUpdate {
	hu.mutation.AddCompletionIDs(ids...)
	return hu
}

// AddCompletions adds the "completions" edges to the Completion entity.
func (hu *HabitUpdate) AddCompletions(c ...*Completion) *HabitUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hu.AddCompletionIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (hu *HabitUpdate) AddInvitationIDs(ids ...int) *HabitUpdate {
	hu.mutation.AddInvitationIDs(ids...)
	return hu
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (hu *HabitUpdate) AddInvitations(i ...*Invitation) *HabitUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return hu.AddInvitationIDs(ids...)
}

// Mutation returns the HabitMutation object of the builder.
func (hu *HabitUpdate) Mutation() *HabitMutation {
	return hu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (hu *HabitUpdate) ClearUsers() *HabitUpdate {
	hu.mutation.ClearUsers()
	return hu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (hu *HabitUpdate) RemoveUserIDs(ids ...int) *HabitUpdate {
	hu.mutation.RemoveUserIDs(ids...)
	return hu
}

// RemoveUsers removes "users" edges to User entities.
func (hu *HabitUpdate) RemoveUsers(u ...*User) *HabitUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hu.RemoveUserIDs(ids...)
}

// ClearAdmins clears all "admins" edges to the Admin entity.
func (hu *HabitUpdate) ClearAdmins() *HabitUpdate {
	hu.mutation.ClearAdmins()
	return hu
}

// RemoveAdminIDs removes the "admins" edge to Admin entities by IDs.
func (hu *HabitUpdate) RemoveAdminIDs(ids ...int) *HabitUpdate {
	hu.mutation.RemoveAdminIDs(ids...)
	return hu
}

// RemoveAdmins removes "admins" edges to Admin entities.
func (hu *HabitUpdate) RemoveAdmins(a ...*Admin) *HabitUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return hu.RemoveAdminIDs(ids...)
}

// ClearPosts clears all "posts" edges to the Post entity.
func (hu *HabitUpdate) ClearPosts() *HabitUpdate {
	hu.mutation.ClearPosts()
	return hu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (hu *HabitUpdate) RemovePostIDs(ids ...int) *HabitUpdate {
	hu.mutation.RemovePostIDs(ids...)
	return hu
}

// RemovePosts removes "posts" edges to Post entities.
func (hu *HabitUpdate) RemovePosts(p ...*Post) *HabitUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return hu.RemovePostIDs(ids...)
}

// ClearCompletions clears all "completions" edges to the Completion entity.
func (hu *HabitUpdate) ClearCompletions() *HabitUpdate {
	hu.mutation.ClearCompletions()
	return hu
}

// RemoveCompletionIDs removes the "completions" edge to Completion entities by IDs.
func (hu *HabitUpdate) RemoveCompletionIDs(ids ...int) *HabitUpdate {
	hu.mutation.RemoveCompletionIDs(ids...)
	return hu
}

// RemoveCompletions removes "completions" edges to Completion entities.
func (hu *HabitUpdate) RemoveCompletions(c ...*Completion) *HabitUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hu.RemoveCompletionIDs(ids...)
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (hu *HabitUpdate) ClearInvitations() *HabitUpdate {
	hu.mutation.ClearInvitations()
	return hu
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (hu *HabitUpdate) RemoveInvitationIDs(ids ...int) *HabitUpdate {
	hu.mutation.RemoveInvitationIDs(ids...)
	return hu
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (hu *HabitUpdate) RemoveInvitations(i ...*Invitation) *HabitUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return hu.RemoveInvitationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HabitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		if err = hu.check(); err != nil {
			return 0, err
		}
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HabitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hu.check(); err != nil {
				return 0, err
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			if hu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HabitUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HabitUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HabitUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HabitUpdate) check() error {
	if v, ok := hu.mutation.Description(); ok {
		if err := habit.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Habit.description": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Frequency(); ok {
		if err := habit.FrequencyValidator(v); err != nil {
			return &ValidationError{Name: "frequency", err: fmt.Errorf(`ent: validator failed for field "Habit.frequency": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Title(); ok {
		if err := habit.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Habit.title": %w`, err)}
		}
	}
	return nil
}

func (hu *HabitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   habit.Table,
			Columns: habit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: habit.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldDescription,
		})
	}
	if value, ok := hu.mutation.Frequency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: habit.FieldFrequency,
		})
	}
	if value, ok := hu.mutation.AddedFrequency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: habit.FieldFrequency,
		})
	}
	if value, ok := hu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldTitle,
		})
	}
	if value, ok := hu.mutation.PhotoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldPhotoURL,
		})
	}
	if value, ok := hu.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: habit.FieldDateCreated,
		})
	}
	if hu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !hu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedAdminsIDs(); len(nodes) > 0 && !hu.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !hu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.CompletionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedCompletionsIDs(); len(nodes) > 0 && !hu.mutation.CompletionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.CompletionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !hu.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{habit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// HabitUpdateOne is the builder for updating a single Habit entity.
type HabitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HabitMutation
}

// SetDescription sets the "description" field.
func (huo *HabitUpdateOne) SetDescription(s string) *HabitUpdateOne {
	huo.mutation.SetDescription(s)
	return huo
}

// SetFrequency sets the "frequency" field.
func (huo *HabitUpdateOne) SetFrequency(i int) *HabitUpdateOne {
	huo.mutation.ResetFrequency()
	huo.mutation.SetFrequency(i)
	return huo
}

// AddFrequency adds i to the "frequency" field.
func (huo *HabitUpdateOne) AddFrequency(i int) *HabitUpdateOne {
	huo.mutation.AddFrequency(i)
	return huo
}

// SetTitle sets the "title" field.
func (huo *HabitUpdateOne) SetTitle(s string) *HabitUpdateOne {
	huo.mutation.SetTitle(s)
	return huo
}

// SetPhotoURL sets the "photo_url" field.
func (huo *HabitUpdateOne) SetPhotoURL(s string) *HabitUpdateOne {
	huo.mutation.SetPhotoURL(s)
	return huo
}

// SetDateCreated sets the "date_created" field.
func (huo *HabitUpdateOne) SetDateCreated(t time.Time) *HabitUpdateOne {
	huo.mutation.SetDateCreated(t)
	return huo
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (huo *HabitUpdateOne) SetNillableDateCreated(t *time.Time) *HabitUpdateOne {
	if t != nil {
		huo.SetDateCreated(*t)
	}
	return huo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (huo *HabitUpdateOne) AddUserIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.AddUserIDs(ids...)
	return huo
}

// AddUsers adds the "users" edges to the User entity.
func (huo *HabitUpdateOne) AddUsers(u ...*User) *HabitUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return huo.AddUserIDs(ids...)
}

// AddAdminIDs adds the "admins" edge to the Admin entity by IDs.
func (huo *HabitUpdateOne) AddAdminIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.AddAdminIDs(ids...)
	return huo
}

// AddAdmins adds the "admins" edges to the Admin entity.
func (huo *HabitUpdateOne) AddAdmins(a ...*Admin) *HabitUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return huo.AddAdminIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (huo *HabitUpdateOne) AddPostIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.AddPostIDs(ids...)
	return huo
}

// AddPosts adds the "posts" edges to the Post entity.
func (huo *HabitUpdateOne) AddPosts(p ...*Post) *HabitUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return huo.AddPostIDs(ids...)
}

// AddCompletionIDs adds the "completions" edge to the Completion entity by IDs.
func (huo *HabitUpdateOne) AddCompletionIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.AddCompletionIDs(ids...)
	return huo
}

// AddCompletions adds the "completions" edges to the Completion entity.
func (huo *HabitUpdateOne) AddCompletions(c ...*Completion) *HabitUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return huo.AddCompletionIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (huo *HabitUpdateOne) AddInvitationIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.AddInvitationIDs(ids...)
	return huo
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (huo *HabitUpdateOne) AddInvitations(i ...*Invitation) *HabitUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return huo.AddInvitationIDs(ids...)
}

// Mutation returns the HabitMutation object of the builder.
func (huo *HabitUpdateOne) Mutation() *HabitMutation {
	return huo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (huo *HabitUpdateOne) ClearUsers() *HabitUpdateOne {
	huo.mutation.ClearUsers()
	return huo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (huo *HabitUpdateOne) RemoveUserIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.RemoveUserIDs(ids...)
	return huo
}

// RemoveUsers removes "users" edges to User entities.
func (huo *HabitUpdateOne) RemoveUsers(u ...*User) *HabitUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return huo.RemoveUserIDs(ids...)
}

// ClearAdmins clears all "admins" edges to the Admin entity.
func (huo *HabitUpdateOne) ClearAdmins() *HabitUpdateOne {
	huo.mutation.ClearAdmins()
	return huo
}

// RemoveAdminIDs removes the "admins" edge to Admin entities by IDs.
func (huo *HabitUpdateOne) RemoveAdminIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.RemoveAdminIDs(ids...)
	return huo
}

// RemoveAdmins removes "admins" edges to Admin entities.
func (huo *HabitUpdateOne) RemoveAdmins(a ...*Admin) *HabitUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return huo.RemoveAdminIDs(ids...)
}

// ClearPosts clears all "posts" edges to the Post entity.
func (huo *HabitUpdateOne) ClearPosts() *HabitUpdateOne {
	huo.mutation.ClearPosts()
	return huo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (huo *HabitUpdateOne) RemovePostIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.RemovePostIDs(ids...)
	return huo
}

// RemovePosts removes "posts" edges to Post entities.
func (huo *HabitUpdateOne) RemovePosts(p ...*Post) *HabitUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return huo.RemovePostIDs(ids...)
}

// ClearCompletions clears all "completions" edges to the Completion entity.
func (huo *HabitUpdateOne) ClearCompletions() *HabitUpdateOne {
	huo.mutation.ClearCompletions()
	return huo
}

// RemoveCompletionIDs removes the "completions" edge to Completion entities by IDs.
func (huo *HabitUpdateOne) RemoveCompletionIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.RemoveCompletionIDs(ids...)
	return huo
}

// RemoveCompletions removes "completions" edges to Completion entities.
func (huo *HabitUpdateOne) RemoveCompletions(c ...*Completion) *HabitUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return huo.RemoveCompletionIDs(ids...)
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (huo *HabitUpdateOne) ClearInvitations() *HabitUpdateOne {
	huo.mutation.ClearInvitations()
	return huo
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (huo *HabitUpdateOne) RemoveInvitationIDs(ids ...int) *HabitUpdateOne {
	huo.mutation.RemoveInvitationIDs(ids...)
	return huo
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (huo *HabitUpdateOne) RemoveInvitations(i ...*Invitation) *HabitUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return huo.RemoveInvitationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HabitUpdateOne) Select(field string, fields ...string) *HabitUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Habit entity.
func (huo *HabitUpdateOne) Save(ctx context.Context) (*Habit, error) {
	var (
		err  error
		node *Habit
	)
	if len(huo.hooks) == 0 {
		if err = huo.check(); err != nil {
			return nil, err
		}
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HabitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = huo.check(); err != nil {
				return nil, err
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			if huo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HabitUpdateOne) SaveX(ctx context.Context) *Habit {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HabitUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HabitUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HabitUpdateOne) check() error {
	if v, ok := huo.mutation.Description(); ok {
		if err := habit.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Habit.description": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Frequency(); ok {
		if err := habit.FrequencyValidator(v); err != nil {
			return &ValidationError{Name: "frequency", err: fmt.Errorf(`ent: validator failed for field "Habit.frequency": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Title(); ok {
		if err := habit.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Habit.title": %w`, err)}
		}
	}
	return nil
}

func (huo *HabitUpdateOne) sqlSave(ctx context.Context) (_node *Habit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   habit.Table,
			Columns: habit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: habit.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Habit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, habit.FieldID)
		for _, f := range fields {
			if !habit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != habit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldDescription,
		})
	}
	if value, ok := huo.mutation.Frequency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: habit.FieldFrequency,
		})
	}
	if value, ok := huo.mutation.AddedFrequency(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: habit.FieldFrequency,
		})
	}
	if value, ok := huo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldTitle,
		})
	}
	if value, ok := huo.mutation.PhotoURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: habit.FieldPhotoURL,
		})
	}
	if value, ok := huo.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: habit.FieldDateCreated,
		})
	}
	if huo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !huo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   habit.UsersTable,
			Columns: habit.UsersPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedAdminsIDs(); len(nodes) > 0 && !huo.mutation.AdminsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.AdminsTable,
			Columns: []string{habit.AdminsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !huo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.PostsTable,
			Columns: []string{habit.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.CompletionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedCompletionsIDs(); len(nodes) > 0 && !huo.mutation.CompletionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.CompletionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.CompletionsTable,
			Columns: []string{habit.CompletionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: completion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: invitation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !huo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   habit.InvitationsTable,
			Columns: []string{habit.InvitationsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Habit{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{habit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
