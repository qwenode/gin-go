package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CODE_SUCCESS = 0
	CODE_ERROR   = 1
)

func SuccessWithList(data interface{}, total int, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_SUCCESS,
		"data": data,
	})
}

func SuccessWithData(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_SUCCESS,
		"data": data,
	})
}

func SuccessWithMessage(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_SUCCESS,
		"msg":  msg,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_SUCCESS,
		"msg":  "success",
	})
}

func ErrorWithMessage(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": CODE_ERROR,
		"msg":  msg,
	})
}

func MessageWithCode(msg string, code int, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
func ErrorWithAuthFail(msg string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  msg,
	})
}

func Refresh(c *gin.Context) {
	Redirect302(c, c.Request.URL.String())
}

func Redirect302(c *gin.Context, url string) {
	c.Redirect(302, url)
}
func Redirect301(c *gin.Context, url string) {
	c.Redirect(301, url)
}
