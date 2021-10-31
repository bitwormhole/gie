package main

import (
	"github.com/bitwormhole/gie"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.Use(gie.Module())
	i.Run()
}
