package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/metrics"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		metrics.HTTPRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			string(rune(status)),
		).Inc()

		metrics.HTTPDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}
