package main

import (
	"github.com/taise-hub/edush/router"
)

func main() {
	r := router.Init()
	r.Run()
}