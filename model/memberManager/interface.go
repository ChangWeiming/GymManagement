package membermanager

import (
	"GymManagement/model/member"
	"log"

	"github.com/gin-gonic/gin"
)

//CreateMember create new member
func CreateMember(c *gin.Context) {
	var mem member.Member
	if err := c.ShouldBindJSON(&mem); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": "failed",
		})
		return
	}
	if err := saveMember(&mem); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"status": "failed",
		})
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}
