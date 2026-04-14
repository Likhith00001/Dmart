package main

import (
	"context"
	"fmt"
	"net/http"
	"order-service/internal/config"
	"order-service/internal/handler"
	"order-service/internal/model"
	"order-service/internal/repository"
	"order-service/internal/service"
	"os"
	"os/signal"
	"shared/pkg/logger"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	logger.Init(cfg.Server.Env)
	log := logger.Get()

	log.Info("starting Order Service...", zap.String("env", cfg.Server.Env),
		zap.String("port", cfg.Server.Port))

	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	if err := db.AutoMigrate(&model.Order{}); err != nil {
		log.Warn("Automigrate failed", zap.Error(err))
	} else {
		log.Info("Database migrated successfully")
	}

	orderRepo := repository.NewOrderRepository(db)
	orderSvc := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSvc)

	// Gin Router
	r := gin.Default()

	// Routes
	api := r.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.POST("", orderHandler.CreateOrder)
			products.GET("", orderHandler.GetOrder)
			// products.GET("/:id", productHandler.GetByID)
			// products.PUT("/:id", productHandler.Update)
			// products.DELETE("/:id", productHandler.Delete)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "order-service",
			"version": "0.1.0",
		})
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Info("HTTP server started", zap.String("address", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Product Service...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Product Service exited gracefully")
}
