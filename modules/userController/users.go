package usercontroller

import (
	"log"
	"net/http"

	dbconnection "github.com/anshumanpatil/goserver-print-api/database"
	usermodel "github.com/anshumanpatil/goserver-print-api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type userParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AllUsers is - For generating login
func AllUsers(c *gin.Context) {
	log.Println("====== Get All Users ======")
	c.JSON(http.StatusOK, usermodel.GetAllUsers(dbconnection.DbCon))
}

// LoginResult is - For generating login
func LoginResult(c *gin.Context) {
	log.Println("====== Get All Users ======")
	var u userParams
	if err := c.ShouldBindWith(&u, binding.JSON); err == nil {
		result := usermodel.LoginResult(dbconnection.DbCon, u.Username, u.Password)
		if result != nil {
			c.JSON(http.StatusOK, gin.H{"error": nil, "data": result})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "User Not Found", "data": nil})
		}
	} else {
		log.Println(err)
	}
}
