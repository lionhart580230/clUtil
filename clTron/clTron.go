package clTron

import (
	"encoding/json"
	"github.com/JFJun/trx-sign-go/genkeys"
	"github.com/lionhart580230/clUtil/clCommon"
	"github.com/lionhart580230/clUtil/clHttpClient"
	"github.com/lionhart580230/clUtil/clJson"
	"github.com/lionhart580230/clUtil/clLog"
)

var isTest = uint32(0)
var phpUrl = "http://lott_php_wallet/"

// 设置测试
func SetTest(_test uint32) {
	isTest = _test
}

// 设置php钱包地址
func setPHPWalletUrl(_url string) {
	phpUrl = _url
}

func GetUSDTContractAddr() string {
	if isTest == 1 {
		return "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj"
	} else {
		return "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	}
}

func GetTronApiUrl() string {
	if isTest == 1 {
		return "https://nile.trongrid.io/"
	}
	return "https://api.trongrid.io/"
}

type SendTrxResp struct {
	Result bool `json:"result"`
}

// 激活地址
func SendTrx(_fromAddr, _fromPriKey, _wallet string, _amount float64) (error, *SendTrxResp) {
	hc := clHttpClient.NewClient(phpUrl)
	hc.SetContentType(clHttpClient.ContentTypeForm)
	hc.AddParam("ac", "send_trx")
	hc.AddParam("from_address", _fromAddr)
	hc.AddParam("from_prikey", _fromPriKey)
	hc.AddParam("to_address", _wallet)
	hc.AddParam("amount", _amount)
	hc.AddParam("test", isTest)
	resp, err := hc.Do()
	if err != nil {
		clLog.Error("请求错误: %v", err)
		return err, nil
	}

	var respObj SendTrxResp
	err = json.Unmarshal([]byte(resp.Body), &respObj)
	clLog.Debug("交易trx地址: %v -> %v 金额: %0.4f 回应:%v", _fromAddr, _wallet, _amount, resp.Body)
	return err, &respObj
}

type BalanceResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Balance string `json:"balance"`
	} `json:"data"`
}

// 获取余额错误
func GetTrc20Balance(_addr string) (error, float64) {
	hc := clHttpClient.NewClient(phpUrl)
	hc.SetContentType(clHttpClient.ContentTypeForm)
	hc.AddParam("ac", "balance_trc20")
	hc.AddParam("address", _addr)
	hc.AddParam("test", isTest)
	resp, err := hc.Do()
	if err != nil {
		return err, 0
	}
	var data BalanceResp

	if err := json.Unmarshal([]byte(resp.Body), &data); err != nil {
		clLog.Error("json: (%v)", resp.Body)
		return err, 0
	}
	return nil, clCommon.Float64(data.Data.Balance) / 1000000
}

type SendUSDTResp struct {
	Result     bool   `json:"result"`
	Visible    bool   `json:"visible"`
	TxID       string `json:"txID"`
	RawDataHex string `json:"raw_data_hex"`
}

// 发送TRC20
func SendTrc20(_fromAddr, _fromPrikey, _toAddr string, _amount float64) (error, *SendUSDTResp) {
	hc := clHttpClient.NewClient("http://lott_php_wallet/")
	hc.SetContentType(clHttpClient.ContentTypeForm)
	hc.AddParam("ac", "send_usdt")
	hc.AddParam("from_address", _fromAddr)
	hc.AddParam("from_prikey", _fromPrikey)
	hc.AddParam("to_address", _toAddr)
	hc.AddParam("amount", _amount)
	hc.AddParam("test", isTest)
	resp, err := hc.Do()
	if err != nil {
		clLog.Error("发送TRC20错误:%v", err)
		return err, nil
	}
	var respObj SendUSDTResp
	if err := json.Unmarshal([]byte(resp.Body), &respObj); err != nil {
		return err, nil
	}
	return nil, &respObj
}

func CreateTronAccount() (string, string) {
	key, address := genkeys.GenerateKey()
	return address, key
}

type QueryTransactionResp struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	}
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
	} `json:"raw_data"`
	Timestamp     uint64 `json:"timestamp"`
	FeeLimit      uint64 `json:"fee_limit"`
	Expiration    uint64 `json:"expiration"`
	RefBlockBytes string `json:"ref_block_bytes"`
	RefBlockHash  string `json:"ref_block_hash"`
}

// 查询交易信息
func QueryTransaction(_txId string) (error, *QueryTransactionResp) {
	client := clHttpClient.NewClient(GetTronApiUrl() + "wallet/gettransactionbyid")
	client.SetContentType(clHttpClient.ContentJson)
	client.SetBody(clJson.CreateBy(clJson.M{
		"value":   _txId,
		"visible": true,
	}).ToStr())
	resp, err := client.Do()
	if err != nil {
		return err, nil
	}

	var data QueryTransactionResp
	if err := json.Unmarshal([]byte(resp.Body), &data); err != nil {
		return err, nil
	}
	return nil, &data
}

// 获取区块信息，如果区块ID为空则查询最新区块
func GetBlockInfo(_blockId string) (error, *Block) {

	httpClient := clHttpClient.NewClient(GetTronApiUrl() + "wallet/getblock")
	httpClient.AddHeader("accept", "application/json")
	httpClient.SetContentType(clHttpClient.ContentJson)
	httpClient.SetBody(clJson.CreateBy(clJson.M{
		"detail":    true,
		"id_or_num": _blockId,
	}).ToStr())
	resp, err := httpClient.Do()
	if err != nil {
		return err, nil
	}

	var block Block
	err = json.Unmarshal([]byte(resp.Body), &block)
	if err != nil {
		return err, nil
	}
	return nil, &block
}

type AccountInfo struct {
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
}

// 获取TRX余额
func GetAccountInfo(_address string) (error, *AccountInfo) {

	hex, _ := genkeys.AddressB58ToHex(_address)

	httpClient := clHttpClient.NewClient(GetTronApiUrl() + "walletsolidity/getaccount")
	httpClient.AddHeader("accept", "application/json")
	httpClient.SetContentType(clHttpClient.ContentJson)
	httpClient.SetBody(clJson.CreateBy(clJson.M{
		"address": hex,
	}).ToStr())
	resp, err := httpClient.Do()
	if err != nil {
		return err, nil
	}
	clLog.Debug("获取账号信息:%v", resp.Body)

	var accountInfo AccountInfo
	err = json.Unmarshal([]byte(resp.Body), &accountInfo)
	if err != nil {
		return err, nil
	}
	return nil, &accountInfo
}
