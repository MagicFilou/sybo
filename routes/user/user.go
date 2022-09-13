package user

import (
	"net/http"
	"strings"

	cfg "sybo/configs"
	userhandler "sybo/handler/user"
	l "sybo/logger"
	usermodel "sybo/models/user"

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
					c.AbortWithStatusJSON(checkError(err))
					return
				}

				c.JSON(cfg.CODE_SUCCESS, gin.H{
					"id":   user.ID,
					"name": user.Name,
				})
			})

		// userRoutes.METHOD("/:id",
		// func(c *gin.Context) {

		// 	ID := c.Param("id")
		// 	intID, err := strconv.Atoi(ID)

	}
}

//TODO check if necessary
func checkError(err error) (int, gin.H) {

	switch {

	case strings.Contains(err.Error(), "out of bounds"):
		return cfg.CODE_BADREQUEST, gin.H{"status": cfg.STATUS_BADREQUEST, "error": err.Error()}

	case strings.Contains(err.Error(), "not found"):
		return cfg.CODE_EMPTY, gin.H{"status": cfg.STATUS_EMPTY, "error": err.Error()}

	default:
		return cfg.CODE_BADREQUEST, gin.H{"status": cfg.STATUS_BADREQUEST, "error": err.Error()}
	}
}
