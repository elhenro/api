package main

import (
	"os"
	"strconv"
	"bytes"
    "encoding/json"
	"log"
    "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	L "./lib"
)

func getTextFromDB(w http.ResponseWriter, r *http.Request) {

	resultFromDb := mysqlGet("api", "texts", "*", 0)
	content:= resultFromDb

	result := content
	json.NewEncoder(w).Encode(result)
}

func getNewIDfromTextDB(w http.ResponseWriter, r *http.Request) {
	res := mysqlGetNewHighestID("api", "texts")
	json.NewEncoder(w).Encode(res)
}

func writeTextToDB(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var content string
	_ = json.NewDecoder(r.Body).Decode(&content)
	content = params["p"]	

	mysqlWriteText("api", "texts", mysqlGetNewHighestID("api", "texts"), string(content))

	result := mysqlGet("api", "texts", "*", 0)
	json.NewEncoder(w).Encode(result)
}

func Read(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile( helloFile )
	if err != nil {
		log.Fatal(err)
	}

	result := string(content)
	json.NewEncoder(w).Encode(result)
}

func Write(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var content string
	_ = json.NewDecoder(r.Body).Decode(&content)
	content = params["p"]	
	f, err := os.OpenFile( helloFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
    json.NewEncoder(w).Encode(content)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	t1 := L.GetLastLineOfFile(timeStartFile)
	t2 := L.GetLastLineOfFile(timeEndFile)

	times := []string{t1, ",", t2}
	var str bytes.Buffer
	for _, l := range times {
		str.WriteString(l)
	}

	json.NewEncoder(w).Encode(str.String())
}
func getTimePercentage(w http.ResponseWriter, r *http.Request){
	p := L.GetTimePercentage(timeStartFile, timeEndFile)
	json.NewEncoder(w).Encode(p)
}

func setTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var timeRequest string
	_ = json.NewDecoder(r.Body).Decode(&timeRequest)
	t1 := params["t1"]	
	t2 := params["t2"]

	var response string
	if (L.IsLetter(t1) || L.IsLetter(t2)) {
		response = "error, not in time hour format"
		err := response
		panic(err)
	} else {	
		L.WriteToFile(timeStartFile, t1)
		L.WriteToFile(timeEndFile, t2)

		res := L.GetTimeframeHours(t1, t2)
    	json.NewEncoder(w).Encode(res)
	}
}

func getHelp(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile( helpFile )
	if err != nil {
		log.Fatal(err)
	}

	result := string(content)
	json.NewEncoder(w).Encode(result)
}

func checkUserWithPass(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	//var timeR string
	//_ = json.NewDecoder(r.Body).Decode(&timeRequest)
	name := params["name"]	
	pass := params["pass"]

	res := mysqlcheckUserWithPass(name, pass)
	json.NewEncoder(w).Encode(res)
}	

func addUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	//var timeR string
	//_ = json.NewDecoder(r.Body).Decode(&timeRequest)
	name := params["name"]	
	pass := params["pass"]

	res := mysqlAddUser(name, pass)
	json.NewEncoder(w).Encode(res)
}

func ChatRead(w http.ResponseWriter, r *http.Request) {
	var res string
	for _, p := range mysqlGetChatItems("api", "chat", 100) {
			res = L.Join(res, strconv.Itoa(p.id), p.message, p.author, p.creation)
		}
	json.NewEncoder(w).Encode(res)
}

func ChatWrite(w http.ResponseWriter, r *http.Request) {
	var res string
	params := mux.Vars(r)
    var contentIn string
	_ = json.NewDecoder(r.Body).Decode(&contentIn)
	content := params["c"]	
	author := params["a"]
	pass := params["p"]

	if mysqlcheckUserWithPass(author, pass){
		newID := mysqlGetNewHighestID("api", "chat")
		mysqlWriteChatText("api", "chat", newID, content, author)

		for _, p := range mysqlGetChatItemByID("api", "chat", strconv.Itoa(newID), 0) {
			res = L.Join(res, strconv.Itoa(p.id), p.message, p.author, p.creation)
		}
	}
	
    json.NewEncoder(w).Encode(res)
}