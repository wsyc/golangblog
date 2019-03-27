package middelware

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"newland/datebase"
)

func MyMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Auth")
		if auth =="" {
			c.String(http.StatusUnauthorized, "用户未登录")
			c.Abort()
			return
		}
		redis :=datebase.Master()
		token,err :=redis.Get(auth).Result()
		if err !=nil {
			println(err)
		}
		if token =="" {
			c.String(http.StatusUnauthorized, "Token已失效")
			c.Abort()
			return
		}
		//token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		//	func(token *jwt.Token) (interface{}, error) {
		//		return []byte("key"), nil
		//	})
		//if err == nil {
		//	if token.Valid {
		//		var p model.Person
		//		var err error
		//		Token, err := token.SignedString([]byte("key"))
		//		p.GerUserId(Token, err)
		//		c.Set("uid", p.Id)
		//	} else {
		//		c.String(http.StatusUnauthorized, "Token is not valid")
		//	}
		//} else {
		//	c.String(http.StatusUnauthorized, "Unauthorized access to this resource")
		//}
	}
}