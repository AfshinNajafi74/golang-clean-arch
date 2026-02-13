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

	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)

	r.POST("/binder/query1", h.QueryBinder1)
	r.POST("/binder/query2", h.QueryBinder2)

	r.POST("/binder/uri/:id/:name", h.UriBinder)

	r.POST("/binder/body", h.BodyBinder)

	r.POST("/binder/form", h.FormBinder)

	r.POST("/binder/file", h.FileBinder)
}
