package clCrypt

import (
	"github.com/lionhart580230/clUtil/clCommon"
	"github.com/lionhart580230/clUtil/clLog"
	"testing"
)

func TestAesCBCEncode(t *testing.T) {
	for i := 0; i < 100; i++ {
		// 随机字符串
		var randomStr = clCommon.GenNonceStr(32)
		var aesKey = string(RandomBlock(32))
		var iv = RandomBlock(16)

		var cryptData = AesCBCEncode(randomStr, aesKey, string(iv))
		var unCryptData = AesCBCDecode([]byte(cryptData), []byte(aesKey), iv)

		if unCryptData != randomStr {
			clLog.Error("加密解密结果不对!!")
		}
	}

}
