package main

import (
	hello "github.com/anshumanpatil/goserver-print-api"
	dbConnection "github.com/anshumanpatil/goserver-print-api/database"
	userController "github.com/anshumanpatil/goserver-print-api/userController"
	"github.com/gin-gonic/gin"
  )
  
func main () {
	dbConnection.DbConnect()
	r := gin.Default()
	r.GET("/", hello.SayHello)
	r.GET("/login", userController.LoginResult)
	r.GET("/getAllUsers", userController.AllUsers)
	r.Run(":5656")
}