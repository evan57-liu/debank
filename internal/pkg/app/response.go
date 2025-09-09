package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiResponse 定义 API 响应的通用格式
type ApiResponse struct {
	Message string      `json:"message"`        // 描述信息
	Data    interface{} `json:"data,omitempty"` // 具体数据
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ApiResponse{
		Message: "success",
		Data:    data,
	})
}

// Error 返回错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, &ApiResponse{
		Message: message,
		Data:    nil,
	})
}

func ErrorWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, &ApiResponse{
		Message: message,
		Data:    data,
	})
}
