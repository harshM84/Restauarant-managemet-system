// controllers/customer_controller.go

package controllers

import (
	"aashirwad/db"
	"aashirwad/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// ShowUsersPage renders the users page with a list of users
func ShowOrderitem(c *gin.Context) {
	database := db.ConnectDB()
	defer database.Close()

	orderitems, err := models.Getorderitem(database)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "waiter.html", gin.H{"orderitems": orderitems})
}

func SubmitItem(c *gin.Context) {
	database := db.ConnectDB()
	defer database.Close()

	c.JSON(http.StatusOK, gin.H{"success": true})
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid input"})
		return
	}
}
