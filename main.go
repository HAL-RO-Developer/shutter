package main

import "github.com/HAL-RO-Developer/shutter/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
