// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"imdbv2/ent/director"
	"imdbv2/ent/movie"
	"imdbv2/ent/predicate"
	"imdbv2/ent/review"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MovieUpdate is the builder for updating Movie entities.
type MovieUpdate struct {
	config
	hooks    []Hook
	mutation *MovieMutation
}

// Where appends a list predicates to the MovieUpdate builder.
func (mu *MovieUpdate) Where(ps ...predicate.Movie) *MovieUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MovieUpdate) SetTitle(s string) *MovieUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MovieUpdate) SetDescription(s string) *MovieUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetRank sets the "rank" field.
func (mu *MovieUpdate) SetRank(i int) *MovieUpdate {
	mu.mutation.ResetRank()
	mu.mutation.SetRank(i)
	return mu
}

// AddRank adds i to the "rank" field.
func (mu *MovieUpdate) AddRank(i int) *MovieUpdate {
	mu.mutation.AddRank(i)
	return mu
}

// SetGenre sets the "genre" field.
func (mu *MovieUpdate) SetGenre(s string) *MovieUpdate {
	mu.mutation.SetGenre(s)
	return mu
}

// SetDirectorID sets the "director_id" field.
func (mu *MovieUpdate) SetDirectorID(i int) *MovieUpdate {
	mu.mutation.SetDirectorID(i)
	return mu
}

// SetNillableDirectorID sets the "director_id" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableDirectorID(i *int) *MovieUpdate {
	if i != nil {
		mu.SetDirectorID(*i)
	}
	return mu
}

// ClearDirectorID clears the value of the "director_id" field.
func (mu *MovieUpdate) ClearDirectorID() *MovieUpdate {
	mu.mutation.ClearDirectorID()
	return mu
}

// SetDirector sets the "director" edge to the Director entity.
func (mu *MovieUpdate) SetDirector(d *Director) *MovieUpdate {
	return mu.SetDirectorID(d.ID)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (mu *MovieUpdate) AddReviewIDs(ids ...int) *MovieUpdate {
	mu.mutation.AddReviewIDs(ids...)
	return mu
}

// AddReviews adds the "reviews" edges to the Review entity.
func (mu *MovieUpdate) AddReviews(r ...*Review) *MovieUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.AddReviewIDs(ids...)
}

// Mutation returns the MovieMutation object of the builder.
func (mu *MovieUpdate) Mutation() *MovieMutation {
	return mu.mutation
}

// ClearDirector clears the "director" edge to the Director entity.
func (mu *MovieUpdate) ClearDirector() *MovieUpdate {
	mu.mutation.ClearDirector()
	return mu
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (mu *MovieUpdate) ClearReviews() *MovieUpdate {
	mu.mutation.ClearReviews()
	return mu
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (mu *MovieUpdate) RemoveReviewIDs(ids ...int) *MovieUpdate {
	mu.mutation.RemoveReviewIDs(ids...)
	return mu
}

// RemoveReviews removes "reviews" edges to Review entities.
func (mu *MovieUpdate) RemoveReviews(r ...*Review) *MovieUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.RemoveReviewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MovieUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MovieUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MovieUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MovieUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MovieUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   movie.Table,
			Columns: movie.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: movie.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldTitle,
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldDescription,
		})
	}
	if value, ok := mu.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: movie.FieldRank,
		})
	}
	if value, ok := mu.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: movie.FieldRank,
		})
	}
	if value, ok := mu.mutation.Genre(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldGenre,
		})
	}
	if mu.mutation.DirectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   movie.DirectorTable,
			Columns: []string{movie.DirectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: director.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DirectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   movie.DirectorTable,
			Columns: []string{movie.DirectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: director.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !mu.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MovieUpdateOne is the builder for updating a single Movie entity.
type MovieUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieMutation
}

// SetTitle sets the "title" field.
func (muo *MovieUpdateOne) SetTitle(s string) *MovieUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetDescription sets the "description" field.
func (muo *MovieUpdateOne) SetDescription(s string) *MovieUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetRank sets the "rank" field.
func (muo *MovieUpdateOne) SetRank(i int) *MovieUpdateOne {
	muo.mutation.ResetRank()
	muo.mutation.SetRank(i)
	return muo
}

// AddRank adds i to the "rank" field.
func (muo *MovieUpdateOne) AddRank(i int) *MovieUpdateOne {
	muo.mutation.AddRank(i)
	return muo
}

// SetGenre sets the "genre" field.
func (muo *MovieUpdateOne) SetGenre(s string) *MovieUpdateOne {
	muo.mutation.SetGenre(s)
	return muo
}

// SetDirectorID sets the "director_id" field.
func (muo *MovieUpdateOne) SetDirectorID(i int) *MovieUpdateOne {
	muo.mutation.SetDirectorID(i)
	return muo
}

// SetNillableDirectorID sets the "director_id" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableDirectorID(i *int) *MovieUpdateOne {
	if i != nil {
		muo.SetDirectorID(*i)
	}
	return muo
}

// ClearDirectorID clears the value of the "director_id" field.
func (muo *MovieUpdateOne) ClearDirectorID() *MovieUpdateOne {
	muo.mutation.ClearDirectorID()
	return muo
}

// SetDirector sets the "director" edge to the Director entity.
func (muo *MovieUpdateOne) SetDirector(d *Director) *MovieUpdateOne {
	return muo.SetDirectorID(d.ID)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (muo *MovieUpdateOne) AddReviewIDs(ids ...int) *MovieUpdateOne {
	muo.mutation.AddReviewIDs(ids...)
	return muo
}

// AddReviews adds the "reviews" edges to the Review entity.
func (muo *MovieUpdateOne) AddReviews(r ...*Review) *MovieUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.AddReviewIDs(ids...)
}

// Mutation returns the MovieMutation object of the builder.
func (muo *MovieUpdateOne) Mutation() *MovieMutation {
	return muo.mutation
}

// ClearDirector clears the "director" edge to the Director entity.
func (muo *MovieUpdateOne) ClearDirector() *MovieUpdateOne {
	muo.mutation.ClearDirector()
	return muo
}

// ClearReviews clears all "reviews" edges to the Review entity.
func (muo *MovieUpdateOne) ClearReviews() *MovieUpdateOne {
	muo.mutation.ClearReviews()
	return muo
}

// RemoveReviewIDs removes the "reviews" edge to Review entities by IDs.
func (muo *MovieUpdateOne) RemoveReviewIDs(ids ...int) *MovieUpdateOne {
	muo.mutation.RemoveReviewIDs(ids...)
	return muo
}

// RemoveReviews removes "reviews" edges to Review entities.
func (muo *MovieUpdateOne) RemoveReviews(r ...*Review) *MovieUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.RemoveReviewIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MovieUpdateOne) Select(field string, fields ...string) *MovieUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Movie entity.
func (muo *MovieUpdateOne) Save(ctx context.Context) (*Movie, error) {
	var (
		err  error
		node *Movie
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MovieUpdateOne) SaveX(ctx context.Context) *Movie {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MovieUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MovieUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MovieUpdateOne) sqlSave(ctx context.Context) (_node *Movie, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   movie.Table,
			Columns: movie.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: movie.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Movie.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movie.FieldID)
		for _, f := range fields {
			if !movie.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != movie.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldTitle,
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldDescription,
		})
	}
	if value, ok := muo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: movie.FieldRank,
		})
	}
	if value, ok := muo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: movie.FieldRank,
		})
	}
	if value, ok := muo.mutation.Genre(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: movie.FieldGenre,
		})
	}
	if muo.mutation.DirectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   movie.DirectorTable,
			Columns: []string{movie.DirectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: director.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DirectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   movie.DirectorTable,
			Columns: []string{movie.DirectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: director.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !muo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   movie.ReviewsTable,
			Columns: []string{movie.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Movie{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
