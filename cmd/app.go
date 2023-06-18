package main

import (
	"log"

	"github.com/Igorprostoff/go-ya-music"
)

func main() {
	c := go_ya_music.Client{}
	err := c.Init("",
		"", "ru", true, nil)
	if err != nil {
		log.Fatal("Unable to init client", err)
	}

	status, err := c.AccountStatus()
	if err != nil {
		log.Fatal("Unable to get account status ", err)
	}
	log.Printf("Account status: %+v", status)

	settings, err := c.AccountSettings()
	if err != nil {
		log.Fatal("Unable to get account settings ", err)
	}
	log.Printf("Account settings: %+v", settings)

	newSettingsMultiple, err := c.AccountSettingsSetMultiple(map[string]interface{}{"diskEnabled": true})
	if err != nil {
		log.Fatal("Unable to set account settings ", err)
	}
	log.Printf("New Account settings: %+v", newSettingsMultiple)

}
