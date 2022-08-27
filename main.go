package main

import (
	"math/rand"
	"time"

	App "github.com/photowey/tsc/cmd/cmder/app"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	App.Run()
}
