package membermanager

import (
	"GymManagement/model/member"
	"log"

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
	if err := c.ShouldBindJSON(&mem); err != nil {
		handleErr(err, c)
		return
	}
	if err := deleteMember(&mem); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}
