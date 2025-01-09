package main

import (
	"context"
	"github.com/bigartists/Modi/client"
	v1 "k8s.io/api/core/v1"
)

func main() {
	req := client.K8sClient.CoreV1().Pods("infra").GetLogs("taichu-web-68d6b4d548-f6skj", &v1.PodLogOptions{
		//Follow: true,
	})
	podLogs := req.Do(context.Background())
	b, _ := podLogs.Raw()
	println(string(b))

	//reader, _ := req.Stream(context.Background())
	//for {
	//	buf := make([]byte, 1024)
	//	n, err := reader.Read(buf)
	//	if err != nil && err != io.EOF {
	//		break
	//	}
	//	fmt.Println(string(buf[0:n]))
	//}

	//req := client.K8sClient.CoreV1().Pods("infra").GetLogs("taichu-web-6555476964-6qx4j", &v1.PodLogOptions{})
	//reader, _ := req.Stream(context.Background())
	//defer reader.Close()
	//var logs string
	//for {
	//	buf := make([]byte, 1024)
	//	n, err := reader.Read(buf)
	//	if err != nil && err != io.EOF {
	//		break
	//	}
	//	logs += string(buf[0:n])
	//}
	//fmt.Println(logs)
}
