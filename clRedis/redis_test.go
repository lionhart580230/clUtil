package clRedis

import (
	"github.com/xiaolan580230/clUtil/clLog"
	"testing"
)

func TestNew(t *testing.T) {

	redis, err := New("127.0.0.1:6379", "", "test")
	if err != nil {
		clLog.Error("连接错误: %v", err)
		return
	}

	err, list := redis.LPOPWait("testKey", 3)
	clLog.Error("err: %v list: %+v", err, list)
}