package controller

import (
	membermanager "GymManagement/model/memberManager"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//StartServer starts http server
func StartServer() {
	router := gin.Default()
	store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("userLogin", store))

	router.POST("/api/member", membermanager.CreateMember)
	router.POST("/api/member/session", loginMember)
	router.POST("/api/admin/session", loginAdmin)
	router.POST("/api/coach/session", loginCoach)

	router.DELETE("/api/member", membermanager.DeleteMember)

	router.PUT("/api/member", membermanager.PutMember)

	router.GET("/api/member", membermanager.GetMember)
	router.Run(":8001")
}
