package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cgcorea/ksidekick/internal/debug"
	"github.com/cgcorea/ksidekick/internal/kannel"
)

func main() {
	fmt.Printf("ksidekick: Kannel Sidekick utility\n\n")
	k := kannel.NewClient("localhost", 4103, "sender", "sender")

	options := kannel.Options{SMSC: "tigo-hn-smsc1"}

	debug.Inspect(options, os.Stdout)
	resp, err := k.Send("1010", "50499821977", "Hello world", &options)

	if err != nil {
		log.Fatal(err)
	}

	debug.Inspect(resp, os.Stdout)
}
