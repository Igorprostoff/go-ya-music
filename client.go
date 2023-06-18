package go_ya_music

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/Igorprostoff/go-ya-music/internal/entities"
	"github.com/Igorprostoff/go-ya-music/internal/utils"
	"golang.org/x/exp/slog"
)

var (
	userAgent      = "Yandex-Music-API"
	headers        = map[string]string{"X-Yandex-Music-Client": "YandexMusicAndroid/24023231", "Authorization": "OAuth "}
	defaultTimeout = 5
)

type Client struct {
	token               string
	baseUrl             string
	language            string
	httpClient          *http.Client
	logger              *log.Logger
	reportUnknownFields bool
	device              string
}

func (c *Client) Init(
	token, baseUrl, language string,
	reportUnknownFields bool,
	httpClient *http.Client,
) error {
	if token == "" {
		return fmt.Errorf("unable to create client without token")
	}
	c.token = token
	headers["Authorization"] = headers["Authorization"] + token
	if !utils.InArray(language, []string{"en", "uz", "uk", "us", "ru", "kk", "hy"}) {
		c.language = "en"
	} else {
		c.language = language
	}

	if baseUrl == "" {
		c.baseUrl = "https://api.music.yandex.net"
	} else {
		c.baseUrl = baseUrl
	}
	if httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	c.reportUnknownFields = reportUnknownFields
	c.device = "os=Golang; os_version=; manufacturer=Igorprostoff;" +
		" model=Yandex Music API; clid=; device_id=random; uuid=random"

	c.logger = slog.NewLogLogger(slog.NewTextHandler(os.Stdout, nil), 0)
	return nil
}

func (c *Client) AccountStatus() (entities.Status, error) {
	var result entities.StatusResult
	url, err := url.JoinPath(c.baseUrl, "/account/status")
	if err != nil {
		return result.Result, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result.Result, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result.Result, err
}

func (c *Client) AccountSettings() (entities.UserSettings, error) {
	var result entities.UserSettingsResult
	url, err := url.JoinPath(c.baseUrl, "/account/settings")
	if err != nil {
		return result.Result, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result.Result, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result.Result, err
}

/*
func (c *Client) AccountSettingsSet(param, value string) (entities.UserSettingsResult, error) {
	var result entities.UserSettingsResult
	url, err := url.JoinPath(c.baseUrl, "/account/settings")
	if err != nil {
		return result, err
	}

	data := map[string]string{param: value}
	data_bin, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data_bin))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result, err
}
*/

func (c *Client) AccountSettingsSetMultiple(data map[string]interface{}) (entities.UserSettings, error) { //TODO: Get editable fields to test
	var result entities.UserSettingsResult
	url, err := url.JoinPath(c.baseUrl, "/account/settings")
	if err != nil {
		return result.Result, err
	}
	if !utils.MapKeysInStructKeys(data, entities.UserSettings{}) {
		return result.Result, fmt.Errorf("map key doesn't match any of user settings")
	}
	dataBin, err := json.Marshal(data)
	if err != nil {
		return result.Result, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(dataBin))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result.Result, err
}

func (c *Client) Settings() (entities.Settings, error) {
	var result entities.SettingsResult
	url, err := url.JoinPath(c.baseUrl, "/settings")
	if err != nil {
		return result.Result, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result.Result, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result.Result, err
}
