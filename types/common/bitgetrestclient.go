package common

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/suhostersky/bitget/config"
	"github.com/suhostersky/bitget/constants"
	"github.com/suhostersky/bitget/types"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
}

func (p *BitgetRestClient) Init(apiKey, secretKey, passphrase string) (*BitgetRestClient, error) {
	if apiKey == "" || secretKey == "" {
		return nil, errors.New("api key or secret key is empty")
	}
	p.ApiKey = apiKey
	p.ApiSecretKey = secretKey
	p.BaseUrl = config.BaseUrl
	p.Passphrase = passphrase
	p.Signer = new(Signer).Init(secretKey)
	p.HttpClient = http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}
	return p, nil
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := types.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	if constants.RSA == config.SignType {
		sign = p.Signer.SignByRSA(constants.POST, uri, params, timesStamp)
	}
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	types.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := types.TimesStamp()
	body := types.BuildGetParams(params)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	types.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
