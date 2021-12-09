package main

import (
	"kriya_Test/routes/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Kriya Test")
	})

	user.User(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
