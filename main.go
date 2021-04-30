/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cgcorea/ksidekick/kannel"

	"github.com/cgcorea/ksidekick/internal/debug"
	"github.com/cgcorea/ksidekick/internal/server"
)

func main() {
	// sendSMS()

	srv := server.NewServer()
	srv.SetRoutes()

	log.Println("Listing on port 7000")
	err := http.ListenAndServe("localhost:7000", srv.Router)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

	// cmd.Execute()

	// sendSMS()
}

func sendSMS() {
	fmt.Printf("ksidekick: Kannel Sidekick utility\n\n")
	k := kannel.NewClient("localhost", 4103, "sender", "sender")

	options := []func(*kannel.Request){kannel.SMSC("tigo-hn-1")}
	options = append(
		options,
		kannel.DLRUrl("http://example.com/?from=%d"),
	)

	req, err := k.NewRequest(
		"1010",
		"50499821977",
		"Hello world",
		kannel.Priority(1),
		kannel.DLRUrl("http://example.com/?from=%d"),
		kannel.DLRMask(3),
	)

	if err != nil {
		log.Fatal(err)
	}

	debug.Inspect(req.Header, os.Stdout)

	resp, err := k.Send(req)
	if err != nil {
		log.Fatal(err)
	}

	debug.Inspect(resp, os.Stdout)
}
