package controller

import (
	membermanager "GymManagement/model/memberManager"

	"github.com/gin-gonic/gin"
)

//StartServer starts http server
func StartServer() {
	router := gin.Default()
	router.POST("/api/member", membermanager.CreateMember)
	router.DELETE("/api/member", membermanager.DeleteMember)
}

func loginMember(c *gin.Context) {

}
