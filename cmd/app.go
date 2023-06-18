package main

import (
	"log"
	"os"

	"github.com/Igorprostoff/go-ya-music"
)

func main() {
	c := go_ya_music.Client{}
	err := c.Init(os.Getenv("TOKEN"),
		"", "ru", true, nil)
	if err != nil {
		log.Fatal("Unable to init client", err)
	}

	status, err := c.AccountStatus()
	if err != nil {
		log.Fatal("Unable to get account status ", err)
	}
	log.Printf("Account status: %+v", status)

	userSettings, err := c.AccountSettings()
	if err != nil {
		log.Fatal("Unable to get account settings ", err)
	}
	log.Printf("Account settings: %+v", userSettings)

	newUserSettingsMultiple, err := c.AccountSettingsSetMultiple(map[string]interface{}{"diskEnabled": true})
	if err != nil {
		log.Fatal("Unable to set account settings ", err)
	}
	log.Printf("New Account settings: %+v", newUserSettingsMultiple)

	settings, err := c.Settings()
	if err != nil {
		log.Fatal("Unable to get settings ", err)
	}
	log.Printf("Settings: %+v", settings)

}
