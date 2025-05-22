package initialize

import (
	"fmt"
	"go_ecommerce_backend/global"

	"go.uber.org/zap"
)

func Run() {
	LoadCofig()
	fmt.Println("Load configuration mysql:: ", global.Config.MySql.Username)
	fmt.Println("Load configuration mysql:: ", global.Config.MySql.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!", zap.String("Ok", "Success"))
	InitMysql()
	InitRedis()
	r := InitRouters()

	r.Run(":8002")
}
