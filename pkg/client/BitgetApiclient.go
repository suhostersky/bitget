package client

import (
	"log"

	"github.com/suhostersky/bitget/types"
	"github.com/suhostersky/bitget/types/common"
)

type BitgetApiClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *BitgetApiClient) Init(apikey, secretKey, passphrase string) *BitgetApiClient {
	bitgetRestClient, err := new(common.BitgetRestClient).Init(apikey, secretKey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	p.BitgetRestClient = bitgetRestClient
	return p
}

func (p *BitgetApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := types.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}
