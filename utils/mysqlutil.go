package utils

import (
	"database/sql"
	"log"
	"strconv"

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
		log.Fatalf("创建 Connector 失败：%v", err)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(10) // 设置最大打开连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
}

func Select(query string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		log.Fatal("数据库连接未初始化")
	}
	return db.Query(query, args...)
}

func Update(query string, args ...interface{}) (int64, error) {
	if db == nil {
		log.Fatal("数据库连接未初始化")
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalf("更新失败：%v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("获取受影响行数失败：%v", err)
		return 0, err
	}
	return rowsAffected, nil
}

func Insert(query string, args ...interface{}) (int64, error) {
	if db == nil {
		log.Fatal("数据库连接未初始化")
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalf("插入失败：%v", err)
		return 0, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("获取插入的 ID 失败：%v", err)
		return 0, err
	}
	return insertedID, nil
}

// func Connect() error {
// 	host := GlobalConfig.Mysql.Host
// 	port := GlobalConfig.Mysql.Port
// 	username := GlobalConfig.Mysql.Username
// 	password := GlobalConfig.Mysql.Password
// 	databaseName := GlobalConfig.Mysql.DataBase

// 	var err error
// 	db, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+databaseName)
// 	if err != nil {
// 		log.Fatalf("创建 Connector 失败：%v", err)
// 	}
// 	return err
// }

// func Select(query string, args ...interface{}) (*sql.Rows, error) {
// 	if db == nil {
// 		if err := Connect(); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return db.Query(query, args...)
// }

// func Update(query string, args ...interface{}) (int64, error) {
// 	if db == nil {
// 		if err := Connect(); err != nil {
// 			return 0, err
// 		}
// 	}
// 	result, err := db.Exec(query, args...)
// 	if err != nil {
// 		log.Fatalf("更新失败：%v", err)
// 		return 0, err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatalf("获取受影响行数失败：%v", err)
// 		return 0, err
// 	}
// 	return rowsAffected, nil
// }

// func Insert(query string, args ...interface{}) (int64, error) {
// 	if db == nil {
// 		if err := Connect(); err != nil {
// 			return 0, err
// 		}
// 	}
// 	result, err := db.Exec(query, args...)
// 	if err != nil {
// 		log.Fatalf("插入失败：%v", err)
// 		return 0, err
// 	}
// 	insertedID, err := result.LastInsertId()
// 	if err != nil {
// 		log.Fatalf("获取插入的 ID 失败：%v", err)
// 		return 0, err
// 	}
// 	return insertedID, nil
// }
