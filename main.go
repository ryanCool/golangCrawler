package main

import "github.com/gin-gonic/gin"

func main() {
	r:=gin.Default()

	initRouter(r)

	r.Run()
}