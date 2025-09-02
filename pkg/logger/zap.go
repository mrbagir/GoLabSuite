package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	z, _ := zap.NewProduction()
	return &ZapLogger{logger: z.Sugar()}
}

// func (z *ZapLogger) Info(msg string, fields ...Field) {
// 	z.logger.Infow(msg, toZapFields(fields...)...)
// }

// // Implement Error & Fatal

// func toZapFields(fields ...Field) []interface{} {
// 	zapFields := []interface{}{}
// 	for _, field := range fields {
// 		zapFields = append(zapFields, field.Key, field.Value)
// 	}
// 	return zapFields
// }
