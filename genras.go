package main

import (
	"github.com/bigartists/Modi/src/utils"
	"log"
)

func main() {
	err := utils.GenRSAPubAndPri(1024, "./resources/pem")
	if err != nil {
		log.Fatal(err)
	}
}
