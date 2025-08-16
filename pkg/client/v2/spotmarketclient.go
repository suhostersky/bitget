package v2

import (
	"log"

	"github.com/suhostersky/bitget/types"
	"github.com/suhostersky/bitget/types/common"
)

type SpotMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotMarketClient) Init(apiKey, secretKey, passphrase string) *SpotMarketClient {
	bitgetRestClient, err := new(common.BitgetRestClient).Init(apiKey, secretKey, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	p.BitgetRestClient = bitgetRestClient
	return p
}

func (p *SpotMarketClient) Coins() (string, error) {
	params := types.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/coins", params)
	return resp, err
}

func (p *SpotMarketClient) Symbols(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
	return resp, err
}

func (p *SpotMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/fills", params)
	return resp, err
}

func (p *SpotMarketClient) Orderbook(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/orderbook", params)
	return resp, err
}

func (p *SpotMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/tickers", params)
	return resp, err
}

func (p *SpotMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/candles", params)
	return resp, err
}
