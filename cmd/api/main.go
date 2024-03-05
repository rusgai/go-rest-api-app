package main

import (
	"fmt"
	"log"
)

func Run() error {
	fmt.Println("Run server api to routing gin.")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalln()
	}
}
