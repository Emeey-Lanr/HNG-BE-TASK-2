package helpers

import (
	"be-task2/models"

	"github.com/gin-gonic/gin"
)

func ErrorResponse (method int, data models.ErrorResp,   c *gin.Context){
 c.JSON(method,  data)
}

func SuccessResponse (method int, data gin.H, c *gin.Context){
  c.JSON(method, data)
}