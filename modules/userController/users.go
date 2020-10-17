package userController

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
  )

func AllUsers(c *gin.Context) {
	log.Println("====== Get All Users ======")
	c.JSON(http.StatusOK, gin.H{"users": "All"})    
}