package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)



func main() {
    // 连接 MySQL
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/your_database_name")
    if err!= nil {
        log.Fatal(err)
    }
    defer db.Close()

    r := gin.Default()

    // 获取内容的接口
    r.GET("/content", func(c *gin.Context) {
        rows, err := db.Query("SELECT id, data FROM your_table_name")
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        var contents []Content
        for rows.Next() {
            var content Content
            err := rows.Scan(&content.ID, &content.Data)
            if err!= nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            contents = append(contents, content)
        }

        c.JSON(http.StatusOK, contents)
    })

    fmt.Println("Server running on port 8080")
    r.Run(":8080")
}