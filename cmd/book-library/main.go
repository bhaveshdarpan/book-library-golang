package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/gorilla/mux"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}