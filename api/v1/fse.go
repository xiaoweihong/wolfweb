package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoweihong/wolfweb/service"
	"strings"
)

func GetFseInfo(c *gin.Context) {
	address := c.Query("address")
	if strings.EqualFold(address, "") {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "address不能为空",
		})
		return
	}
	repos, err := service.GetRepos(address)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprintf(err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": repos,
	})
}

func DeleteRepoById(c *gin.Context) {
	id := c.Param("id")
	address := c.Query("address")
	result, err := service.DeleteRepoById(address, id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprintf(err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": result,
	})
}
