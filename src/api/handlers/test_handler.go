package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (testHandler *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (testHandler *TestHandler) Users(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (testHandler *TestHandler) UserById(context *gin.Context) {
	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (testHandler *TestHandler) GetByUserName(context *gin.Context) {
	userName := context.Param("username")
	context.JSON(http.StatusOK, gin.H{
		"result": "GetByUserName",
		"user":   userName,
	})
}

func (testHandler *TestHandler) Accounts(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     context.Param("id"),
	})
}

func (testHandler *TestHandler) AddUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
		"id":     context.Param("id"),
	})
}
