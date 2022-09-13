package mw

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMW(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("Apply auth middleware according to the auth for the game. Could be parsing a jwt with claims and other")
		//Code here for handling auth

		c.Next()
	}
}
