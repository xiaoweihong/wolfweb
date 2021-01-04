package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaoweihong/wolfweb/api/v1"
)

func InitRedisRouter(r *gin.RouterGroup)  {
	redisGroup := r.Group("redis")
	{
		redisGroup.GET("info",v1.GetRedisInfo)
		redisGroup.DELETE("delete/:id",v1.DeleteRedisDbById)
		redisGroup.DELETE("deleteall",v1.DeleteRedisDbAll)
	}
}
