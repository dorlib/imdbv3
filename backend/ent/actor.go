// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"imdbv2/ent/actor"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Actor is the model entity for the Actor schema.
type Actor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CharacterName holds the value of the "character_name" field.
	CharacterName string `json:"character_name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActorQuery when eager-loading is set.
	Edges ActorEdges `json:"edges"`
}

// ActorEdges holds the relations/edges for other nodes in the graph.
type ActorEdges struct {
	// Actors holds the value of the actors edge.
	Actors []*Movie `json:"actors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ActorsOrErr returns the Actors value or an error if the edge
// was not loaded in eager-loading.
func (e ActorEdges) ActorsOrErr() ([]*Movie, error) {
	if e.loadedTypes[0] {
		return e.Actors, nil
	}
	return nil, &NotLoadedError{edge: "actors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Actor) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case actor.FieldID:
			values[i] = new(sql.NullInt64)
		case actor.FieldName, actor.FieldCharacterName, actor.FieldImage:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Actor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Actor fields.
func (a *Actor) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case actor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case actor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case actor.FieldCharacterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field character_name", values[i])
			} else if value.Valid {
				a.CharacterName = value.String
			}
		case actor.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				a.Image = value.String
			}
		}
	}
	return nil
}

// QueryActors queries the "actors" edge of the Actor entity.
func (a *Actor) QueryActors() *MovieQuery {
	return (&ActorClient{config: a.config}).QueryActors(a)
}

// Update returns a builder for updating this Actor.
// Note that you need to call Actor.Unwrap() before calling this method if this Actor
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Actor) Update() *ActorUpdateOne {
	return (&ActorClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Actor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Actor) Unwrap() *Actor {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Actor is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Actor) String() string {
	var builder strings.Builder
	builder.WriteString("Actor(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", character_name=")
	builder.WriteString(a.CharacterName)
	builder.WriteString(", image=")
	builder.WriteString(a.Image)
	builder.WriteByte(')')
	return builder.String()
}

// Actors is a parsable slice of Actor.
type Actors []*Actor

func (a Actors) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
