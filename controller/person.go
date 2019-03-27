package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"log"
	"fmt"
	. "newland/models"
	"newland/util"
	"newland/datebase"
	"time"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	p := Person{UserName: username, Password: password}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	var p Person
	ra, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"date" : ra,
	})
}

func LoginPersonApi(c *gin.Context) {
	u := util.Util{}
	redis := datebase.Master()
	//todo 判断用户登录
	var p Person
	s := c.Request.Header.Get("auth")
	p.GetUserId(strconv.Atoi(s))
	uid := p.Id
	//uid:=1
	val, err := redis.Get("uid|" + string(uid)).Result()
	if val!="" {
		redis.Del(val)
		redis.Del("uid|"+string(uid))
	}
	token := u.GetToken(24)
	err = redis.Set(token, uid, 3600* time.Second).Err()
	if err != nil {
		panic("11")
	}
	redis.Set("uid|"+string(uid), token, 3600* time.Second)
	//defer redis.Close()
	//ch := make(chan int)
	//go func() {
		c.Set("uid", uid)
	//	ch <-1
	//}()
	//<- ch
	// c.Set("uid", uid)
	c.JSON(http.StatusOK, gin.H{
		"msg": "1",
		"date" : token,
	})
}