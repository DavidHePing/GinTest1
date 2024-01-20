package customer_logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (fileLogger *zap.Logger) {
	createFileWithDir("./test.log")

	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.OutputPaths = []string{"./test.log"}
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileLogger, _ = config.Build()
	return
}

func createFileWithDir(filePath string) {
	// Create the directory along with any necessary parents
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	return
}
