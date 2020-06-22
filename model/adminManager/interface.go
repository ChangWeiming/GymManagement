package adminmanager

import (
	"GymManagement/model/admin"
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

//CreateAdmin create new administrator
func CreateAdmin(c *gin.Context) {
	var administrator admin.Admin
	if err := c.ShouldBindJSON(&administrator); err != nil {
		handleErr(err, c)
		return
	}
	if err := saveAdmin(&administrator); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

//DeleteAdmin delete certain administrator
func DeleteAdmin(c *gin.Context) {
	var administrator admin.Admin
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	administrator.ID = id.(int)
	if err := deleteAdmin(&administrator); err != nil {
		handleErr(err, c)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

//GetAdmin returns obj(admin)
func GetAdmin(c *gin.Context) {
	var administrator admin.Admin
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	administrator.ID = id.(int)
	if json, err := getAdmin(&administrator); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, json)
	}
}

// PutAdmin create new admin
func PutAdmin(c *gin.Context) {
	var administrator admin.Admin
	c.ShouldBindJSON(&administrator)
	session := sessions.Default(c)
	id := session.Get("adminID")
	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "failed",
			"errlog": "not logged in",
		})
		return
	}
	administrator.ID = id.(int)
	if err := putAdmin(&administrator); err != nil {
		handleErr(err, c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
