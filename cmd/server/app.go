package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type App struct {
	engine *gin.Engine
}

func ProvideApp(e *gin.Engine) *App {
	return &App{
		engine: e,
	}
}

func (a *App) Run(port int) error {
	err := a.engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

func (a *App) NewApiServerCommand() (cmd *cobra.Command) {
	// 集成 cobra命令
	cmd = &cobra.Command{
		Use: "appServer",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			return a.Run(port)
		},
	}
	// 添加 flag, name=port, 默认值是 9090
	cmd.Flags().Int("port", 9090, "appserver port")
	return
}
