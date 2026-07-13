package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mdtajulislammt/zvonsystem_backend/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/middleware"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/metrics"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetupRouter(lc fx.Lifecycle, ctg *config.Config, r *gin.Engine, log *zap.Logger) {
	// logger
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger(log))

	// prometheus metrics
	metrics.Register()
	r.Use(middleware.Prometheus())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/health", func(c *gin.Context) {
		log.Info("health_check")
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to ZvonSystem API",
			"success": true,
		})
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				listenAddr := fmt.Sprintf("0.0.0.0:%s", ctg.App.Port)
				if err := r.Run(listenAddr); err != nil && err != http.ErrServerClosed {
					fmt.Println("Failed to start server:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
