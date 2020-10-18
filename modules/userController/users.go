package userController

import (
	"log"
	"net/http"

	dbConnection "github.com/anshumanpatil/goserver-print-api/database"
	userModel "github.com/anshumanpatil/goserver-print-api/models"
	"github.com/gin-gonic/gin"
)

type List struct {
	Messages []string `key:"required"`
}

func AllUsers(c *gin.Context) {
	log.Println("====== Get All Users ======")
	c.JSON(http.StatusOK, userModel.GetAllUsers(dbConnection.DbCon))
}

func LoginResult(c *gin.Context) {
	log.Println("====== Get All Users ======")
	NrawBody, _ := c.GetRawData()
	obj := c.BindJSON(NrawBody)
	log.Println(NrawBody.foo)
	c.JSON(http.StatusOK, userModel.LoginResult(dbConnection.DbCon, "anshuman", "ansh"))
}
