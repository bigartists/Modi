package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"modi/client"
	"modi/src/helpers"
	"modi/src/wscore"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		shellClient := wscore.NewWsShellClient(wsClient)

		//params := &struct {
		//	Namespace string `json:"ns"  required:"true"`
		//	Pod       string `json:"pname" required:"true"`
		//	Container string `json:"cname" required:"true"`
		//}{}
		//err = c.ShouldBindJSON(params)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//?ns=infra&pname=taichu-web-68d6b4d548-f6skj&cname=taichu-web

		//fmt.Println(params.Namespace, params.Pod, params.Container)

		err = helpers.HandleCommand("infra", "taichu-web-68d6b4d548-f6skj", "taichu-web", client.K8sClient, client.K8sClientRestConfig, []string{"sh"}).Stream(
			remotecommand.StreamOptions{
				Stdin:  shellClient,
				Stdout: shellClient,
				Stderr: shellClient,
				Tty:    true,
			})
		if err != nil {
			log.Println(err)
		}
	})
	r.Run(":7777")
}
