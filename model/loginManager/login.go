package loginmanager

import (
	mysql "GymManagement/MySQL"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginPair struct {
	Username string `json:"phone_number"`
	Password string `json:"password"`
}

func passwordAuth(pair *loginPair, role string) (int, error) {
	db := mysql.GetDB()

	//avoid sql injection
	for _, x := range pair.Username {
		if x < '0' || x > '9' {
			return -1, nil
		}
	}

	res, err := db.Query("SELECT password,id FROM " + role + " WHERE phone_number = " + pair.Username)

	if err != nil {
		return -1, err
	}
	ans, err := mysql.GetResult(res)
	// fmt.Print(pair.Password)
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

//LoginMember member login auth & sets cookie
func LoginMember(c *gin.Context) {
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
	session.Set("memberID", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//LoginCoach coach login auth & sets cookie
func LoginCoach(c *gin.Context) {
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
	session.Set("coachID", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

//LoginAdmin admin login auth & sets cookie
func LoginAdmin(c *gin.Context) {
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
	session.Set("adminID", ans)
	session.Set("username", pair.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
