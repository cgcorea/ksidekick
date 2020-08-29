package main

import (
	"fmt"
	"log"

	"github.com/cgcorea/ksidekick/internal/kannel"
)

func main() {
	fmt.Println("ksidekick: Kannel Sidekick utility")
	k := kannel.NewClient("localhost", 4103, "sender", "sender")

	resp, err := k.SendText("1010", "50499821977", "Hello world")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
