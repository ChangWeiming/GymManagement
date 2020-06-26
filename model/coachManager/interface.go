package coachmanager

import (
	"GymManagement/model/coach"
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

//CreateCoach create new coach
func CreateCoach(c *gin.Context) {
	var coachTmp coach.Coach
	if err := c.ShouldBindJSON(&coachTmp); err != nil {
		handleErr(err, c)
		return
	}
	if err := saveCoach(&coachTmp); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

type coachID struct {
	ID int `json:"coachID"`
}

//DeleteCoach delete certain Coach
func DeleteCoach(c *gin.Context) {
	var coachTmp coach.Coach
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not admin logged in",
		})
		return
	}

	var coa coachID
	c.ShouldBindJSON(coa)

	coachTmp.ID = coa.ID
	if err := deleteCoach(&coachTmp); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

//GetCoach returns obj(coach)
func GetCoach(c *gin.Context) {
	var coachTmp coach.Coach
	session := sessions.Default(c)
	id := session.Get("coachID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	coachTmp.ID = id.(int)
	if json, err := getCoach(&coachTmp); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, json)
	}
}

// PutCoach create new coach
func PutCoach(c *gin.Context) {
	var coachTmp coach.Coach
	c.ShouldBindJSON(&coachTmp)
	session := sessions.Default(c)
	id := session.Get("coachID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	coachTmp.ID = id.(int)
	if err := putCoach(&coachTmp); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

//GetCoachList returns list of coach
func GetCoachList(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}

	if coachList, err := getCoachList(); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, coachList)
	}
}
