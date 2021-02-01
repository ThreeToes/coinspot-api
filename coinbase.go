//coinspot-api, an API client library for coinspot.com.au
//Copyright (C) 2021 Stephen Gream
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Coinbase represents the API
type CoinspotApi interface {
	//LatestPrices gets the prices of coins
	LatestPrices() (*LatestPricesResponse, error)
	//ListOpenOrders lists all open orders
	ListOpenOrders(coinType string) (*ListOpenOrdersResponse, error)
	ListOrderHistory(coinType string) (*ListOrderHistoryResponse, error)
	//DepositCoins generates a deposit address
	DepositCoins(coinType string) (*DepositCoinsResponse, error)
	QuickBuyQuote(coinType string, amount float64) (*QuickBuyQuoteResponse, error)
	QuickSellQuote(coinType string, amount float64) (*QuickSellQuoteResponse, error)
	ListBalances() (*ListBalancesResponse, error)
	//ListMyOrders lists all your open orders
	ListMyOrders() (*ListMyOrdersResponse, error)
	PlaceBuyOrder(coinType string, amount float64, rate float64) (*PlaceBuyOrderResponse, error)
	PlaceSellOrder(coinType string, amount float64, rate float64) (*PlaceSellOrderResponse, error)
	CancelBuyOrder(id string) (*CancelBuyOrderResponse, error)
	CancelSellOrder(id string) (*CancelSellOrderResponse, error)
}

type coinSpotApi struct {
	httpClient *http.Client
	secret string
	key string
}

func (c *coinSpotApi) LatestPrices() (*LatestPricesResponse, error) {
	var ret LatestPricesResponse
	var err error

	resp, err := c.httpClient.Get(LatestPricesUrl)
	if err != nil {
		return nil, err
	}
	bodyString, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyString, &ret)

	return &ret, err
}

func (c *coinSpotApi) ListOpenOrders(coinType string) (*ListOpenOrdersResponse, error) {
	req := &listOpenOrdersRequest{
		CoinType:        coinType,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(ListOpenOrdersUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret ListOpenOrdersResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) ListOrderHistory(coinType string) (*ListOrderHistoryResponse, error) {
	req := &listOrderHistoryRequest {
		CoinType:        coinType,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(ListOrderHistoryUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret ListOrderHistoryResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) DepositCoins(coinType string) (*DepositCoinsResponse, error) {
	req := &depositCoinsRequest{
		CoinType:        coinType,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(DepositCoinsUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret DepositCoinsResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) QuickBuyQuote(coinType string, amount float64) (*QuickBuyQuoteResponse, error) {
	req := &quickBuyQuoteRequest{
		CoinType:        coinType,
		Amount: amount,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(QuickBuyQuoteUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret QuickBuyQuoteResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) QuickSellQuote(coinType string, amount float64) (*QuickSellQuoteResponse, error) {
	req := &quickSellQuoteRequest{
		CoinType:        coinType,
		Amount: amount,
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(QuickBuyQuoteUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret QuickSellQuoteResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) ListBalances() (*ListBalancesResponse, error) {
	req := &listBalancesRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(ListMyBalancesUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret ListBalancesResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) ListMyOrders() (*ListMyOrdersResponse, error) {
	req := &listMyOrdersRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := c.doRequest(ListMyOrdersUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret ListMyOrdersResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) PlaceBuyOrder(coinType string, amount float64, rate float64) (*PlaceBuyOrderResponse, error) {
	req := &placeBuyOrderRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
		CoinType: coinType,
		Amount: amount,
		Rate: rate,
	}
	resp, err := c.doRequest(PlaceBuyOrderUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret PlaceBuyOrderResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) PlaceSellOrder(coinType string, amount float64, rate float64) (*PlaceSellOrderResponse, error) {
	req := &placeSellOrderRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
		CoinType: coinType,
		Amount: amount,
		Rate: rate,
	}
	resp, err := c.doRequest(PlaceSellOrderUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret PlaceSellOrderResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) CancelBuyOrder(id string) (*CancelBuyOrderResponse, error) {
	req := &cancelBuyOrderRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
		Id: id,
	}
	resp, err := c.doRequest(CancelBuyOrderUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret CancelBuyOrderResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) CancelSellOrder(id string) (*CancelSellOrderResponse, error) {
	req := &cancelSellOrderRequest{
		Nonce: fmt.Sprintf("%d", time.Now().Unix()),
		Id: id,
	}
	resp, err := c.doRequest(CancelSellOrderUrl, req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret CancelSellOrderResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}

func (c *coinSpotApi) doRequest(url string, body interface{}) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha512.New, []byte(c.secret))
	mac.Write(b)
	sigStr := hex.EncodeToString(mac.Sum(nil))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("sign", sigStr)
	req.Header.Add("key", c.key)

	resp, err := c.httpClient.Do(req)

	return resp, err
}

func NewCoinSpotApi(key, secret string) CoinspotApi {
	c := &http.Client{}
	return &coinSpotApi{
		httpClient: c,
		secret:     secret,
		key:        key,
	}
}