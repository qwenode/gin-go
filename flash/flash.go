package flash

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	labelSuccess = "_SUCCESS"
	labelError   = "_ERROR"
	labelInfo    = "_INFO"
)

func Assert(c *gin.Context, err error) bool {
	if err != nil {
		Error(c, err.Error())
		return false
	}
	return true
}
func Success(c *gin.Context, message string) {
	setMessage(c, labelSuccess, message)
}

func Error(c *gin.Context, message string) {
	setMessage(c, labelError, message)
}

func Info(c *gin.Context, message string) {
	setMessage(c, labelInfo, message)
}
func GetSuccessMessage(c *gin.Context) string {
	return getMessage(c, labelSuccess)
}
func GetErrorMessage(c *gin.Context) string {
	return getMessage(c, labelError)
}
func GetInfoMessage(c *gin.Context) string {
	return getMessage(c, labelInfo)
}

func getMessage(c *gin.Context, t string) string {
	session := sessions.Default(c)
	m := session.Get(t)
	session.Delete(t)
	session.Save()
	if m == nil {
		return ""
	}
	return m.(string)
}
func setMessage(c *gin.Context, t string, message string) {
	session := sessions.Default(c)
	session.Set(t, message)
	session.Save()
}
