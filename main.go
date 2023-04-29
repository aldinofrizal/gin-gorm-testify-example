package main

import (
	"golang-web-testing/route"
)

func main() {
	r := route.SetupRoutes()
	r.Run()
}
