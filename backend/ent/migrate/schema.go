// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DirectorsColumns holds the columns for the "directors" table.
	DirectorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// DirectorsTable holds the schema information for the "directors" table.
	DirectorsTable = &schema.Table{
		Name:       "directors",
		Columns:    DirectorsColumns,
		PrimaryKey: []*schema.Column{DirectorsColumns[0]},
	}
	// MoviesColumns holds the columns for the "movies" table.
	MoviesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "rank", Type: field.TypeInt},
		{Name: "genre", Type: field.TypeString},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "director_id", Type: field.TypeInt, Nullable: true},
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "movies_directors_movies",
				Columns:    []*schema.Column{MoviesColumns[6]},
				RefColumns: []*schema.Column{DirectorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "topic", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "rank", Type: field.TypeInt},
		{Name: "review_movie", Type: field.TypeInt, Nullable: true},
		{Name: "user_reviews", Type: field.TypeInt, Nullable: true},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_movies_movie",
				Columns:    []*schema.Column{ReviewsColumns[4]},
				RefColumns: []*schema.Column{MoviesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "reviews_users_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "firstname", Type: field.TypeString},
		{Name: "lastname", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "birth_day", Type: field.TypeString},
		{Name: "profile", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DirectorsTable,
		MoviesTable,
		ReviewsTable,
		UsersTable,
	}
)

func init() {
	MoviesTable.ForeignKeys[0].RefTable = DirectorsTable
	ReviewsTable.ForeignKeys[0].RefTable = MoviesTable
	ReviewsTable.ForeignKeys[1].RefTable = UsersTable
}
