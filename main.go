package main

import (
	cfg "sybo/configs"
	l "sybo/logger"
	"sybo/routes"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	// Init the logger
	logger := l.GetLogger()
	defer logger.Sync()

	r := ginInit(logger)

	//setup routes for the gin
	routes.CollectGroups(r)

	logger.Info("server started",
		zap.String("PORT", cfg.GetConfig().Port),
	)

	// Run the server
	err := r.Run(cfg.GetConfig().Port)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

// ginInit: initialize the gin server
func ginInit(logger *zap.Logger) *gin.Engine {

	//Setup if dev or not
	if cfg.GetConfig().ENV != "local" {

		gin.SetMode(gin.ReleaseMode)
	}

	eng := gin.New()

	eng.Use(ginzap.RecoveryWithZap(logger, true))

	return eng
}
