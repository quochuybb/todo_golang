package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"todolist/global"
)

func Test() {
	LoadConfig()
	fmt.Println("Load Config Success", global.Config.MySQL.Host)
	InitLogger()
	global.Logger.Info("initialize success", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":4040")
}
