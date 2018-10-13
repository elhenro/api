package main

import (
    "encoding/json"
	"log"
	"os"
    "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
)

func Read(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("testdata/hello")
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
	f, err := os.OpenFile("testdata/hello", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}


    json.NewEncoder(w).Encode(content)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/read", Read).Methods("GET")
	router.HandleFunc("/write/{p}", Write).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}