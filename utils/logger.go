package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLoggerFile() {
	// 配置日志输出到文件
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log", // 日志文件路径
		MaxSize:    100,            // 日志文件最大大小（MB）
		MaxBackups: 3,              // 保留的旧日志文件最大数量
		MaxAge:     28,             // 保留旧日志文件的最大天数
		Compress:   true,           // 是否压缩旧日志文件
	})

	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON 格式
		writer,                                // 输出到文件
		zap.InfoLevel,                         // 日志级别
	)

	// 创建 Logger
	Logger = zap.New(core, zap.AddCaller())
	defer Logger.Sync()

	// 替换全局的日志器和 SugaredLogger
	zap.ReplaceGlobals(Logger)
}

func InitLogger() {
	var err error
	// 使用生产环境的日志配置（JSON 格式，包含调用堆栈）
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// 替换全局的日志器和 SugaredLogger
	zap.ReplaceGlobals(Logger)
}
