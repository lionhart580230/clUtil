package clConfig

import (
	"github.com/lionhart580230/clUtil/clSuperMap"
	"io/ioutil"
	"os"
	"strings"
)

var mConfig map[string]*clSuperMap.SuperMap
var aConfig map[string][]string

func init() {
	mConfig = make(map[string]*clSuperMap.SuperMap)
	aConfig = make(map[string][]string)
}

// 通过文件加载配置
// 调用多次后面的会覆盖前面
// @Param _filename 配置文件所在路径
// @Param _overWrite 是否覆盖之前读取的内容， true为覆盖，反之不覆盖
func LoadFromFile(_filename string, _overWrite bool) error {
	buffer, err := ioutil.ReadFile(_filename)
	if err != nil {
		return err
	}

	section := "GLOBAL"
	if _overWrite {
		mConfig = make(map[string]*clSuperMap.SuperMap)
		mConfig[section] = clSuperMap.NewSuperMap()
		aConfig = make(map[string][]string)
		aConfig[section] = make([]string, 0)
	}

	var confLines = strings.Split(string(buffer), "\n")
	for _, line := range confLines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "--") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = strings.ToUpper(strings.TrimPrefix(strings.TrimSuffix(line, "]"), "["))
			mConfig[section] = clSuperMap.NewSuperMap()
			aConfig[section] = make([]string, 0)
			continue
		}

		idx := strings.Index(line, "=")
		if idx <= 0 {
			// 可能是数组
			aConfig[section] = append(aConfig[section], line)
			continue
		}
		mConfig[section].Add(strings.ToUpper(line[:idx]), line[idx+1:])
	}
	return nil
}

// 从环境变量中加载配置
func LoadFromENV(_section string, _key string) {
	_section = strings.ToUpper(_section)
	_key = strings.ToUpper(_key)
	val := os.Getenv(_section + "_" + _key)
	if val == "" {
		return
	}
	_, isOK := mConfig[_section]
	if !isOK {
		mConfig[_section] = clSuperMap.NewSuperMap()
	}
	mConfig[_section].Add(_key, val)
}

// 获取Uint32配置
func GetUint32(_key string, _def uint32) uint32 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetUInt32(_key[idx+1:], _def)
}

// 获取Uint64配置
func GetUint64(_key string, _def uint64) uint64 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetUInt64(_key[idx+1:], _def)
}

// 获取int32配置
func GetInt32(_key string, _def int32) int32 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetInt32(_key[idx+1:], _def)
}

// 获取int64配置
func GetInt64(_key string, _def int64) int64 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetInt64(_key[idx+1:], _def)
}

// 获取bool配置
func GetBool(_key string, _def bool) bool {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetBool(_key[idx+1:], _def)
}

// 获取Str配置
func GetStr(_key string, _def string) string {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {

		return _def
	}
	return configItem.GetStr(_key[idx+1:], _def)
}

// 获取Float32配置
func GetFloat32(_key string, _def float32) float32 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetFloat32(_key[idx+1:], _def)
}

// 获取Float32配置
func GetFloat64(_key string, _def float64) float64 {
	_key = strings.ToUpper(_key)
	idx := strings.Index(_key, ".")
	if idx <= 0 {
		return _def
	}
	configItem, exists := mConfig[_key[:idx]]
	if !exists {
		return _def
	}
	return configItem.GetFloat64(_key[idx+1:], _def)
}

// 获取Arr配置
func GetStrArray(_key string) []string {
	configItem, exists := aConfig[strings.ToUpper(_key)]
	if !exists {
		return []string{}
	}
	return configItem
}

// 获取Arr配置
func GetMap(_key string) map[string]string {
	configItem, exists := mConfig[strings.ToUpper(_key)]
	if !exists {
		return map[string]string{}
	}
	return configItem.GetMap()
}
