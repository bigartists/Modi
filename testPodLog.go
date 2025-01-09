package main

import (
	"context"
	"fmt"
	client2 "github.com/bigartists/Modi/client"
	"io"
	v1 "k8s.io/api/core/v1"
)

func main() {
	client := client2.K8sClient
	req := client.CoreV1().Pods("aigc").GetLogs("aigc-web-7d7dd565f4-swwmh", &v1.PodLogOptions{Follow: true})

	reader, _ := req.Stream(context.Background())

	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		fmt.Println(string(buf[0:n]))
	}

}
