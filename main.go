package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	
	router := gin.Default()
	router.Handle("GET","/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

}