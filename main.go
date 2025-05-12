package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	err := r.Run(":3100")

	if err != nil {
		return
	}
}
