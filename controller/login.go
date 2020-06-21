package controller

import (
	mysql "GymManagement/MySQL"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginPair struct {
	Username string `json:"phone_number"`
	Password string `json:"passowrd"`
}

func passwordAuth(pair *loginPair, role string) (int, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT password,id FROM "+role+" WHRER phone_number = ?", pair.Username)
	if err != nil {
		return -1, err
	}
	ans, err := mysql.GetResult(res)

	if len(ans) == 0 {
		return -1, nil
	}

	if ans[0]["password"] == pair.Password {
		id, _ := strconv.Atoi(ans[0]["id"])
		return id, nil
	} else {
		return -1, nil
	}
}

func handleErr(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": err.Error(),
		})
		return true
	}
	return false
}

func loginMember(c *gin.Context) {
	var pair loginPair
	c.ShouldBindJSON(&pair)
	ans, err := passwordAuth(&pair, "member")
	if handleErr(err, c) {
		return
	}
	if ans < 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("id", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func loginCoach(c *gin.Context) {
	var pair loginPair
	c.ShouldBindJSON(&pair)
	ans, err := passwordAuth(&pair, "coach")
	if handleErr(err, c) {
		return
	}
	if ans < 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("id", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func loginAdmin(c *gin.Context) {
	var pair loginPair
	c.ShouldBindJSON(&pair)
	ans, err := passwordAuth(&pair, "admin")
	if handleErr(err, c) {
		return
	}
	if ans < 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("id", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
