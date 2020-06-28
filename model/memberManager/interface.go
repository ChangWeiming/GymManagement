package membermanager

import (
	"GymManagement/model/member"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func handleErr(err error, c *gin.Context) {
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{
			"status": "failed",
			"errlog": err.Error(),
		})
	}
}

//CreateMember create new member
func CreateMember(c *gin.Context) {
	var mem member.Member
	/*
		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(body))
	*/
	if err := c.ShouldBindJSON(&mem); err != nil {
		fmt.Print("err1")
		handleErr(err, c)
		return
	}
	if err := saveMember(&mem); err != nil {
		fmt.Print("err2")
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

type memberID struct {
	ID string `json:"memberID"`
}

//DeleteMember delete certain member
func DeleteMember(c *gin.Context) {
	var mem member.Member
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	var member memberID
	c.ShouldBindJSON(&member)
	mem.ID, _ = strconv.Atoi(member.ID)
	if err := deleteMember(&mem); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

//GetMember returns obj(member)
func GetMember(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("coachID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	if json, err := getMember(id.(int)); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, json)
	}
}

// PutMember create new member
func PutMember(c *gin.Context) {
	var mem member.Member
	c.ShouldBindJSON(&mem)
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	mem.ID = id.(int)
	if err := putMember(&mem); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//PutTerm set term
func PutTerm(c *gin.Context) {
	var mem member.Member
	c.ShouldBindJSON(&mem)
	if err := putTerm(&mem); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetMemberList returns list of coach
func GetMemberList(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	if memberList, err := getMemberList(); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, memberList)
	}
}

//PostStartTime posts start time of certain member
func PostStartTime(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	startTimeTest := session.Get("start_time")

	if startTimeTest != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "you have checked in",
		})
		return
	}

	type data struct {
		StartTime string `json:"start_time"`
		ID        string `json:"courseID"`
	}

	var courseID data
	c.ShouldBindJSON(&courseID)

	session.Set("start_time", courseID.StartTime)
	session.Save()

	if err := postStartTime(id.(int), courseID.ID, courseID.StartTime); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//PostLeaveTime posts leave time of certain member
func PostLeaveTime(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("memberID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	type data struct {
		LeaveTime string `json:"leave_time"`
		ID        string `json:"courseID"`
	}

	var courseID data
	c.ShouldBindJSON(&courseID)

	startTime := session.Get("start_time")
	if startTime == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not checked in",
		})
		return
	}

	if err := postLeaveTime(id.(int), courseID.ID, startTime.(string), courseID.LeaveTime); err != nil {
		handleErr(err, c)
		return
	} else {
		session.Delete("start_time")
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetTerm return term of certain phone number
func GetTerm(c *gin.Context) {
	var mem member.Member
	mem.PhoneNumber = c.Param("phone_number")
	//fmt.Println(mem.PhoneNumber)
	if res, err := getTerm(&mem); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}
