package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Igorprostoff/go-ya-music/internal/entities"
	"github.com/Igorprostoff/go-ya-music/internal/utils"
)

var (
	userAgent = "Yandex-Music-API"
	headers   = map[string][]string{"X-Yandex-Music-Client": {"YandexMusicAndroid/24023231"}, "Authorization": {"OAuth "}}
	defaultTimeout = 5
)

type Client struct {
	token    string
	baseUrl  string
	language string
	httpClient          *http.Client
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
	headers["Authorization"] = []string{headers["Authorization"][0] + token}
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
	return nil
}

func (c *Client) AccountStatus() (entities.StatusResult, error) {
	var result entities.StatusResult
	strUrl, err := url.JoinPath(c.baseUrl, "/account/status")
	if err != nil {
		return result, err
	}
	url, err := url.Parse(strUrl)
	if err != nil {
		return result, err
	}

	req := http.Request{Method: http.MethodGet, Header: headers, URL: url}
	resp, err := c.httpClient.Do(&req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
	}
	err = result.ParseFromReader(resp.Body)
	return result, err
}
