package clAliSDK

import (
	"github.com/xiaolan580230/clUtil/clLog"
	"testing"
)

func TestApiGetBalance(t *testing.T) {
	err, data := ApiGetBalance("oss的access key", "oss的secret key")
	if err != nil {
		clLog.Error("获取余额错误: %v", err)
		return
	}

	clLog.Debug("账号: %v 余额: %v", "video-best", data.Data.AvailableCashAmount)
}
