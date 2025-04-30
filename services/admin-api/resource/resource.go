package resource

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
	Code int         `json:"code"` // 业务码，200成功，其他失败
	Data interface{} `json:"data"` // 数据部分，可为对象、列表、分页等
	Msg  string      `json:"msg"`  // 提示信息
}

// Success 成功：带数据
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Resource{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// SuccessMsg 成功：自定义消息
func SuccessMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Resource{
		Code: 200,
		Data: nil,
		Msg:  msg,
	})
}

// ErrBind 绑定数据错误
func ErrBind(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, Resource{
		Code: 400,
		Data: nil,
		Msg:  "failed",
	})
}

// Error 错误：默认错误
func Error(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, Resource{
		Code: 500,
		Data: nil,
		Msg:  "error",
	})
}

// ErrorMsg 错误：自定义消息
func ErrorMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, Resource{
		Code: 500,
		Data: nil,
		Msg:  msg,
	})
}

// ErrorCode 错误：自定义错误码和消息
func ErrorCode(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Resource{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}

// Fail 快捷：失败
func Fail(ctx *gin.Context, msg string) {
	ErrorMsg(ctx, msg)
}

// PageData 分页结构体
type PageData struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// Page 分页响应
func Page(ctx *gin.Context, list interface{}, total int64, page, limit int) {
	ctx.JSON(http.StatusOK, Resource{
		Code: 200,
		Data: PageData{
			List:  list,
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Msg: "success",
	})
}
