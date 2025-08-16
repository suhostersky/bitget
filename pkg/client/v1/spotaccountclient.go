package v1

import (
	"log"

	"github.com/suhostersky/bitget/types"
	"github.com/suhostersky/bitget/types/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init(apikey, secretKey, passphrase string) *SpotAccountClient {
	bitgetRestClient, err := new(common.BitgetRestClient).Init(apikey, secretKey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	p.BitgetRestClient = bitgetRestClient
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := types.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/getInfo", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/assets-lite", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/spot/v1/account/transferRecords", params)
	return resp, err
}
