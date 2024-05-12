package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/handle-click", func(c *gin.Context) {

		session := sessions.Default(c)

		buttonClicked := c.PostForm("button")
		switch buttonClicked {
		case "get-cookies":
			session.Set("example", "examplevalue")
			session.Save()
			value := session.Get("example")
			resp := map[string]string{
				"message": "cookie created",
				"exampleValue": value.(string), // Convert the value to string
			}
			c.HTML(200, "index.html", resp)
		case "remove-cookies":
			c.SetCookie("mysession", "", -1, "/", "localhost", false, true)
			resp := map[string]string{
				"message": "cookie removed",
			}
			c.HTML(200, "index.html", resp)
		case "get-cookie-value":
			v, err := c.Cookie("mysession")
			if err != nil {
				v = "Cookie not found"
			}
			resp := map[string]interface{}{
				"cookieValue": v,
			}
			c.HTML(200,"index.html",resp)
		default:
			c.String(http.StatusBadRequest, "Invalid button clicked")
		}
	})

	r.Run()
}
