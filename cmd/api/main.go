package main



import (
	"hire/pkg/config"
	"hire/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"log"
)




func main() {
	cfg := config.Load()

	zlogs := logger.New()
	defer zlogs.Sync()
	
	router := gin.New()
	router.Use(gin.Recovery())

	
	router.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "ok"})
	})

	zlogs.Info("Server started", zap.String("port", cfg.HTTP_PORT))


	err := router.Run(":" + cfg.HTTP_PORT)
	if err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
