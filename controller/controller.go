package controller

import (
	mysql "GymManagement/MySQL"
	adminmanager "GymManagement/model/adminManager"
	coachmanager "GymManagement/model/coachManager"
	coursemanager "GymManagement/model/courseManager"
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
	router.POST("/api/course", coursemanager.CreateCourse)
	router.POST("/api/selection", coursemanager.SelectCourse)
	router.POST("/api/start_time", membermanager.PostStartTime)
	router.POST("/api/leave_time", membermanager.PostLeaveTime)

	router.DELETE("/api/member", membermanager.DeleteMember)
	router.DELETE("/api/coach", coachmanager.DeleteCoach)
	router.DELETE("/api/admin", adminmanager.DeleteAdmin)
	router.DELETE("/api/course", coursemanager.DeleteCourse)

	router.PUT("/api/member", membermanager.PutMember)
	router.PUT("/api/coach", coachmanager.PutCoach)
	router.PUT("/api/admin", adminmanager.PutAdmin)
	router.PUT("/api/course", coursemanager.PutCourse)

	router.GET("/api/member", membermanager.GetMember)
	router.GET("/api/coach", coachmanager.GetCoach)
	router.GET("/api/admin", adminmanager.GetAdmin)
	router.GET("/api/courselist", coursemanager.GetCourseList)
	router.GET("/api/memberlist", membermanager.GetMemberList)
	router.GET("/api/coachlist", coachmanager.GetCoachList)
	router.GET("/api/term", membermanager.GetTerm)
	router.GET("/api/coachcourse", coursemanager.GetCoachCourse)
	router.GET("/api/selectcourse", coursemanager.GetSelectedCourse)
	router.GET("/unselectcourse", coursemanager.GetUnselectCourse)
	router.Run(":8001")
}
