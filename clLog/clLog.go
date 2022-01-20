package clLog

import (
	"fmt"
	"github.com/xiaolan580230/clUtil/clTime"
	"runtime"
	"strings"
	"time"
)

var logFlag uint32 = LogFlagAll
var version string = ""
var showTime bool = true
const (
	LogFlagInfo = 1 << 0
	LogFlagDebug = 1 << 1
	LogFlagWarning = 1 << 2
	LogFlagError = 1 << 3
	LogFlagAll = LogFlagInfo | LogFlagWarning | LogFlagDebug | LogFlagError
)


const (
	LogColorOrange = uint8(iota+91)
	LogColorGreen
	LogColorYellow
	LogColorPurple
	LogColorMagenta
	LogColorBlue
	LogColorWhite
)


// 设置日志输出掩码
func SetLogFlag(_flag uint32) {
	logFlag = _flag
}

// 设置版本号
func SetVersion(_version string) {
	version = _version
}

// 设置时间
func SetShowTime(_show bool) {
	showTime = _show
}

// 输出Info日志
func Info(_fmt string, _args ...interface{}) {
	if logFlag & LogFlagInfo == 0 {
		return
	}
	print(LogFlagInfo, _fmt, _args...)
}


// 输出Debug日志
func Debug(_fmt string, _args ...interface{}) {
	if logFlag & LogFlagDebug == 0 {
		return
	}
	print(LogFlagDebug, _fmt, _args...)
}


// 输出Warning日志
func Warning(_fmt string, _args ...interface{}) {
	if logFlag & LogFlagWarning == 0 {
		return
	}
	print(LogFlagWarning, _fmt, _args...)
}


// 输出Error日志
func Error(_fmt string, _args ...interface{}) {
	if logFlag & LogFlagError == 0 {
		return
	}
	print(LogFlagError, _fmt, _args...)
}


// 打印
func print(_type uint32, _fmt string, _args ...interface{}) {
	var timeStr = clTime.GetDateByFormat(uint32(time.Now().Unix()), "15:04:05")
	var content = _fmt
	if _args != nil && len(_args) > 0 {
		content = fmt.Sprintf(_fmt, _args...)
	}

	debugFile := ""
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fileItems := strings.Split(file,"/")
		debugFile = fmt.Sprintf("%v:%v", fileItems[len(fileItems)-1], line)
	}

	logContent := ""
	if showTime {
		logContent += timeStr
	}

	if version != "" {
		logContent += " " + version
	}
	logContent += " " + debugFile
	switch _type {
	case LogFlagInfo:
		logContent += "[Info]"
	case LogFlagWarning:
		logContent += "[Warn]"
	case LogFlagError:
		logContent += "[Err]"
	case LogFlagDebug:
		logContent += "[Debug]"
	}

	fmt.Println(logContent + " " + content)
}