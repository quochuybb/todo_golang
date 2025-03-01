package initialize

import (
	"todolist/global"
	"todolist/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	managerRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User
	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitUserAdminRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}

	return r
}
