package main

import (
	hell "github.com/anshumanpatil/goserver-print-api"
	"github.com/gin-gonic/gin"
  )
  
func main () {
	r := gin.Default()
	r.GET("/", hell.SayHello)
	r.Run(":5656")
}