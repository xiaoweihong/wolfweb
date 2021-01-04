package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redisgo "github.com/xiaoweihong/wolfweb/utils/redis"
	"strconv"
)

var cacher *redisgo.Cacher

func GetRedisInfo(c *gin.Context) {
	addr := c.Query("address")
	password := c.Query("password")
	cacher, _ = redisgo.New(redisgo.Options{
		Addr:     addr,
		Password: password,
		Db:       0,
	})
	dbs, err := cacher.GetAllDb()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprint(err),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "",
		"data":    dbs,
	})
}

func DeleteRedisDbById(c *gin.Context) {
	dbId := c.Param("id")
	dId, _ := strconv.Atoi(dbId)

	err := cacher.Flush(dId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprint(err),
		})
		return
	}
}

func DeleteRedisDbAll(c *gin.Context) {
	err := cacher.FlushALL()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprint(err),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "",
	})
}
