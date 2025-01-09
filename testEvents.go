package main

import (
	"context"
	"fmt"
	"github.com/bigartists/Modi/client"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {
	list, err := client.K8sClient.CoreV1().Events("groot").List(
		context.Background(), v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range list.Items {
		fmt.Println(item.Name, item.Type,
			item.Reason,
			item.Message,
			item.Namespace,
			item.InvolvedObject,
		)
	}
}
