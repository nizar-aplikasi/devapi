package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data,omitempty"`
}

func Success(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: "operation success",
		Data:    data,
	}
}

func Error(message string) APIResponse {
	return APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
	}
}

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, Success(data))
}

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, Error(message))
}
