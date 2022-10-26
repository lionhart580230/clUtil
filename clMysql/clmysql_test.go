package clMysql

import (
	"fmt"
	"github.com/xiaolan580230/clUtil/clLog"
	"testing"
)

func TestAddMaster(t *testing.T) {

	db := NewDBSimple("127.0.0.1", "root", "root", "miner_new")
	if db == nil {
		fmt.Printf("connect to mysql failed\n")
		return
	}

	tableList, err := db.GetTables("")
	if err != nil {
		clLog.Error("获取数据库表格错误: %v", err)
		return
	}
	for _, table := range tableList {
		clLog.Debug("表格: %v", table)
		columnList, err := db.GetTableColumns(table)
		if err != nil {
			clLog.Error("获取表格字段数据错误: %v", err)
			continue
		}
		clLog.Debug("表格: %v 字段数据: %+v", table, columnList)
	}

	clLog.Debug("生成完毕!")
}
