package service

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/markup"
)

// G.I.E 的图形界面通过调用系统的浏览器实现

type GuiService struct {
	markup.Component `id:"gui-service" initMethod:"Init"`

	ClientFactory cli.ClientFactory `inject:"#cli-client-factory"`

	Port int `inject:"${server.port}"`
}

func (inst *GuiService) enabled() bool {
	return false
}

func (inst *GuiService) Init() error {

	if !inst.enabled() {
		return nil
	}

	paths := []string{
		"C:\\Program Files (x86)\\Google\\Chrome\\Application",
		"C:\\Program Files (x86)\\Microsoft\\Edge\\Application",
		"C:\\Program Files\\Mozilla Firefox",
	}
	const key = "PATH"
	path := os.Getenv(key)
	for _, p := range paths {
		path = path + ";" + p
	}
	os.Setenv(key, path)

	port := strconv.Itoa(inst.Port)
	cmd := exec.Command("msedge", "http://localhost:"+port)
	cmd.Start()
	// cmd.Wait()

	return nil
}
