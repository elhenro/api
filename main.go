package main

import (
	"os"
	"bytes"
    "encoding/json"
	"log"
    "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	L "./lib"
)

const (
	timeFormat = "2006-01-02 15:04 MST"
	timeStartFile = "testdata/timeStart"
	timeEndFile = "testdata/timeEnd"
	helloFile = "testdata/hello"
)

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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time", getTime).Methods("GET")
	router.HandleFunc("/time/set/{t1}/{t2}", setTime).Methods("POST")
	router.HandleFunc("/time/percent", getTimePercentage).Methods("GET")

	router.HandleFunc("/read", Read).Methods("GET")
	router.HandleFunc("/write/{p}", Write).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}