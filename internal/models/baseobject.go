package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// Use for querying a table
type Query struct{}

// Use for Updating a db object
type CreateParams struct{}

// use for updating objects
type UpdateParams struct{}

// use for deleting
type DeleteParams struct{}

// filter params

type BaseModel interface {
	getURI() string
}

type BaseDbService interface {
}

// Table Helpers
type Table struct {
	Name     string
	Columns  []string
	PkColumn string
}

type TableManager interface {
	// Get Operations
	getOne(*Query) []BaseModel
	getMultiple(*Query) []BaseModel

	// Creation Operations
	createOne(*CreateParams) BaseModel
	createMultiple(*[]BaseModel) []BaseModel

	// Update Operations
	updateOne(*UpdateParams) int
	updateMultiple(*UpdateParams) int

	// Delete Operations
	deleteOne(*DeleteParams) int
	deleteMultiple(*DeleteParams) int
}

func buildQuestionArray(num int) string {
	var qArray []string
	for i := 0; i < num; i++ {
		qArray[i] = "?"

	}
	qStr := strings.Join(qArray, ", ")
	return qStr
}

func BuildCreateStatment(table *Table, tm TableManager, db *sql.DB) (*sql.Stmt, error) {
	columnNames := strings.Join(table.Columns, ", ")

	qStr := buildQuestionArray(len(table.Columns))

	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table.Name, columnNames, qStr
	)
	return db.Prepare(queryString)

}
