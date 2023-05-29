package wecom

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
)

const BASE_URL = "https://qyapi.weixin.qq.com"

type HttpMethod string

var (
	GET  HttpMethod = http.MethodGet
	POST HttpMethod = http.MethodPost
)

type Common struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Api struct {
	method HttpMethod
	path   string
}

var GetTokenApi = Api{
	method: GET,
	path:   "/cgi-bin/gettoken",
}

type GetTokenApiResp struct {
	Common
	AccessToken string  `json:"access_token"`
	ExpiresIn   big.Int `json:"expires_in"`
}

func invoke(api Api, params map[string]string) ([]byte, error) {
	client := http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), string(api.method), genUrl(api, params), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(fmt.Sprintf("invoke wecom request %s response %s\n", params, string(bytes)))
	return bytes, nil
}

func genUrl(api Api, params map[string]string) string {
	if api.method == POST {
		return BASE_URL + api.path
	}
	querys := []string{}
	for k, v := range params {
		querys = append(querys, k+"="+v)
	}
	return BASE_URL + api.path + "?" + strings.Join(querys, "&")
}

func GetToken(corpid string, corpsecret string) *GetTokenApiResp {
	params := map[string]string{
		"corpid":     corpid,
		"corpsecret": corpsecret,
	}
	respBody, err := invoke(GetTokenApi, params)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	resp := &GetTokenApiResp{}
	_ = json.Unmarshal(respBody, resp)
	return resp
}
