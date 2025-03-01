package initialize

import (
	"todolist/global"
	"todolist/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
