package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

func main() {

	db := NewSsoDbUser()
	db.drop_table()
	db.create_table()
	db.insert("phpor2", "ssologin2")
	db.insert("phpor3", "ssologin3")
	db.query()
	db.update()
	db.query2()
}

type SsoDbUser struct {
	db *xorm.Engine
}

func NewSsoDbUser() *SsoDbUser {
	orm, err := xorm.NewEngine("mysql", "root:sina@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	orm.ShowSQL = true
	return &SsoDbUser{db: orm}
}

func (this *SsoDbUser) query() {
	orm := this.db
	result, err := orm.Query("select * from user")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
func (this *SsoDbUser) query2() {
	orm := this.db
	user := &User{Name:"phpor2"}
	exists, err := orm.Get(user)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	if exists {
		fmt.Println(user)
	} else {
		fmt.Println("Not found")
	}
}

func (this *SsoDbUser) update() {
	orm := this.db
	user := &User{Name:"phpor2", Password:"password2"}
	effected, err := orm.Cols("password").Update(user)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	if effected > 0 {
		fmt.Println("update succ")
	}
}

func (this *SsoDbUser) insert(name, password string) {
	orm := this.db
	user := &User{Name:name, Password:password}

	effected, err := orm.Insert(user)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(effected)
}
func (this *SsoDbUser) show_tables() {
	orm := this.db
	tables, err := orm.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, table := range tables {
		fmt.Println(table.Name)
	}
}
func (this *SsoDbUser) drop_table() {
	this.db.Exec("drop table user")
}
func (this *SsoDbUser) create_table() {
	orm := this.db
	sql := `create table user (
		uid int(10) unsigned NOT NULL AUTO_INCREMENT,
		name varchar(64) NOT NULL,
		password varchar(20) NOT NULL,
		primary key (uid)
	)`
	orm.Exec(sql)
}

type User struct {
	Uid      int64  `xorm:"int(10) unsigned not null pk autoincr uid"`
	Name     string `xorm:"char(16) not null name"`
	Password string `xorm:"char(32) not null password"`
}

func (this *User) TableName() string {
	// 对于分库分表的情况，这里就可以根据Uid来计算表名字了
	return "user"
}
