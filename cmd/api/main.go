package main



import (
	"hire/pkg/config"
	"hire/pkg/logger"
	"hire/pkg/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
)




func main() {
	// configs
	cfg := config.Load()


	// logger
	zlogs := logger.New()
	defer zlogs.Sync()
	
	// database
	db, err := database.New(cfg)
	if err != nil {
		zlogs.Fatal("Failed to connect db: ", zap.Error(err))
	}
	defer db.Close()
	zlogs.Info("Database connected");

	// gin http framework
	router := gin.New()
	router.Use(gin.Recovery())
	
	router.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	// creating http server config
	server := &http.Server {
		Addr: ":" + cfg.HTTPport,
		Handler: router,
	}
	

	// start server in goroutine
	go func() {
		zlogs.Info("Server started", zap.String("port", cfg.HTTPport))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			zlogs.Fatal("Server failed", zap.Error(err))
		}
	}()

	
	// create channel signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	
	zlogs.Info("Shutdown signal received")

	// shutdown server gracefuly
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	if err = server.Shutdown(ctx); err != nil {
		zlogs.Fatal("Shutdown failed", zap.Error(err))
	}

	zlogs.Info("Closing database connection")

	zlogs.Info("Server exited cleanly")
}
