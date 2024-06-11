package clTron

import (
	"github.com/lionhart580230/clUtil/clLog"
	"testing"
)

func TestCreateTronAccount(t *testing.T) {
	addr, key := CreateTronAccount()
	clLog.Info("addr: %v", addr)
	clLog.Info("key: %v", key)
}

func TestGetBlockInfo(t *testing.T) {
	err, blockInfo := GetBlockInfo("")
	if err != nil {
		clLog.Error("获取区块信息错误:%v", err)
		return
	}
	clLog.Info("区块: %+v", *blockInfo)
}

func TestQueryTransaction(t *testing.T) {
	SetTest(1)
	err, info := QueryTransaction("e9f44ce91c0c3bd7b2e6c4b97e81669ca364cd7d179f3997878c485d848cdb33")
	if err != nil {
		clLog.Error("查询交易记录错误:%v", err)
		return
	}
	if info == nil {
		return
	}

	clLog.Info("订单信息: %+v", info)
}

/*
*
15:22:50 clTron.go:201[Debug] 获取账号信息:{"address": "41c24cbac543646d60ff414cf55135500c30be5531","balance": 30,"create_time": 1716876252000,"latest_opration_time": 1716880542000,"free_net_usage": 338,"latest_consume_free_time": 1716879297000,"net_window_size": 28800000,"net_window_optimized": true,"account_resource": {"latest_consume_time_for_energy": 1716880542000,"energy_window_size": 28800000,"energy_window_optimized": true},"owner_permission": {"permission_name": "owner","threshold": 1,"keys": [{"address": "41c24cbac543646d60ff414cf55135500c30be5531","weight": 1}]},"active_permission": [{"type": "Active","id": 2,"permission_name": "active","threshold": 1,"operations": "7fff1fc0033ec30f000000000000000000000000000000000000000000000000","keys": [{"address": "41c24cbac543646d60ff414cf55135500c30be5531","weight": 1}]}],"frozenV2": [{},{"type": "ENERGY"},{"type": "TRON_POWER"}],"asset_optimized": true}
*/
func TestGetAccountInfo(t *testing.T) {
	GetAccountInfo("TTga1YAxuUTbJVdLqsVjU7FNqXWbTPLfKB")
}

func TestGetAccountResources(t *testing.T) {
	GetAccountResources("TS4yRVvuozWpKScEC4SKjEq7Y2g8QYEyx5")
}
