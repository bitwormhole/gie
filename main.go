package main

import (
	"embed"

	"github.com/bitwormhole/gie/modules/gie"
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/main/resources
var theMainRes embed.FS

func main() {

	res := collection.LoadEmbedResources(&theMainRes, "src/main/resources")

	i := starter.InitApp()
	i.Use(gie.Module(res))
	i.Run()
}
