package main

import (
	"fmt"
	//"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	L "./lib"
	"golang.org/x/crypto/bcrypt"
)

type Text struct {
	ID   int    `json:"id"`
	Subject string `json:"name"`
	Creation string `json:"name`
}

type ChatItem struct {
	id int `json:"id"`
	message string `json:"message"`
	author string `json:"author"`
	creation string `json:"creation"`
}

type User struct {
	id int `json:"id"`
	name string `json:id`
	pass string `json:pass`
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

func mysqlcheckUserWithPass(name string, pass string) bool{
	var r User
	q := make([]User, 0, 2)

	c := L.Join(mysqlUser, ":", mysqlPass, "@/", "api")
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query(L.Join("SELECT id,name,pass FROM users where name like '", name, "';"))
	if err != nil {
		panic(err.Error())
		return false
	}
	for results.Next() {
		err = results.Scan(&r.id, &r.name,&r.pass)
		if err != nil {
			panic(err.Error())
			return false
		}
		q = append(q, User{r.id, r.name, r.pass})
	}
	fmt.Println("user ", name)

	fmt.Println("comparing ", pass, " and ", r.pass)

	if name != r.name{
		fmt.Println("user ", name, " not found")
		return false
	}
	if CheckPasswordHash(pass, r.pass) {
		fmt.Println("true")
		return true
	} else {
		return false
	}
}

func mysqlAddUser(name string, pass string)bool{
	c := L.Join(mysqlUser, ":", mysqlPass, "@/", "api")
	db, err := sql.Open("mysql", c)
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()

	stmtIns, err := db.Prepare(L.Join("INSERT INTO users ( id, name, pass ) VALUES( ?, ?, ? )"))
	if err != nil {
		panic(err.Error())
		return false
	}
	defer stmtIns.Close() 
	hashedPass, err := HashPassword(pass)
	if err != nil {
		panic(err.Error())
		return false
	}
	_, err = stmtIns.Exec( mysqlGetNewHighestID("api","users"), name, hashedPass)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func mysqlWriteChatText(database string, row string, id int, content string, author string) {
	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare(L.Join("INSERT INTO ", row, " ( id, message, author, creation ) VALUES( ?, ?, ?, NOW() )"))
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() 

	_, err = stmtIns.Exec( id, content, author)
	if err != nil {
		panic(err.Error())
	}

}

func mysqlGetChatItems(database string, row string, limit int) []ChatItem {
	q := make([]ChatItem, 0, 2)

	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query(L.Join("SELECT id, message, author, creation FROM ", row, ";"))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var r ChatItem
		err = results.Scan(&r.id, &r.message, &r.author, &r.creation)
		if err != nil {
			panic(err.Error())
		}
		q = append(q, ChatItem{r.id, r.message, r.author, r.creation})
	}
	return q
}

func mysqlGetChatItemByID(database string, row string, itemID string, limit int) []ChatItem {
	q := make([]ChatItem, 0, 2)

	c := L.Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query(L.Join("SELECT id, message, author, creation FROM ", row, " where id= ", itemID,";"))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var r ChatItem
		err = results.Scan(&r.id, &r.message, &r.author, &r.creation)
		if err != nil {
			panic(err.Error())
		}
		q = append(q, ChatItem{r.id, r.message, r.author, r.creation})
	}
	return q
}