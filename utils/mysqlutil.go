package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	host := GlobalConfig.Mysql.Host
	port := GlobalConfig.Mysql.Port
	username := GlobalConfig.Mysql.Username
	password := GlobalConfig.Mysql.Password
	databaseName := GlobalConfig.Mysql.DataBase

	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+databaseName)
	if err != nil {
		log.Fatalf("Create connector failed: %v", err)
	}

	// setup connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
}

func Select(query string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		log.Fatal("Database not connect!")
	}
	return db.Query(query, args...)
}

func Upsert(table string, values []interface{}, updateColumns []string) (int64, error) {
	if db == nil {
		log.Fatal("Database not connect!")
	}

	// make value string
	valueHolders := make([]string, len(values))
	for i := range values {
		valueHolders[i] = "?"
	}
	valueStr := strings.Join(valueHolders, ",")

	// make update column string
	updateColHolders := make([]string, len(updateColumns))
	for i := range updateColumns {
		updateColHolders[i] = updateColumns[i] + "(" + updateColumns[i] + ")"
	}
	updateColStr := strings.Join(valueHolders, ",")

	// execute upsert query
	upsertQuery := fmt.Sprintf("insert into %s values(%s) on duplicate key update %s", table, valueStr, updateColStr)
	result, err := db.Exec(upsertQuery, values...)
	if err != nil {
		log.Fatalf("Upsert failed: %v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Get affected rows failed: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}

func Insert(table string, args ...interface{}) (int64, error) {
	if db == nil {
		log.Fatal("Database not connect!")
	}

	placeholders := strings.Join(strings.Split(strings.Repeat("?", len(args)), ""), ",")
	insertQuery := fmt.Sprintf("insert into %s values(%s)", table, placeholders)
	result, err := db.Exec(insertQuery, args...)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
		return 0, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Get insert id failed：%v", err)
		return 0, err
	}
	return insertedID, nil
}

func BulkUpsert(table string, values [][]interface{}, updateColumns []string) (int64, error) {
	if db == nil {
		log.Fatal("Database not connect!")
	}

	lineHolders := make([]string, len(values[0]))
	for i := range values[0] {
		lineHolders[i] = "?"
	}
	lineStr := strings.Join(lineHolders, ",")

	valueHolders := make([]string, len(values))
	var args []interface{}
	for i, row := range values {
		valueHolders[i] = "(" + lineStr + ")"
		args = append(args, row...)
	}
	valueStr := strings.Join(valueHolders, ",")

	updateColHolders := make([]string, len(updateColumns))
	for i := range len(updateColumns) {
		updateColHolders[i] = updateColumns[i] + "=values(" + updateColumns[i] + ")"
	}
	updateColStr := strings.Join(updateColHolders, ",")

	upsertQuery := fmt.Sprintf("insert into %s values %s on duplicate key update %s", table, valueStr, updateColStr)
	result, err := db.Exec(upsertQuery, args...)
	if err != nil {
		log.Fatalf("Upsert failed: %v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("Get affected rows failed: %v", err)
		return 0, err
	}
	return rowsAffected, nil
}
