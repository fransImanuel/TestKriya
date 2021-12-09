package user

import (
	"fmt"
	"kriya_Test/middleware"
	"kriya_Test/utilities/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

func User(route *gin.Engine) {
	v1 := route.Group("/user")

	v1.Use(middleware.CheckRole())

	v1.GET("/get/list/:page", func(c *gin.Context) {

		var param GetUserListParam

		if err := c.BindUri(&param); err != nil {
			fmt.Println(err)
			response := Response{
				Message:       "Failed",
				Error_key:     "internal_Server",
				Error_message: "internal_Server",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		db := db.Connect()
		defer db.Close()

		result, err := getListUser(db, param)
		if err != nil {
			fmt.Println(err)
			response := Response{
				Message:       "Failed",
				Error_key:     "internal_Server",
				Error_message: "internal_Server",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := Response{
			Message: "Success",
			Data:    result,
		}

		c.JSON(http.StatusOK, response)
	})

	v1.GET("/get/:userid", func(c *gin.Context) {
		var param GetUserParam

		if err := c.BindUri(&param); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		db := db.Connect()
		defer db.Close()

		result, err := getUser(db, param)
		if err != nil {
			response := Response{
				Message:       "Failed",
				Error_key:     "internal_Server",
				Error_message: "internal_Server",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := Response{
			Message: "Success",
			Data:    result,
		}

		c.JSON(http.StatusOK, response)

	})

	v1.POST("/post", func(c *gin.Context) {
		str, err := c.GetRawData()
		if err != nil {
			fmt.Println("error:", err)
			c.JSON(http.StatusBadRequest, err)
		}

		role_id := c.MustGet("id")
		body := strings.ToLower(string(str))

		u4, err := uuid.NewV4()
		if err != nil {
			fmt.Println("error:", err)
			c.JSON(http.StatusBadRequest, err)
		}

		db := db.Connect()
		defer db.Close()

		err = addUser(db, u4.String(), body, role_id.(string))
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
		}

		response := Response{
			Message: "Success",
		}

		c.JSON(http.StatusOK, response)

	})

	v1.PUT("/update/:userid", func(c *gin.Context) {
		var param GetUserParam

		c.BindUri(&param)

		str, err := c.GetRawData()
		if err != nil {
			fmt.Println("error:", err)
			c.JSON(http.StatusBadRequest, err)
		}

		body := strings.ToLower(string(str))

		db := db.Connect()
		defer db.Close()

		err = UpdateUser(db, body, param.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadGateway, err)
		}

		response := Response{
			Message: "Success",
		}

		c.JSON(http.StatusOK, response)

	})

	v1.DELETE("/delete/:userid", func(c *gin.Context) {
		var param GetUserParam

		c.BindUri(&param)

		db := db.Connect()
		defer db.Close()

		err := DeleteUser(db, param.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadGateway, err)
		}

		response := Response{
			Message: "Success",
		}

		c.JSON(http.StatusOK, response)
	})

}
