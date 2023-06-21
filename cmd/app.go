package main

import (
	"os"

	"github.com/Igorprostoff/go-ya-music"
	"golang.org/x/exp/slog"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	c := go_ya_music.Client{}
	err := c.Init(os.Getenv("TOKEN"),
		"", "ru", true, nil)
	if err != nil {
		log.Error("Unable to init client", err)
	}

	status, err := c.AccountStatus()
	if err != nil {
		log.Error("Unable to get account status ", err)
	}
	log.Info("Account status: ", status)

	userSettings, err := c.AccountSettings()
	if err != nil {
		log.Error("Unable to get account settings ", err)
	}
	log.Info("Account settings: ", userSettings)

	newUserSettingsMultiple, err := c.AccountSettingsSetMultiple(map[string]interface{}{"diskEnabled": true})
	if err != nil {
		log.Error("Unable to set account settings ", err)
	}
	log.Info("New Account settings: ", newUserSettingsMultiple)

	settings, err := c.Settings()
	if err != nil {
		log.Error("Unable to get settings ", err)
	}
	log.Info("Settings: ", settings)

	permissionAlerts, err := c.PermissionAlerts()
	if err != nil {
		log.Error("Unable to get permissionAlerts ", err)
	}
	log.Info("permissionAlerts: ", permissionAlerts)

	accountExperiments, err := c.AccountExperiments()
	if err != nil {
		log.Error("Unable to get accountExperiments ", err)
	}
	log.Info("accountExperiments: ", accountExperiments)

	promocode, err := c.ConsumePromocode("K6TS5BREDG", "en")
	if err != nil {
		log.Error("Unable to get promocode ", err)
	}
	log.Info("promocode: ", promocode)

	/*likes, err := c.GetLikesTracks(strconv.Itoa(status.Account.Uid))
	if err != nil {
		log.Error("Unable to get likes ", err)
	}*/

}
