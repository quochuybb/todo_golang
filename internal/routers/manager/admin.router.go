package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *UserRouter) InitAdminRouter(Router *gin.RouterGroup) {
	userAdminRouterPublic := Router.Group("/admin")
	{
		userAdminRouterPublic.POST("/login")
	}
	userAdminRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Authen())
	//userRouterPrivate.Use(Permission())
	{
		userAdminRouterPrivate.POST("/active_user")
	}
}
