package clCrypt

import (
	"fmt"
	"testing"
)

func TestAesCBCEncode(t *testing.T) {
	var data = `{"k": "hello"}`
	var aesKey = string(RandomBlock(32))
	var iv = RandomBlock(32)
	var cryptData = AesCBCEncode(data, aesKey, string(iv))
	fmt.Printf("加密后内容: %v\n", cryptData)

	var unCryptData = AesCBCDecode([]byte(cryptData), []byte(aesKey), iv)
	fmt.Printf("解密后内容: %v\n", unCryptData)
}