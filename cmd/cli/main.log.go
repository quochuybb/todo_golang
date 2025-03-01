package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	sugar := zap.NewExample().Sugar()
	sugar.Info("hello world")
	encoder := GetEncoderLog()
	sync := GetWriterLog()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("hello world", zap.Int("number", 1))
	logger.Error("error", zap.Int("number", 2))
}

func GetEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func GetWriterLog() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fileSync := zapcore.AddSync(file)
	consoleSync := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(fileSync, consoleSync)
}
