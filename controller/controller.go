package controller

import (
	mysql "GymManagement/MySQL"
	adminManager "GymManagement/model/adminManager"
	coachManager "GymManagement/model/coachManager"
	courseManager "GymManagement/model/courseManager"
	loginManager "GymManagement/model/loginManager"
	memberManager "GymManagement/model/memberManager"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//StartServer starts http server
func StartServer() {
	router := gin.Default()
	store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("userLogin", store))
	router.Use(func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("data: %v\n", string(data))
		//很关键
		//把读过的字节流重新放到body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	})
	mysql.RunMySQL()

	router.POST("/api/member", memberManager.CreateMember)
	router.POST("/api/coach", coachManager.CreateCoach)
	router.POST("/api/admin", adminManager.CreateAdmin)
	router.POST("/api/member/session", loginManager.LoginMember)
	router.POST("/api/admin/session", loginManager.LoginAdmin)
	router.POST("/api/coach/session", loginManager.LoginCoach)
	router.POST("/api/course", courseManager.CreateCourse)
	router.POST("/api/selection", courseManager.SelectCourse)
	router.POST("/api/start_time", memberManager.PostStartTime)
	router.POST("/api/leave_time", memberManager.PostLeaveTime)

	router.DELETE("/api/member", memberManager.DeleteMember)
	router.DELETE("/api/coach", coachManager.DeleteCoach)
	router.DELETE("/api/admin", adminManager.DeleteAdmin)
	router.DELETE("/api/course", courseManager.DeleteCourse)

	router.PUT("/api/member", memberManager.PutMember)
	router.PUT("/api/coach", coachManager.PutCoach)
	router.PUT("/api/admin", adminManager.PutAdmin)
	router.PUT("/api/course", courseManager.PutCourse)
	router.PUT("/api/term", memberManager.PutTerm)

	router.GET("/api/member", memberManager.GetMember)
	router.GET("/api/coach", coachManager.GetCoach)
	router.GET("/api/admin", adminManager.GetAdmin)
	router.GET("/api/courselist", courseManager.GetCourseList)
	router.GET("/api/memberlist", memberManager.GetMemberList)
	router.GET("/api/coachlist", coachManager.GetCoachList)
	router.GET("/api/term/:phone_number", memberManager.GetTerm)
	router.GET("/api/coachcourse", courseManager.GetCoachCourse)
	router.GET("/api/selectcourse", courseManager.GetSelectedCourse)
	router.GET("/api/unselectcourse", courseManager.GetUnselectCourse)
	router.Run(":8001")
}
