package user

import (
	"fmt"
	"net/http"

	cfg "sybo/configs"
	userhandler "sybo/handler/user"
	l "sybo/logger"
	usermodel "sybo/models/user"
	"sybo/utils"

	"sybo/mw"

	"github.com/gin-gonic/gin"
)

func UserGroup(r *gin.Engine) {

	userRoutes := r.Group("/user",
		mw.AuthMW(l.GetLogger()),
	)
	{

		userRoutes.POST("",
			func(c *gin.Context) {

				var user usermodel.User

				if err := c.ShouldBindJSON(&user); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				err := userhandler.New(&user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"id":   user.ID,
					"name": user.Name,
				})
			})

		userRoutes.PUT("/:userid/state",
			func(c *gin.Context) {

				var user usermodel.User

				if err := c.ShouldBindJSON(&user); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(400, gin.H{"error": fmt.Errorf("User ID provided is not a valid uuid")})
					return
				}

				err := userhandler.SaveState(&user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
					return
				}

			})

		userRoutes.GET("/:userid/state",
			func(c *gin.Context) {

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(400, gin.H{"error": fmt.Errorf("User ID provided is not a valid uuid")})
					return
				}

				err := userhandler.LoadState(&user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"gamesPlayed": user.GamesPlayed,
					"score":       user.Score,
				})
			})

		userRoutes.PUT("/:userid/friends",
			func(c *gin.Context) {

				var friends usermodel.FriendsList

				if err := c.ShouldBindJSON(&friends); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(400, gin.H{"error": fmt.Errorf("User ID provided is not a valid uuid")})
					return
				}

				err := userhandler.UpdateFriends(friends, &user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
					return
				}

			})

		userRoutes.GET("/:userid/friends",
			func(c *gin.Context) {

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(400, gin.H{"error": fmt.Errorf("User ID provided is not a valid uuid")})
					return
				}

				friends, err := userhandler.GetFriends(&user)
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
					return
				}

				if len(friends) == 0 {
					c.Status(cfg.CODE_EMPTY)
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"friends": friends,
				})

			})
	}
}
