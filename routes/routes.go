package routes

import (
	u "sybo/routes/user"

	"github.com/gin-gonic/gin"
)

// CollectGroups: group of routegroups if we want to extend for more routing
func CollectGroups(r *gin.Engine) {

	u.UserGroup(r)
}
