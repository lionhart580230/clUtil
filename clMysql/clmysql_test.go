package clMysql

import (
	"fmt"
	"testing"
)

func TestAddMaster(t *testing.T) {

	db := NewDBSimple("127.0.0.1", "root", "root", "supermaket")
	if db == nil {
		fmt.Printf("connect to mysql failed\n")
		return
	}

	db.NewBuilder().Table("table1").LeftJoin("table2", "table1.id = table2.id").Where("table1.id > 100").Count()
	fmt.Printf("SQL: %v\n", db.GetLastSql())
}
