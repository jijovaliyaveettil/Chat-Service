package handlers

import "github.com/gin-gonic/gin"

func LoginFormHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
