package query

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type rowValue struct {
	value interface{}
}

func newRowValue(value interface{}) *rowValue {
	v := value.(*interface{})
	data := *v

	return &rowValue{data}
}

// AsString returns the value as string
func (r *rowValue) AsString() *sql.NullString {
	result := &sql.NullString{}

	if r != nil && r.value != nil {
		result.Scan(r.value)
	}

	return result
}

// AsTime returns the value as string
func (r *rowValue) AsTime() *mysql.NullTime {
	result := &mysql.NullTime{}

	if r != nil && r.value != nil {
		result.Scan(r.value)
	}

	return result
}

// AsInt64 returns the value as string
func (r *rowValue) AsInt64() *sql.NullInt64 {
	result := &sql.NullInt64{}

	if r != nil && r.value != nil {
		result.Scan(r.value)
	}

	return result
}

// AsBool returns the value as string
func (r *rowValue) AsBool() *sql.NullBool {
	result := &sql.NullBool{}

	if r != nil && r.value != nil {
		result.Scan(r.value)
	}

	return result
}
