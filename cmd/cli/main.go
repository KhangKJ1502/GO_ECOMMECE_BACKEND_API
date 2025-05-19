package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "tipgo", 40)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipGO"), zap.Int("age", 40))

	// logger := zap.NewExample()
	// logger.Info("Hello")
	// //Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// === CUSTOM LOGGER ===

	// Tạo encoder định dạng log
	encoder := getEncoderLog()

	// Tạo nơi ghi log: vừa ghi vào file vừa in ra console
	sync := getWriterSync()

	// Tạo core của logger với encoder, writer và mức log là Info
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	// Tạo logger từ core
	logger := zap.New(core, zap.AddCaller())

	// Ghi log với message và field kèm theo
	logger.Info("Info log", zap.Int("Line", 1))
	logger.Info("Error log", zap.Int("line", 2))
}

// Hàm tạo encoder để format log
func getEncoderLog() zapcore.Encoder {
	// Dùng cấu hình mặc định cho môi trường production
	encodeConfig := zap.NewProductionEncoderConfig()

	// Định dạng thời gian theo ISO8601: ví dụ "2025-05-19T10:26:11.170+0700"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Tên field thời gian là "Time" thay vì mặc định "ts"
	encodeConfig.TimeKey = "Time"

	// Ghi log level dưới dạng chữ in hoa: INFO, ERROR,...
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Hiển thị vị trí gọi log: "main.go:20"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Trả về encoder với định dạng JSON
	return zapcore.NewJSONEncoder(encodeConfig)
}

// Hàm tạo nơi ghi log (WriteSyncer): ghi ra file và console
func getWriterSync() zapcore.WriteSyncer {
	// Mở file log (tạo nếu chưa có, ghi tiếp nếu đã có)
	// Tạo thư mục nếu chưa tồn tại
	err := os.MkdirAll("./log", 0755)
	if err != nil {
		panic(err)
	}
	// 0666 là quyền truy cập file mặc định (rw-rw-rw-)
	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err) // Nếu mở file lỗi thì dừng chương trình
	}

	// Tạo syncer để ghi ra file
	syncFile := zapcore.AddSync(file)

	// Tạo syncer để ghi ra màn hình lỗi (stderr)
	syncConsole := zapcore.AddSync(os.Stderr)

	// Kết hợp cả 2: ghi ra file và console cùng lúc
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
