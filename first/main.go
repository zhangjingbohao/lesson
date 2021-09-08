package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

func main() {
	if err := Job(); err != nil {
		log.Print("err:", err)
	}
}

func Job() error {
	return errors.WithMessage(QueryMysql(), "Job.QueryMysql")
}

// dao层数据库操作
func QueryMysql() error {
	db, err := sql.Open("mysql", "user:pwd@tcp(ip:port)/DB")
	if err != nil {
		panic(err)
	}
	_, err = db.Query("select * from table_name")
	return errors.Wrap(err, "QueryMysql.err")
}
