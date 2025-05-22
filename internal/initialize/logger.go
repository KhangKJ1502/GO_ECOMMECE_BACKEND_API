package initialize

import (
	"go_ecommerce_backend/global"
	"go_ecommerce_backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
