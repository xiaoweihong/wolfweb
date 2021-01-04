package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaoweihong/wolfweb/api/v1"
)

func InitSysInfo(r *gin.RouterGroup)  {
	sysinfoGroup := r.Group("sys")
	{
		sysinfoGroup.GET("/info",v1.GetSysInfo)
	}
}
