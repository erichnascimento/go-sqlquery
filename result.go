package query

import "database/sql"

// Row is a row structure
type Row map[string]*rowValue

// Result is result of a query
type Result struct {
	*sql.Rows
	columns []string
}

func (qr *Result) Read() (Row, error) {
	numColumns := len(qr.columns)

	row := make([]interface{}, numColumns)
	rowPointer := make([]interface{}, numColumns)
	for i := range row {
		row[i] = ""
		rowPointer[i] = &row[i]
	}

	err := qr.Scan(rowPointer...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*rowValue, numColumns)
	for i, name := range qr.columns {
		result[name] = newRowValue(rowPointer[i])
	}

	return result, nil
}

// RowToResult retur
func RowToResult(rows *sql.Rows) (*Result, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	return &Result{rows, columns}, err
}

// Query performs a query and return multiples rows
func Query(conn *sql.DB, query string, args ...interface{}) (*Result, error) {
	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return RowToResult(rows)
}

// QueryStmt performs a query based on a Statement
func QueryStmt(stmt *sql.Stmt, args ...interface{}) (*Result, error) {
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return RowToResult(rows)
}
