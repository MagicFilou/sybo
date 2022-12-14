package user

import (
	"fmt"
	"net/http"
	"strings"

	l "sybo/logger"
	"sybo/models"
	"sybo/mw"
	"sybo/utils"

	userhandler "sybo/handler/user"
	usermodel "sybo/models/user"

	"github.com/gin-gonic/gin"
)

//UserGroup: func for the routes of the user group model
func UserGroup(r *gin.Engine) {

	userRoutes := r.Group("/user",
		mw.AuthMW(l.GetLogger()),
	)
	{

		//Endpoint to get all users
		userRoutes.GET("",
			func(c *gin.Context) {

				paramMap := c.Request.URL.Query()

				var whereData []models.WhereData
				for key, value := range paramMap {
					whereData = append(whereData, models.WhereData{
						Field: key,
						//TODO room for improvement but for now just using the first instance of each param
						Value: value[0],
					})
				}

				users, err := userhandler.Get(whereData)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				//if empty return code empty
				if len(users) == 0 {
					c.Status(http.StatusNoContent)
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"users": users,
				})
			})

		//Endpoint to add  a new user
		userRoutes.POST("",
			func(c *gin.Context) {

				var user usermodel.User

				if err := c.ShouldBindJSON(&user); err != nil {
					c.JSON(checkError(err))
					return
				}

				if len(user.Name) == 0 {
					c.JSON(checkError(fmt.Errorf("User data given should have a name")))
					return
				}

				err := userhandler.New(&user)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":   user.ID,
					"name": user.Name,
				})
			})

		//Endpoint to get the state of a user
		userRoutes.GET("/:userid/state",
			func(c *gin.Context) {

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("User ID provided is not a valid uuid")))
					return
				}

				err := userhandler.LoadState(&user)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"gamesPlayed": user.GamesPlayed,
					"score":       user.Score,
				})
			})

		//Endpoint to update the state of a user
		userRoutes.PUT("/:userid/state",
			func(c *gin.Context) {

				var user usermodel.User

				if err := c.ShouldBindJSON(&user); err != nil {
					c.JSON(checkError(err))
					return
				}

				if user.GamesPlayed == 0 {
					c.JSON(checkError(fmt.Errorf("User data given should have at least one game played")))
					return
				}

				user.ID = c.Param("userid")

				//Check if userid is a valid
				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("User ID provided is not a valid uuid")))
					return
				}

				err := userhandler.SaveState(&user)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}
			})

		//Endpoint to get the friends list for a user
		userRoutes.GET("/:userid/friends",
			func(c *gin.Context) {

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("User ID provided is not a valid uuid")))
					return
				}

				friends, err := userhandler.GetFriends(&user)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				if len(friends) == 0 {
					c.Status(http.StatusNoContent)
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"friends": friends,
				})
			})

		//Endpoint to update the friends list for a user
		userRoutes.PUT("/:userid/friends",
			func(c *gin.Context) {

				//Get the given list. Because I keep it simple, it is jsut a []string converted further into a comma separated list
				var friends usermodel.FriendsList

				if err := c.ShouldBindJSON(&friends); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if len(friends.Friends) == 0 {
					c.JSON(checkError(fmt.Errorf("User data given should have at least one friend")))
					return
				}

				var user usermodel.User

				user.ID = c.Param("userid")

				if ok := utils.IsValidUUID(user.ID); !ok {
					c.AbortWithStatusJSON(checkError(fmt.Errorf("User ID provided is not a valid uuid")))
					return
				}

				err := userhandler.UpdateFriends(friends, &user)
				if err != nil {
					c.AbortWithStatusJSON(checkError(err))
					return
				}
			})
	}
}

// checkError: it's a convenience function to return a bad request or internal server error according to the error received.
// It can be extend to have any sort of errors
func checkError(err error) (int, gin.H) {

	switch {

	case strings.Contains(err.Error(), "not a valid uuid"), strings.Contains(err.Error(), "No users with"), strings.Contains(err.Error(), "User data given"):
		return http.StatusBadRequest, gin.H{"error": err.Error()}

	default:
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}
}
