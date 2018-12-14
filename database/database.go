package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDatabase is result database
func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/s.a.d")
	if err != nil {
		fmt.Println("error from connect", err)
		return nil, err
	}

	return db, nil
}

// QueryString is function result database
func QueryString(queryString string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("error from querystring", err)
		return nil, err
	}
	return rows, nil
}

func ResultValue(rows *sql.Rows) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		item := make(map[string]interface{})
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			item[columns[i]] = value
		}

		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
