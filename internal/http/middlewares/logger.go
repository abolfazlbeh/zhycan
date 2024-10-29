package middlewares

import (
	"github.com/abolfazlbeh/zhycan/internal/logger"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func ZapLogger() gin.HandlerFunc {
	logge, err := logger.GetManager().GetLogger()
	if err != nil {
		return nil
	}
	logger1 := logge.(*logger.ZapWrapper).Instance()

	return ginzap.Ginzap(logger1, time.RFC3339, true)
}

func ZapRecoveryLogger() gin.HandlerFunc {
	logge, err := logger.GetManager().GetLogger()
	if err != nil {
		return nil
	}
	logger1 := logge.(*logger.ZapWrapper).Instance()

	return ginzap.RecoveryWithZap(logger1, true)
}
