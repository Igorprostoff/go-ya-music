package main

import (
	"log"

	"github.com/Igorprostoff/go-ya-music/client"
)

func main() {
	c := client.Client{}
	err := c.Init("",
		"", "ru", true, nil)
	if err != nil {
		log.Fatal("Unable to init client", err)
	}
	status, err := c.AccountStatus()
	if err != nil {
		log.Fatal("Unable to get account status ", err)
	}
	log.Printf("Account status: %+v", status.Result)
}
