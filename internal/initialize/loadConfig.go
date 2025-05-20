package initialize

import (
	"fmt"
	"go_ecommerce_backend/global"

	"github.com/spf13/viper"
)

func LoadCofig() {
	v := viper.New()

	// Đúng đường dẫn: từ cmd/vipper → ../config/
	v.AddConfigPath("config/")

	v.SetConfigName("local")
	v.SetConfigType("yaml")

	// In ra thư mục hiện tại để kiểm tra
	// dir, _ := os.Getwd()
	// fmt.Println("Working directory:", dir)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w\n", err))
	}

	fmt.Println("Server Port:", v.GetInt("server.port"))
	fmt.Println("JWT Key:", v.GetString("security.jwt.key"))

	//config structure

	if err := v.Unmarshal(&global.Config); err != nil { // ✅ Sử dụng v thay vì viper
		fmt.Println("Unable to decode configuration:", err)
	}
	// fmt.Println("Config port ::", &global.Config.Server.Port)

	// for _, db := range &global.Config.Database {
	// 	fmt.Printf("Database User: %s, password: %s, Username: %s \n", db.User, db.Password, db.Host)
	// }
}
