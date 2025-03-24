package negoutils

import (
	"fmt"
	"os"
	"testing"
)

func TestDBase_Conn(t *testing.T) {
	myconf := MySQLConf{
		Host:       "100.90.150.28",
		User:       "xiaoju",
		Passwd:     "25119c08c893d783cf70fc57eb4a58f7",
		DbName:     "xiaoju",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	db.SetDebug(true)
	_, err := db.Conn()
	fmt.Println(err)
	if db != nil {
		defer db.Close()
	} else {
		panic("db is nil")
	}
	sql := "SELECT * FROM `trainstation_pickup`"
	ret, err := db.FetchRows(sql)
	fmt.Println(ret, err)
}

func TestDBase_InsertBatchData(t *testing.T) {
	myconf := MySQLConf{
		Host:       "100.90.150.28",
		User:       "xiaoju",
		Passwd:     "25119c08c893d783cf70fc57eb4a58f7",
		DbName:     "xiaoju",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	db.SetDebug(true)
	_, err := db.Conn()
	defer db.Close()
	fmt.Println(err)
	var data [][]interface{}
	fields := []string{"id", "uid"}
	for i := 0; i < 100; i++ {
		var tmp []interface{}
		tmp = append(tmp, i)
		tmp = append(tmp, i+1)
		data = append(data, tmp)
	}
	ret, b, e := db.InsertBatchData("xiaoju", fields, data, true)
	fmt.Println(ret, b, e)
}

func TestGetMySQLTableStruct(t *testing.T) {
	myconf := MySQLConf{
		Host:       "100.90.150.28",
		User:       "xiaoju",
		Passwd:     "25119c08c893d783cf70fc57eb4a58f7",
		DbName:     "xiaoju",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	_, err := db.Conn()
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, err := GetMySQLTableStruct(db, "trainstation_pickup")
	fmt.Println(err, ret)
}

func TestGetAllMySQLTables(t *testing.T) {
	myconf := MySQLConf{
		Host:       "100.90.150.28",
		User:       "xiaoju",
		Passwd:     "25119c08c893d783cf70fc57eb4a58f7",
		DbName:     "xiaoju",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	_, err := db.Conn()
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, err := GetAllMySQLTables(db)
	fmt.Println(err, ret)
}

func TestGetMySQLAllTablesStruct(t *testing.T) {
	myconf := MySQLConf{
		Host:       "100.90.150.28",
		User:       "xiaoju",
		Passwd:     "25119c08c893d783cf70fc57eb4a58f7",
		DbName:     "xiaoju",
		Charset:    "utf8",
		Timeout:    5,
		Port:       3306,
		AutoCommit: true,
	}
	db := NewDBase(myconf)
	_, err := db.Conn()
	fmt.Println(err)
	str, _ := GetMySQLAllTablesStruct(db)
	fmt.Println(str)
}

func TestFormatFieldNameToGolangType(t *testing.T) {
	fields := []string{
		"api",
		"1user",
		"user_name",
		"user1",
		"menuName",
		"user_Name",
		"_http_status_",
	}
	for _, f := range fields {
		fmt.Fprintf(os.Stdout, "fieldName:%s\tformatFieldName:%s\n", f, FormatFieldNameToGolangType(f))
	}
}
