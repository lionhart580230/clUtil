package clCrypt

import (
	"crypto/sha1"
	"fmt"
	"github.com/lionhart580230/clUtil/clCommon"
	"github.com/lionhart580230/clUtil/clLog"
	"golang.org/x/crypto/pbkdf2"
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

func TestAesCBCDecode(t *testing.T) {
	a := pbkdf2.Key([]byte("1233123"), []byte("hello"), 1000, 32, sha1.New)
	fmt.Printf("加密: %v\n", a)
	fmt.Printf("长度: %v\n", len(a))
}
