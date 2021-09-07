package clConfig

import (
	"github.com/xiaolan580230/clUtil/clSuperMap"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var mConfig *clSuperMap.SuperMap

func init() {
	mConfig = clSuperMap.NewSuperMap()
}


// 通过文件加载配置
// 调用多次后面的会覆盖前面
func LoadFromFile(_filename string) {
	buffer, err := ioutil.ReadFile(_filename)
	if err != nil {
		return
	}

	var confLines = strings.Split(string(buffer), "\n")
	for _, line := range confLines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "--") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx <= 0 {
			continue
		}
		mConfig.Add(line[:idx], line[idx+1:])
	}
}



// 获取Uint32配置
func GetUint32(_key string, _def uint32) uint32 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseInt(os.Getenv(_key), 10, 32)
		if e != nil {
			return _def
		}
		return uint32(v)
	}
	return mConfig.GetUInt32(_key, _def)
}


// 获取Uint64配置
func GetUint64(_key string, _def uint64) uint64 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseInt(os.Getenv(_key), 10, 64)
		if e != nil {
			return _def
		}
		return uint64(v)
	}
	return mConfig.GetUInt64(_key, _def)
}


// 获取int32配置
func GetInt32(_key string, _def int32) int32 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseInt(os.Getenv(_key), 10, 32)
		if e != nil {
			return _def
		}
		return int32(v)
	}
	return mConfig.GetInt32(_key, _def)
}


// 获取int64配置
func GetInt64(_key string, _def int64) int64 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseInt(os.Getenv(_key), 10, 64)
		if e != nil {
			return _def
		}
		return int64(v)
	}
	return mConfig.GetInt64(_key, _def)
}


// 获取bool配置
func GetBool(_key string, _def bool) bool {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseBool(os.Getenv(_key))
		if e != nil {
			return _def
		}
		return v
	}
	return mConfig.GetBool(_key, _def)
}


// 获取Str配置
func GetStr(_key string, _def string) string {
	if os.Getenv(_key) != "" {
		return os.Getenv(_key)
	}
	return mConfig.GetStr(_key, _def)
}


// 获取Float32配置
func GetFloat32(_key string, _def float32) float32 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseFloat(os.Getenv(_key), 32)
		if e != nil {
			return _def
		}
		return float32(v)
	}
	return mConfig.GetFloat32(_key, _def)
}


// 获取Float32配置
func GetFloat64(_key string, _def float64) float64 {
	if os.Getenv(_key) != "" {
		v, e := strconv.ParseFloat(os.Getenv(_key), 64)
		if e != nil {
			return _def
		}
		return v
	}
	return mConfig.GetFloat64(_key, _def)
}

