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
