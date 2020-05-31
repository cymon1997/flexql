package flexql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

const querySelect = "SELECT %s FROM %s WHERE %s"

func (s *DB) SelectContext(ctx context.Context, fields []*Field, tables []*Table, cds []*Condition) ([]*Result, error) {
	rows, err := s.db.QueryContext(ctx, buildQuery(querySelect, fields, tables, cds))
	if err != nil {
		return nil, err
	}
	var result []*Result
	for rows.Next() {
		temp, err := resultMap(rows, fields)
		if err != nil {
			return nil, err
		}
		result = append(result, NewResult(temp))
	}
	return result, nil
}

func buildQuery(query string, fields []*Field, tables []*Table, cds []*Condition) string {
	var field string
	field = fmt.Sprint(fields[0].StringAs())
	for _, f := range fields[1:] {
		field = fmt.Sprint(field, ", ", f.StringAs())
	}
	var table string
	table = fmt.Sprint(tables[0].Name)
	for _, t := range tables[1:] {
		table = fmt.Sprint(table, ", ", t.Name)
	}
	var cond string
	cond = fmt.Sprintf(cds[0].String())
	for _, c := range cds[1:] {
		cond = fmt.Sprint(cond, ", ", c.String())
	}
	log.Printf("debug query: %v", fmt.Sprintf(query, field, table, cond))
	return fmt.Sprintf(query, field, table, cond)
}
