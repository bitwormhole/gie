package main

import "github.com/bitwormhole/gie/boot"

func main() {
	b := boot.Default()
	err := boot.Run(b)
	if err != nil {
		panic(err)
	}
}
