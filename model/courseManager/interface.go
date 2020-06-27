package coursemanager

import (
	"GymManagement/model/course"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func handleErr(err error, c *gin.Context) {
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": err.Error(),
		})
	}
}

//CreateCourse create new course
func CreateCourse(c *gin.Context) {
	var courseTmp course.Course
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	if err := c.ShouldBindJSON(&courseTmp); err != nil {
		handleErr(err, c)
		return
	}
	if err := saveCourse(&courseTmp); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

type courseID struct {
	ID string `json:"courseID"`
}

//DeleteCourse delete certain course
func DeleteCourse(c *gin.Context) {
	var courseTmp course.Course
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	var course courseID
	c.ShouldBindJSON(&course)
	courseTmp.ID = course.ID

	if err := deleteCourse(&courseTmp); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

/*
func getCourseABONDON(c *gin.Context) {
	var courseTmp course.Course
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	var course courseID
	c.ShouldBindJSON(course)
	courseTmp.ID = course.ID
	if json, err := getCourse(&courseTmp); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, json)
	}
}
*/
// PutCourse create new course
func PutCourse(c *gin.Context) {
	var courseTmp course.Course

	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	c.ShouldBindJSON(&courseTmp)
	//fmt.Println(courseTmp)
	if err := putCourse(&courseTmp); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetCourseList returns list of course
func GetCourseList(c *gin.Context) {
	if courseList, err := getCourseList(); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, courseList)
	}
}

//SelectCourse logged in user selects specific course
func SelectCourse(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	var csID courseID
	c.ShouldBindJSON(&csID)
	if err := selectCourse(csID.ID, id.(int)); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetCoachCourse return courses of certain coach
func GetCoachCourse(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("coachID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	if res, err := getCoachCourse(id.(int)); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

//GetSelectedCourse returns selected course of certain member
func GetSelectedCourse(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	if res, err := getSelectedCourse(id.(int)); err != nil {
		handleErr(err, c)
		return
	} else {
		//fmt.Println(res)
		c.JSON(http.StatusOK, res)
	}
}

//GetUnselectCourse returns course of unselected
func GetUnselectCourse(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	if res, err := getUnelectCourse(id.(int)); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}
