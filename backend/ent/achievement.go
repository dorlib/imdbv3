// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"imdbv2/ent/achievement"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Achievement is the model entity for the Achievement schema.
type Achievement struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Achievement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case achievement.FieldID:
			values[i] = new(sql.NullInt64)
		case achievement.FieldName, achievement.FieldImage, achievement.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Achievement", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Achievement fields.
func (a *Achievement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case achievement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case achievement.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case achievement.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				a.Image = value.String
			}
		case achievement.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Achievement.
// Note that you need to call Achievement.Unwrap() before calling this method if this Achievement
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Achievement) Update() *AchievementUpdateOne {
	return (&AchievementClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Achievement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Achievement) Unwrap() *Achievement {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Achievement is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Achievement) String() string {
	var builder strings.Builder
	builder.WriteString("Achievement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(a.Image)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Achievements is a parsable slice of Achievement.
type Achievements []*Achievement

func (a Achievements) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
