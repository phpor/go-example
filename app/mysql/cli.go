package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// "user:password@tcp(127.0.0.1:3306)/hello"
func main() {
	username := flag.String("username", "root", "username")
	password := flag.String("password", "123456", "password")
	host := flag.String("host", "127.0.0.1", "host")
	port := flag.String("port", "3306", "port")
	dbname := flag.String("dbname", "mysql", "dbname")
	sqlStr := flag.String("sql", "select 'ok' as Res", "sql")
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *username, *password, *host, *port, *dbname))
	if err != nil {
		println(err)
		return
	}
	r := db.QueryRowContext(context.Background(), *sqlStr)
	res := ""
	if err := r.Scan(&res); err != nil {
		println(err.Error())
		return
	}
	println(res)
}
