package logger

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	logger.Info("trhrtrthrrhr")
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "kjnknkjnknknk",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "kjnknkjnknknk")
}
