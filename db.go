package main

import (
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	L "./lib"
)

type Text struct {
	ID   int    `json:"id"`
	Subject string `json:"name"`
	Creation string `json:"name`
}

const (
	mysqlUser = "tester"
	mysqlPass = "test"
)

func mysqlGet(database string, row string, item string, limit int) []Text {
	
	q := make([]Text, 0, 2)

	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	//defer db.Close()	
	results, err := db.Query(L.Join("SELECT ", item, " FROM ", row ," limit ", strconv.Itoa(limit)))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var r Text
		err = results.Scan(&r.ID, &r.Subject, &r.Creation)
		if err != nil {
			panic(err.Error())
		}
		q = append(q, Text{r.ID, r.Subject, r.Creation})
	}
	db.Close();
	return q
}
