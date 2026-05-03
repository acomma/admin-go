package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/acomma/admin/handler"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.NewConfig()
	config.User = GetEnv("MYSQL_USER", "root")
	config.Passwd = GetEnv("MYSQL_PASSWD", "123456")
	config.Net = "tcp"
	config.Addr = GetEnv("MYSQL_ADDR", "127.0.0.1:3306")
	config.DBName = GetEnv("MYSQL_DBNAME", "admin-go")

	database, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("打开数据库失败：%v\n", err)
	}

	userHandler := handler.NewUserHandler(database)
	roleHandler := handler.NewRoleHandler(database)

	http.HandleFunc("GET /users/{userId}", userHandler.GetUserByUserId)
	http.HandleFunc("GET /roles/{roleId}", roleHandler.GetRoleByRoleId)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动服务失败：%v\n", err)
	}
}

func GetEnv(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
