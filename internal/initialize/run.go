package initialize

import (
	"fmt"
	"go_ecommerce_backend/global"
)

func Run() {
	LoadCofig()
	fmt.Println("Load configuration mysql:: ", global.Config.MySql.Username)
	fmt.Println("Load configuration mysql:: ", global.Config.MySql.Password)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouters()

	r.Run(":8002")
}
