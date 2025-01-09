package main

import (
	"fmt"
	"github.com/bigartists/Modi/client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

func main() {
	// 远程执行pod命令的方法
	// kubectl exec -it   nginx-7875-5qqbr  -c nginx  -- sh -c ls
	// 实现上述需求

	option := &v1.PodExecOptions{
		Container: "taichu-web",
		Command:   []string{"sh", "-c", "ls"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
	}
	req := client.K8sClient.CoreV1().
		RESTClient().Post().
		Resource("pods").
		Namespace("infra").
		Name("taichu-web-6d65b576d4-7v272").
		SubResource("exec").
		VersionedParams(
			option,
			scheme.ParameterCodec)

	fmt.Print(req.URL())

	// http 1.1 SPDY 2.0
	exec, err := remotecommand.NewSPDYExecutor(client.K8sClientRestConfig, "POST", req.URL())
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})

	if err != nil {
		log.Fatal(err)
	}

}
