package main

import (
	hello "github.com/anshumanpatil/goserver-print-api"
	"github.com/gin-gonic/gin"
  )
  
func main () {
	r := gin.Default()
	r.GET("/", hello.SayHello)
	r.Run(":5656")
}