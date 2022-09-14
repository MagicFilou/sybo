package mw

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthMW: is an example middleware on how we could implement the authentication on the route level.
func AuthMW(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Apply auth middleware according to the auth for the game. Could be parsing a jwt with claims and other
		//Code here for handling auth like check the claims, see if the token is expired or others.

		c.Next()
	}
}
