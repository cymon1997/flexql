package flexql

import (
	"database/sql"
)

func resultMap(rows *sql.Rows, fields []*Field) (map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	temp := make([]interface{}, len(cols))
	for i := range temp {
		switch fields[i].Type {
		case TypeInt:
			temp[i] = new(int)
		default:
			temp[i] = new(string)
		}
	}
	err = rows.Scan(temp...)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	for i, data := range temp {
		result[cols[i]] = data
	}
	return result, nil
}
