package web

import (
	em "github.com/Etpmls/Etpmls-Micro/v2"
	"github.com/Etpmls/Etpmls-Micro/v2/define"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	KvServiceTemplate = "/template"
	KvServiceDefaultTemplate = "/default-template"
)

const (
	LayoutDirName = "layout"
	BaseTemplate = "app"
)



const (
	WebErrorCode = "400000"
	WebSuccessCode = "200000"
)

// Return error message in json format
// 返回json格式的错误信息
func JsonError(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{}, err error)  {
	s, _ := em.Kv.ReadKey(define.KvAppUseHttpCode)
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if strings.ToLower(s) == "true" {
		code = httpStatusCode
	}

	c.JSON(httpStatusCode, gin.H{"code": code, "error":"error", "message": message + "Error: " + err.Error(), "data": data})
	return
}


// Return success information in json format
// 返回json格式的成功信息
func JsonSuccess(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{})  {
	s, _ := em.Kv.ReadKey(define.KvAppUseHttpCode)
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if strings.ToLower(s) == "true" {
		code = httpStatusCode
	}

	c.JSON(httpStatusCode, gin.H{"code": code, "status":"success", "message": message, "data": data})
	return
}



