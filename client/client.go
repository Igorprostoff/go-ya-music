package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/Igorprostoff/go-ya-music/account"
	"github.com/Igorprostoff/go-ya-music/utils"
)

var (
	user_agent      = "Yandex-Music-API"
	headers         = map[string][]string{"X-Yandex-Music-Client": {"YandexMusicAndroid/24023231"}}
	default_timeout = 5
)

type Client struct {
	token                 string
	base_url              string
	language              string
	http_client           *http.Client
	report_unknown_fields bool
	device                string
}

func (c *Client) Init(
	token, base_url, language string,
	report_unknown_fields bool,
	http_client *http.Client,
) error {
	if token == "" {
		return fmt.Errorf("unable to create client without token")
	}
	c.token = token

	if !utils.In_array(language, []string{"en", "uz", "uk", "us", "ru", "kk", "hy"}) {
		c.language = "en"
	} else {
		c.language = language
	}

	if base_url == "" {
		c.base_url = "https://api.music.yandex.net"
	} else {
		c.base_url = base_url
	}
	if http_client == nil {
		c.http_client = http.DefaultClient
	}
	c.report_unknown_fields = report_unknown_fields
	c.device = "os=Golang; os_version=; manufacturer=Igorprostoff;" +
		" model=Yandex Music API; clid=; device_id=random; uuid=random"
	return nil
}

func (c *Client) Account_status() (account.Status, error) {
	var res account.Status
	url, err := url.Parse(path.Join(c.base_url, "/account/status"))
	if err != nil {
		return res, err
	}
	req := http.Request{Method: http.MethodGet, Header: headers, URL: url}
	resp, err := c.http_client.Do(&req)
	if err != nil {
		return res, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, &res)
}
