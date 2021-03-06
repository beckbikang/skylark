package main

import (
	"log"
	"skylark/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%+v", err)
	}
}
