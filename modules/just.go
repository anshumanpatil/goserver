package hello

import (
	"net/http"
	"github.com/gin-gonic/gin"
  )

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Api Working OK!!!"})    
}
