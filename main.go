package main

import (
	"aashirwad/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"titlde": "trils only",
		})
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
		})
	})
	router.GET("/kitchen", func(c *gin.Context) {
		c.HTML(http.StatusOK, "kitchen.html", gin.H{
			"title": "kitchen",
		})
	})
	router.GET("/manager", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manager.html", gin.H{
			"title": "manager",
		})
	})
	router.GET("/team", func(c *gin.Context) {
		c.HTML(http.StatusOK, "team.html", nil)
	})
	router.POST("/open", func(c *gin.Context) {
		file := c.PostForm("file")

		var templateName string
		switch file {
		case "admin":
			templateName = "admin.html"
		case "kitchen":
			templateName = "kitchen.html"
		case "manager":
			templateName = "manager.html"
		case "waiter":
			templateName = "waiter.html"
		default:
			templateName = "team.html"
		}

		c.HTML(http.StatusOK, templateName, nil)
	})
	// router.GET("/waiter", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "waiter.html", gin.H{
	// 		"title": "waiter",
	// 	})
	// })

	router.GET("/waiter", controllers.ShowOrderitem)
	router.POST("/submit", controllers.SubmitItem)
	// router.POST("/submit", func(c *gin.Context) {
	// 	var item Item
	// 	if err := c.BindJSON(&item); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid input"})
	// 		return
	// 	}
	// 	_, err = db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error inserting data"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"success": true})
	// })

	// return router
	router.Run(":9090")

}
