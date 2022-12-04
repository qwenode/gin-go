package reqs

import (
	"github.com/gin-gonic/gin"
	"github.com/qwenode/gogo/convert"
)

var (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

func IsPOST(c *gin.Context) bool {
	return c.Request.Method == METHOD_POST
}
func IsGET(c *gin.Context) bool {
	return c.Request.Method == METHOD_GET
}
func IsAjax(c *gin.Context) bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}
func GetPage(c *gin.Context) int {
	query, b := c.GetQuery("page")
	if !b {
		return 1
	}
	toInt := convert.ToInt(query)
	if toInt > 0 {
		return toInt
	}
	return 1
}
