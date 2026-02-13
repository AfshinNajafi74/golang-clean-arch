package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Header struct {
	UserId  string
	Browser string
}

type PersonData struct {
	FirstName string
	LastName  string
}

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

func (testHandler *TestHandler) HeaderBinder1(context *gin.Context) {
	userId := context.GetHeader("userId")

	context.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (testHandler *TestHandler) HeaderBinder2(context *gin.Context) {
	h := Header{}
	_ = context.BindHeader(&h)
	context.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": h,
	})
}

func (testHandler *TestHandler) QueryBinder1(context *gin.Context) {
	id := context.Query("id")
	name := context.Query("name")

	context.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	})
}

func (testHandler *TestHandler) QueryBinder2(context *gin.Context) {
	ids := context.QueryArray("id")
	name := context.Query("name")
	context.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id":     ids,
		"name":   name,
	})
}

func (testHandler *TestHandler) UriBinder(context *gin.Context) {
	id := context.Param("id")
	name := context.Param("name")

	context.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (testHandler *TestHandler) BodyBinder(context *gin.Context) {
	p := PersonData{}
	context.BindJSON(&p)
	context.JSON(http.StatusOK, gin.H{
		"result":    "UriBinder",
		"firstName": p.FirstName,
		"lastName":  p.LastName,
		"p":         p,
	})
}

func (testHandler *TestHandler) FormBinder(context *gin.Context) {
	p := PersonData{}
	context.Bind(&p)
	context.JSON(http.StatusOK, gin.H{
		"result":    "FormBinder",
		"firstName": p.FirstName,
		"lastName":  p.LastName,
		"p":         p,
	})
}

func (testHandler *TestHandler) FileBinder(context *gin.Context) {
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, "file")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	})
}
