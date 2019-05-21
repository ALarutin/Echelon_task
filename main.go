package main

import (
	"github.com/ALarutin/Echelon_task/router"
)

func main() {
	PORT := ":8080"
	r := router.GetRouter()
	r.Run(PORT)
}
