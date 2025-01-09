package main

import (
	"fmt"
	"github.com/bigartists/Modi/cmd/app"
	"github.com/bigartists/Modi/config"
	_ "github.com/bigartists/Modi/config"
)

func main() {
	cmd := app.NewApiServerCommand()
	fmt.Println("config.SysYamlconfig.Server.Name = ", config.SysYamlconfig.Server.Name)
	cmd.Execute()
}

// go run cmd/apiserver.go apiserver --port=8888
