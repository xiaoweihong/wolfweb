package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaoweihong/wolfweb/api/v1"
)

func InitFseRouter(r *gin.RouterGroup) {
	fseGroup := r.Group("fse")
	{
		fseGroup.GET("info", v1.GetFseInfo)
		fseGroup.DELETE("repo/:id", v1.DeleteRepoById)
	}
}
