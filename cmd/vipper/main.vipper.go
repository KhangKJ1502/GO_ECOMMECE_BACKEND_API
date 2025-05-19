package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DBName   string `mapstructure:"dbName"` // ✅ Thêm vào đây
	} `mapstructure:"database"`
}

func main() {
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
	var config Config

	if err := v.Unmarshal(&config); err != nil { // ✅ Sử dụng v thay vì viper
		fmt.Println("Unable to decode configuration:", err)
	}
	fmt.Println("Config port ::", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("Database User: %s, password: %s, Username: %s \n", db.User, db.Password, db.Host)
	}
}
