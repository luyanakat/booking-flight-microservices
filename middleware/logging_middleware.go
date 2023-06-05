package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Since(start)

		logger.Info("Completed server call",
			zap.String("address", c.Request.RequestURI),
			zap.String("duration", end.String()))
	}
}
