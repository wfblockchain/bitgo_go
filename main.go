package bitgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/google/go-querystring/query"
)

type BitGo struct {
	debug   bool
	timeout time.Duration
	host    string
	token   string
	coin    string
}

type ListParams struct {
	Limit     int    `url:"limit,omitempty"`
	PrevId    string `url:"prevId,omitempty"`
	AllTokens bool   `url:"allTokens,omitempty"`
}

func New(env string, timeout time.Duration) (b *BitGo, err error) {
	if env == "" {
		return nil, errors.New("empty env")
	}
	switch env {
	case "express_test":
		env = "http://localhost:3080"
	case "test":
		env = "https://app.bitgo-test.com"
		//env = "https://test.bitgo.com"
	case "prod":
		env = "https://app.bitgo.com"
		//env = "https://www.bitgo.com"
	}
	return &BitGo{
		host:    env + "/api/v2",
		timeout: timeout,
	}, nil
}

func (b *BitGo) clone() *BitGo {
	return &BitGo{
		host:    b.host,
		token:   b.token,
		coin:    b.coin,
		timeout: b.timeout,
		debug:   b.debug,
	}
}

func (b *BitGo) Token(token string) *BitGo {
	c := b.clone()
	c.token = token
	return c
}

func (b *BitGo) Coin(coin string) *BitGo {
	c := b.clone()
	c.coin = coin
	return c
}

func (b *BitGo) Debug(debug bool) *BitGo {
	c := b.clone()
	c.debug = debug
	return c
}

func (b *BitGo) get(url string, params interface{}, responce interface{}) (err error) {
	if params != nil {
		_, err = valid.ValidateStruct(params)
		if err != nil {
			return
		}

		v, _ := query.Values(params)
		url = url + "?" + v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, b.host+"/"+url, nil)
	if err != nil {
		return
	}

	return b.request(req, responce)
}

func (b *BitGo) modify(method string, url string, params interface{}, responce interface{}) (err error) {
	var req *http.Request

	if params != nil {
		var body *bytes.Buffer

		_, err = valid.ValidateStruct(params)
		if err != nil {
			return
		}

		body = new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(params)
		if err != nil {
			return
		}
		req, err = http.NewRequest(method, b.host+"/"+url, body)
		if err != nil {
			return
		}

		req.Header.Add("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, b.host+"/"+url, nil)
		if err != nil {
			return
		}
	}

	return b.request(req, responce)
}

func (b *BitGo) post(url string, params interface{}, responce interface{}) (err error) {
	return b.modify(http.MethodPost, url, params, responce)
}

func (b *BitGo) put(url string, params interface{}, responce interface{}) (err error) {
	return b.modify(http.MethodPut, url, params, responce)
}

func (b *BitGo) request(req *http.Request, responce interface{}) (err error) {
	req.Header.Add("Authorization", "Bearer "+b.token)

	client := &http.Client{
		Timeout: b.timeout,
	}

	if b.debug {
		log.Println(req.Method, req.URL.EscapedPath())
	}

	r, err := client.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	if r.StatusCode != http.StatusOK {
		berr := Error{}

		err = json.Unmarshal(resp, &berr)
		if err != nil {
			return err
		}

		return berr
	}

	if b.debug {
		log.Println(r.Status, string(resp))
	}

	err = json.Unmarshal(resp, &responce)
	return
}
