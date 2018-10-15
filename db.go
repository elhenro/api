package main

import (
	//"strconv"
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

func mysqlGet(database string, row string, item string/*, limit int*/) []Text {
	q := make([]Text, 0, 2)

	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query(L.Join("SELECT ", item, " FROM ", row /*," limit ", strconv.Itoa(limit)*/))
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
	//db.Close();
	return q
}

func mysqlGetNewHighestID(database string, row string) int {
	
	//q := make([]Text, 0, 2)
	var q int
	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query( L.Join("SELECT id FROM ", row, " ORDER BY id DESC LIMIT 0, 1;"))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var r Text
		err = results.Scan(&r.ID)
		if err != nil {
			panic(err.Error())
		}
		q = r.ID + 1
	}
	return q
}

func mysqlWriteText(database string, row string, id int, content string) {
	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare(L.Join("INSERT INTO ", row, " ( id, subject, creation ) VALUES( ?, ?, NOW() )"))
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() 

	_, err = stmtIns.Exec( id, content)
	if err != nil {
		panic(err.Error())
	}

}