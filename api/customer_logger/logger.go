package customer_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (fileLogger *zap.Logger) {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.OutputPaths = []string{"./test.log"}
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileLogger, _ = config.Build()
	return
}
