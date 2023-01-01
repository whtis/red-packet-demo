package main

import (
	"ginDemo/dal"
	"github.com/gin-gonic/gin"
)

func main() {
	dal.InitDB()
	r := gin.Default()
	register(r)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
