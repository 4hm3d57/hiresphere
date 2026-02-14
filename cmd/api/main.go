package main



import (
	"hire/pkg/config"
	"hire/pkg/logger"
	"hire/pkg/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"log"
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

	zlogs.Info("Server started", zap.String("port", cfg.HTTPport))


	err = router.Run(":" + cfg.HTTPport)
	if err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
