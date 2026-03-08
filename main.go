package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/acomma/admin/handler"
	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "123456"
	config.Net = "tcp"
	config.Addr = "127.0.0.1:3306"
	config.DBName = "admin-go"

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
