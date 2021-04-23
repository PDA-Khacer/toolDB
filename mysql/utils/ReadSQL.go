package utils

import (
	"database/sql"
)

type Cols struct {
	Name string
	TypeData string
	Key string
}

type Rows struct{
	Col map[string][]byte
}


type RowData struct {
	Col map[string]string
}


func GetAllSchemas(db *sql.DB) ([]string, error) {
	//if schema, err := db.Query("select schema_name as database_name from information_schema.schemata"); err != nil{
	if schema, err := db.Query("show schemas;"); err != nil{
		return nil, err
	} else {
		var re []string
		var token string
		for schema.Next() {
			err = schema.Scan(&token)
		}
		re = append(re, token)
		return re ,err
	}
}

func GetAllTables(db *sql.DB, schema string) ([]string, error) {
	_, _ = db.Query("use "+schema+";")

	if tables, err := db.Query("show tables;"); err != nil{
		return nil, err
	} else {
		var re []string
		var token string
		for tables.Next() {
			err = tables.Scan(&token)
		}
		re = append(re, token)
		return re ,err
	}
}

func GetAllDataRow(db *sql.DB, table string) ([]*Rows,error) {
	query := "select * from " + table
	if rows, err := db.Query(query); err != nil{
		return nil, err
	} else {
		// get all col name
		nameCols, err := rows.Columns()
		vals := make([]interface{}, len(nameCols))
		for i, _ := range nameCols {
			vals[i] = new(sql.RawBytes)
		}
		var result []*Rows
		for rows.Next() {
			err = rows.Scan(vals...)
			cols := map[string][]byte{}
			for i, _ := range nameCols {
				key := nameCols[i]
				cols[key] = vals[i].(sql.RawBytes)
			}
		}
		return result ,err
	}
}

func InsertToSQLDataRow(db *sql.DB, table string, rows Rows) error {
	var keys string
	var values string
	for  key, value := range rows.Col {
		keys += key + " ,"
		values += "'" + string(value) + "' ,"
	}
	sql := "INSERT INTO "+table+"("+keys[:len(keys)-1]+") VALUES ("+values[:len(values)-1]+")"
	_, err := db.Exec(sql)
	if err != nil{
		return err
	}
	return nil
}

func ToMapJson(table string, rows Rows) string {
	body := "{ "
	for  key, value := range rows.Col {
		body += " \n \t" + key + " : " + string(value) + ","
	}
	return body[:len(body) - 1] + "\n }"
}


