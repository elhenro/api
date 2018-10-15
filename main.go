package main

import (
	"log"
	"strconv"
    "github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
	L "./lib"
)

const (
	timeFormat = "2006-01-02 15:04 MST"
	timeStartFile = "testdata/timeStart"
	timeEndFile = "testdata/timeEnd"
	helloFile = "testdata/hello"
	helpFile = "testdata/help"
	port = 8000
)

func main() {
	router := mux.NewRouter()

	handler := cors.Default().Handler(router)

	router.HandleFunc("/", getHelp).Methods("GET")
	router.HandleFunc("/time", getTime).Methods("GET")
	router.HandleFunc("/time/set/{t1}/{t2}", setTime).Methods("POST")
	router.HandleFunc("/time/percent", getTimePercentage).Methods("GET")

	router.HandleFunc("/read", Read).Methods("GET")
	router.HandleFunc("/write/{p}", Write).Methods("POST")

	router.HandleFunc("/text", getTextFromDB).Methods("GET")
	router.HandleFunc("/text/{p}", writeTextToDB).Methods("POST")

	router.HandleFunc("/getnewid", getNewIDfromTextDB).Methods("GET")
    log.Fatal(http.ListenAndServe(L.Join(":",strconv.Itoa(port)), handler/*handlers.CORS(corsOk, headersOk, methodsOk)(router)*/))
}