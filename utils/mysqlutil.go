package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() error {
	username := "xmiles"
	password := "oR%EwEgdFNv9io8L"
	host := "pc-uf60e09n0v9968tta00.rwlb.rds.aliyuncs.com"
	port := "3306"
	databaseName := "meteor_ssp"
	var err error
	// db, err := sql.Open("mysql", "xmiles:oR%EwEgdFNv9io8L@tcp(jdbc:mysql://pc-uf60e09n0v9968tta.rwlb.rds.aliyuncs.com:3306)/meteor_ssp?useUnicode=true&characterEncoding=UTF8&useAffectedRows=true&useSSL=false&serverTimezone=Asia/Shanghai&autoReconnect=true")
	db, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+databaseName)
	if err != nil {
		log.Fatalf("创建 Connector 失败：%v", err)
	}
	return err
}

func Select(query string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		if err := Connect(); err != nil {
			return nil, err
		}
	}
	return db.Query(query, args...)
}

func Update(query string, args ...interface{}) (int64, error) {
	if db == nil {
		if err := Connect(); err != nil {
			return 0, err
		}
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
