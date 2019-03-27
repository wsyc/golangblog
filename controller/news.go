package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"log"
	"fmt"
	. "newland/models"
)

//func AddNewsApi(c *gin.Context) {
//	username := c.Request.FormValue("username")
//	password := c.Request.FormValue("password")
//
//	p := Person{UserName: username, Password: password}
//
//	ra, err := p.AddPerson()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	msg := fmt.Sprintf("insert successful %d", ra)
//	c.JSON(http.StatusOK, gin.H{
//		"msg": msg,
//	})
//}

func GetNewsApi(c *gin.Context) {
	var n News
	ra, err := n.GetNews()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"date" : ra,
	})
}
