package membermanager

import (
	"GymManagement/model/member"
	"log"
	"net/http"

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
	if err := c.ShouldBindJSON(&mem); err != nil {
		handleErr(err, c)
		return
	}
	if err := saveMember(&mem); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

//DeleteMember delete certain member
func DeleteMember(c *gin.Context) {
	var mem member.Member
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
	var mem member.Member
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
	if json, err := getMember(&mem); err != nil {
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
