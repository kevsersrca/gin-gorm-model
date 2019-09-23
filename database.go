package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type OrmConn struct {
	orm *gorm.DB
}

type Log struct {
	Id      	int       		`gorm:"primary_key:true;auto_increment"`
	Server  	string   		`gorm:"size(255)"`
	Description string 			`gorm:"size(255)"`
	Created 	time.Time 		`gorm:"auto_now_add;type(datetime)"`
	Updated 	time.Time 		`gorm:"auto_now;type(datetime)"`
}




func NewOrm() (conn OrmConn) {
	orm, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	conn.orm = orm
	conn.orm.AutoMigrate(&Log{})
	return
}

func (c OrmConn) AddLog(log Log) error {
	c.orm.Create(&log)
	//return error
	return c.orm.Error
}
func (c OrmConn) GetLogs() (logs []Log, err error) {
	c.orm.Find(&logs)
	//return error
	return logs, c.orm.Error
}

func (c OrmConn) GetLog(key string) (find Log, err error) {
	find = Log{
		Server:     key,
	}
	c.orm.Find(&find)
	//return error
	return find, c.orm.Error
}

func (c OrmConn) RemoveLog(key string) error {
	find, err := c.GetLog(key)
	if err != nil {
		return err
	}
	c.orm.Delete(&find)
	return c.orm.Error
}

func (c OrmConn) RemoveLogs() error {
	find, err := c.GetLogs()
	if err != nil {
		return err
	}
	c.orm.Delete(&find)
	return c.orm.Error
}

