package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func (d DashboardController) Health(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (d DashboardController) Display(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Content-Type": "text/html",
	})
}
