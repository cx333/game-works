package resource

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`    // 业务码，0表示成功
	Type    string      `json:"type"`    // 类型：success 或 error
	Message string      `json:"message"` // 消息提示
	Result  interface{} `json:"result"`  // 返回数据
}

// Success 成功：带数据
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Type:    "success",
		Message: "操作成功",
		Result:  data,
	})
}

// SuccessMsg 成功：带自定义消息
func SuccessMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Type:    "success",
		Message: msg,
		Result:  nil,
	})
}

// ErrBind 失败：请求绑定错误
func ErrBind(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    1,
		Type:    "error",
		Message: "参数绑定失败",
		Result:  nil,
	})
}

// Error 默认错误
func Error(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    1,
		Type:    "error",
		Message: "服务器内部错误",
		Result:  nil,
	})
}

// ErrorMsg 自定义错误消息
func ErrorMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    1,
		Type:    "error",
		Message: msg,
		Result:  nil,
	})
}

// ErrorCode 自定义业务码和消息
func ErrorCode(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Type:    "error",
		Message: msg,
		Result:  nil,
	})
}

// Fail 快捷失败
func Fail(ctx *gin.Context, msg string) {
	ErrorMsg(ctx, msg)
}

// PageData 分页数据结构体
type PageData struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// Page 分页响应
func Page(ctx *gin.Context, list interface{}, total int64, page, limit int) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Type:    "success",
		Message: "操作成功",
		Result: PageData{
			List:  list,
			Total: total,
			Page:  page,
			Limit: limit,
		},
	})
}
