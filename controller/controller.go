package controller

import (
	mysql "GymManagement/MySQL"
	adminmanager "GymManagement/model/adminManager"
	coachmanager "GymManagement/model/coachManager"
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
	mysql.RunMySQL()

	router.POST("/api/member", membermanager.CreateMember)
	router.POST("/api/coach", coachmanager.CreateCoach)
	router.POST("/api/admin", adminmanager.CreateAdmin)
	router.POST("/api/member/session", loginMember)
	router.POST("/api/admin/session", loginAdmin)
	router.POST("/api/coach/session", loginCoach)

	router.DELETE("/api/member", membermanager.DeleteMember)
	router.DELETE("/api/coach", coachmanager.DeleteCoach)
	router.DELETE("/api/admin", adminmanager.DeleteAdmin)

	router.PUT("/api/member", membermanager.PutMember)
	router.PUT("/api/coach", coachmanager.PutCoach)
	router.PUT("/api/admin", adminmanager.PutAdmin)

	router.GET("/api/member", membermanager.GetMember)
	router.GET("/api/coach", coachmanager.GetCoach)
	router.GET("/api/admin", adminmanager.GetAdmin)

	router.Run(":8001")
}
