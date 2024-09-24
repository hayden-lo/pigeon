package mysqlutil

import (
    "database/sql"
    "fmt"
    "log" 
      "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/your_database_name")
    if err!= nil {
        log.Fatal(err)
        return nil, err
    }
    return db, nil
}