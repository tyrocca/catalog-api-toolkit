package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// Table Helpers
type TableColumns []string

type Table struct {
	Name     string
	Columns  TableColumns
	PkColumn string
}

type TableManager interface {
	// Get Operations
	GetOne(*Query) []BaseModel
	GetMultiple(*Query) []BaseModel

	// Creation Operations
	CreateOne(*CreateParams) BaseModel
	CreateMultiple(*[]BaseModel) []BaseModel

	// Update Operations
	UpdateOne(*UpdateParams) int
	UpdateMultiple(*UpdateParams) int

	// Delete Operations
	DeleteOne(*DeleteParams) int
	DeleteMultiple(*DeleteParams) int

	// helper methods
	columnExists(col string) bool
	validColumns(tbCols *TableColumns, cols *TableColumns) (bool, error)

	// Query Helpers -- should these be privare
	BuildCreateStatment(db *sql.DB) (*sql.Stmt, error)
	BuildUpdateStatement(db *sql.DB) (*sql.Stmt, error)
	BuildFilterStatement(db *sql.DB) (*sql.Stmt, error)
}

/*
 * General Helpers
 */

func buildQuestionArray(num int) string {
	var qArray []string
	for i := 0; i < num; i++ {
		qArray[i] = "?"
	}
	qStr := strings.Join(qArray, ", ")
	return qStr
}

func buildSetClause(cols *TableColumns) string {
	outstr := ""
	for _, col := range *cols {
		outstr = fmt.Sprintf("%s%s=?,", outstr, col)
	}
	return outstr
}

func buildANDClause(cols *TableColumns) string {
	var filterArray []string
	for i, col := range *cols {
		filterArray[i] = fmt.Sprintf("%s=?", col)
	}
	return "(" + strings.Join(filterArray, " AND ") + ")"
}

func buildOrderByClause(cols *TableColumns) string {
	if cols != nil && len(*cols) > 0 {
		joined := strings.Join(*cols, ", ")
		return fmt.Sprintf(" ORDER BY (%s)", joined)
	}
	return ""
}

func buildLimitClause(limit int) string {
	if limit > 0 {
		return fmt.Sprintf(" LIMIT %d", limit)
	}
	return ""
}

/*
 * Table Helpers
 */

func (t *Table) columnExists(col string) bool {
	for _, tblCol := range t.Columns {
		if col == tblCol {
			return true
		}
	}
	return false
}

func (t *Table) validColumns(cols *TableColumns) bool {
	for _, col := range *cols {
		if !t.columnExists(col) {
			return false
		}
	}
	return true

}

/*
 * Create Columns
 */

func (t *Table) BuildCreateStatment(db *sql.DB) (*sql.Stmt, error) {
	columnNames := strings.Join(t.Columns, ", ")

	qStr := buildQuestionArray(len(t.Columns))

	queryString := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		t.Name, columnNames, qStr,
	)
	return db.Prepare(queryString)
}

/*
 * Update Helpers
 */

func (t *Table) BuildUpdateStatement(
	db *sql.DB,
	filterCols *TableColumns,
	updateCols *TableColumns,
) (*sql.Stmt, error) {

	// TODO (tyrocca 2022-03-28): handle invalid column
	t.validColumns(updateCols)
	t.validColumns(filterCols)

	// TODO (tyrocca 2022-03-28): Build update
	updateString := buildSetClause(updateCols)
	filterString := buildANDClause(filterCols)

	queryString := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		t.Name, updateString, filterString,
	)
	return db.Prepare(queryString)
}

/*
 * Filter helpers
 */

func (t *Table) BuildFilterStatement(
	db *sql.DB,
	filterCols *TableColumns,
	selectCols *TableColumns,
	orderByCols *TableColumns,
	limit int,
) (*sql.Stmt, error) {

	// TODO (tyrocca 2022-03-28): error handle
	t.validColumns(filterCols)
	t.validColumns(selectCols)
	t.validColumns(orderByCols)

	selectString := strings.Join(*selectCols, ", ")
	filterString := buildANDClause(filterCols)

	// Additional params
	orderByClause := buildOrderByClause(orderByCols)
	limitClause := buildLimitClause(limit)

	queryString := fmt.Sprintf("SELECT %s FROM %s WHERE %s%s%s",
		selectString, t.Name, filterString, orderByClause, limitClause,
	)
	return db.Prepare(queryString)
}
