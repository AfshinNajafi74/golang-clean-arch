package routers

import (
	"golang-clean-arch/api/handlers"

	"github.com/gin-gonic/gin"
)

func TestHandler(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/users/:id", h.UserById)
	r.GET("/users/get-user-by-username/:username", h.GetByUserName)
	r.GET("/users/:id/accounts", h.Accounts)

	r.POST("/add-user", h.AddUser)
}
