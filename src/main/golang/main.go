package main

import "github.com/bitwormhole/gie/app/boot"

func main() {
	// i := starter.InitApp()
	// i.Use(gie.Module())
	// i.Run()

	b := boot.Default()
	err := boot.Run(b)
	if err != nil {
		panic(err)
	}
}
