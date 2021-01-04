package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xiaoweihong/wolfweb/routers"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	group := r.Group("")
	{
		routers.InitSysInfo(group)
		routers.InitRedisRouter(group)
		routers.InitFseRouter(group)
	}
	r.Run(":9998")
}
