package initialize

import (
	"context"
	"fmt"
	"go_ecommerce_backend/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	fmt.Printf("Redis config: host=%s, port=%d\n", r.Host, r.Port) // ✅ Kiểm tra rõ ràng

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error", zap.Error(err))
		return
	}

	fmt.Println("Init Redis is running")
	global.Rdb = rdb
	RedisExample()
}

func RedisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:: ", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting:: ", zap.Error(err))
		return
	}
	global.Logger.Info("Value score:: ", zap.String("score", value))
}
