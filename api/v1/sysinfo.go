package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoweihong/wolfweb/service"
)

func GetSysInfo(c *gin.Context) {
	info, err := service.GetSysInfo()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprintf(err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": info,
	})
}
