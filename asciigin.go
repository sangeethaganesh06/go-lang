package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "english ",
			"tag":  "<br>",
			//kjgytfyunmlojokjounm
		}
		c.AsciiJSON(http.StatusOK, data)
	})


	r.Run( )
}
