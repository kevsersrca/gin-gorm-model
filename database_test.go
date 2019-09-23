package database

import(
	"fmt"
	"testing"
	"time"
)


func TestOrmConn_AddLog(t *testing.T) {
	orm := NewOrm()
	if orm.orm.Error != nil{
		t.Error(orm.orm.Error)
	}
	err := orm.AddLog(Log{
		Server:     "175.5.5.5",
		Description: "Create Edildi",
		Created: time.Time{},
		Updated: time.Time{},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Skip()
}

func TestOrmConn_GetLogs(t *testing.T) {
	orm := NewOrm()
	if orm.orm.Error != nil{
		t.Error(orm.orm.Error)
	}
	find, err := orm.GetLogs()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(find)
	t.Skip()
}

func TestOrmConn_RemoveLogs(t *testing.T) {
	orm := NewOrm()
	if orm.orm.Error != nil{
		t.Error(orm.orm.Error)
	}
	err := orm.RemoveLogs()
	if err != nil {
		t.Error(err)
	}
	t.Skip()
}
