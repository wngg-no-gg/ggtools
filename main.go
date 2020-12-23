package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, test")
    })

    r.Run(":48080") // listen and serve on 0.0.0.0:8080
}