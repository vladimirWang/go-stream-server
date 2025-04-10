package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/video_server?charset=utf8"
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
}
