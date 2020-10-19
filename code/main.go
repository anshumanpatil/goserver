package main

import (
	hello "github.com/anshumanpatil/goserver-print-api"
	dbconnection "github.com/anshumanpatil/goserver-print-api/database"
	usercontroller "github.com/anshumanpatil/goserver-print-api/userController"
	"github.com/gin-gonic/gin"
)

func main() {
	dbconnection.DbConnect()
	r := gin.Default()
	r.GET("/", hello.SayHello)
	r.POST("/login", usercontroller.LoginResult)
	r.POST("/getAllUsers", usercontroller.AllUsers)
	r.Run(":5656")
}
