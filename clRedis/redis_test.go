package clRedis

import (
	"github.com/lionhart580230/clUtil/clLog"
	"testing"
)

func TestNew(t *testing.T) {

	redis, err := New("127.0.0.1:6379", "", "")
	if err != nil {
		clLog.Error("连接错误: %v", err)
		return
	}

	if redis.SetEx("hello", "1", 100) {
		clLog.Debug("设置成功!!")
	} else {
		clLog.Debug("设置失败!!")
	}

	if redis.SetExpire("hello", 200) {
		clLog.Debug("设置时效成功!!")
	} else {
		clLog.Debug("设置时效失败!!")
	}

	clLog.Debug("设置redis...")
}

func TestRedisObject_Listen(t *testing.T) {
	redis, err := New("127.0.0.1:6379", "", "")
	if err != nil {
		clLog.Error("连接错误: %v", err)
		return
	}

	var str = make(chan string)

	go redis.Subscribe("queue_1", str)
	for msg := range str {
		clLog.Debug("收到消息: %v", msg)
	}
}
