package initialize

import (
	"fmt"
	"go_ecommerce_backend/global"
	"go_ecommerce_backend/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySql

	// DSN format: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "InitMysql initialization error")
	global.Logger.Info("InitMysql initialization success")
	global.Mdb = db

	//set pool
	SetPool()

	migrateTable()
}

// InitMysql.SetPool()
func SetPool() {
	m := global.Config.MySql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error: %s ::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns) //giới hạn số lượng kêt nối tối đa
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
func migrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migrating Tables error :: ", err)
	}
}
