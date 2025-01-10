package main

import (
	"fmt"
	"github.com/bigartists/Modi/cmd/inject"
	"github.com/bigartists/Modi/config"
	_ "github.com/bigartists/Modi/config"
	"os"
)

func main() {
	app, err := inject.InitializeApp()
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := app.NewApiServerCommand()

	fmt.Println("config.SysYamlconfig.Server.Name = ", config.SysYamlconfig.Server.Name)
	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
		return
	}
}

// go run cmd/apiserver.go apiserver --port=8888
