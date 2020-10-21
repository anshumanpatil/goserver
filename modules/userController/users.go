package usercontroller

import (
	"log"
	"net/http"

	dbconnection "github.com/anshumanpatil/goserver-print-api/database"
	helperfunctions "github.com/anshumanpatil/goserver-print-api/helper"
	usermodel "github.com/anshumanpatil/goserver-print-api/models"
	"github.com/gin-gonic/gin"
)

type userParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AllUsers is - For generating login
func AllUsers(c *gin.Context) {
	log.Println("====== Get All Users ======")
	log.Println("================================================================================================")
	db := c.MustGet("DB")
	log.Print(db)
	log.Println("================================================================================================")
	c.JSON(http.StatusOK, usermodel.GetAllUsers(dbconnection.DbCon))
}

// LoginResult is - For generating login
func LoginResult(c *gin.Context) {
	log.Println("====== Get Login Status of User ======")
	var u userParams
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := usermodel.LoginResult(dbconnection.DbCon, u.Username, u.Password)
	if result != nil {
		if token, err := helperfunctions.GenerateToken(result[0].ID); err == nil {
			c.JSON(http.StatusOK, gin.H{"error": nil, "data": result, "token": token})
		}
	} else {
		c.JSON(http.StatusLocked, gin.H{"error": "User Not Found", "data": nil})
	}
}
