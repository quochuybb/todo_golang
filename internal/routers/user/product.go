package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("product")
	{
		productRouter.GET("/product")
		productRouter.GET("/detail/:id")
	}
}
