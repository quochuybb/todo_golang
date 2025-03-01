package routers

import (
	"todolist/internal/routers/manager"
	"todolist/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
