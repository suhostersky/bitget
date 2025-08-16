package v2

import (
	"log"

	"github.com/suhostersky/bitget/types"
	"github.com/suhostersky/bitget/types/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init(apiKey, secretKey, passphrase string) *SpotAccountClient {
	bitgetRestClient, err := new(common.BitgetRestClient).Init(apiKey, secretKey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	p.BitgetRestClient = bitgetRestClient
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := types.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/assets", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}
