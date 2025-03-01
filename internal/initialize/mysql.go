package initialize

import (
	"fmt"
	"time"
	"todolist/global"
	"todolist/internal/po"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanics(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)

	}

}

func InitMySQL() {
	//m := global.Config.MySQL
	dsn := "root:admin123@tcp(127.0.0.1:3307)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	//var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanics(err, "InitMySQL initialization error")
	global.Logger.Info("mysql connect success")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("SetPool error: %s", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLife))
}
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.Role{},
		&po.User{},
	)
	if err != nil {
		fmt.Println()
	}
}
