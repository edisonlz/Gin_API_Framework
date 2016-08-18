package sessions

import (
	"github.com/gin-gonic/gin"
	"Gin_API_Framework/middleware/contrib/secure_cookie"
	"net/http"
)

const (
	COOKIE_MAX_AGE = 122344
	COOKIE_DOMAIN = "*"
	COOKIE_PATH = "/"
)


// set secure cookie
func AuthLogin(c *gin.Context, uid string)  {
	secure_cookie.SetSecureCookie(
		c,
		"user_token",
		uid,
		COOKIE_MAX_AGE,
		COOKIE_PATH,
		COOKIE_DOMAIN,
		true,true)
}

// delete cookie user_token
func AuthLogout(c *gin.Context)  {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "user_token",
		Value:    "",
		MaxAge:   -1,
		Path:     COOKIE_PATH,
		Domain:   COOKIE_DOMAIN,
		Secure:   true,
		HttpOnly: true,
	})
}

// Login Require Decorator
func LoginRequired(handle gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		userToken, cookie_err := secure_cookie.GetSecureCookie(c,"user_token",1)

		var is_login  bool = true

		if cookie_err != nil{
			is_login = false
		}

		//Tudo 添加查数据库逻辑

		if is_login == false{
			c.JSON(http.StatusUnauthorized,
				gin.H{
					"status":  "failed",
					"desc": "login requierd",
				})
		}else {
			handle(c)
			c.Set("currentUserId",userToken)
			c.Set("currentUser", userToken)
		}
	}
}