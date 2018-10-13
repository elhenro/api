package main

import (
	"bytes"
	"unicode"
	"fmt"
	"strings"
	"time"
    "encoding/json"
	"log"
	"os"
    "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"strconv"
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
	t1 := getLastLineOfFile(timeStartFile)
	t2 := getLastLineOfFile(timeEndFile)

	times := []string{t1, ",", t2}
	var str bytes.Buffer
	for _, l := range times {
		str.WriteString(l)
	}

	json.NewEncoder(w).Encode(str.String())
}

func setTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var timeRequest string
	_ = json.NewDecoder(r.Body).Decode(&timeRequest)
	t1 := params["t1"]	
	t2 := params["t2"]

	var response string
	if (IsLetter(t1) || IsLetter(t2)) {
		response = "error, not in time hour format"
		err := response
		panic(err)
	} else {	
		writeToFile(timeStartFile, t1)
		writeToFile(timeEndFile, t2)

		res := getTimeframeHours(t1, t2)
    	json.NewEncoder(w).Encode(res)
	}
}

func writeToFile(file string, content string){
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
	// new line
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
}

func getTimeframeHours(t1 string, t2 string) float64{
	// 10:30, 17:30
	// returns timeframe in hours  e.g.: 2.8

	hhmm1 := strings.Split(t1, ":")
	hhmm2 := strings.Split(t2, ":")
	hh1, err := strconv.Atoi(hhmm1[0])
	mm1, err := strconv.Atoi(hhmm1[1])
	hh2, err := strconv.Atoi(hhmm2[0])
	mm2, err := strconv.Atoi(hhmm2[1])
	
	if err != nil {
        fmt.Println(err)
    }

	t := time.Now()
	berlinTime, err := time.LoadLocation("Europe/Berlin")
	s1 := time.Date(t.Year(), t.Month(), t.Day(), hh1, mm1, 0, 0, berlinTime)
	s2 := time.Date(t.Year(), t.Month(), t.Day(), hh2, mm2, 0, 0, berlinTime)

	//d := s2.Since(s1)
	d := s2.Sub(s1)
	return d.Hours()
}

func getTimePercentage(w http.ResponseWriter, r *http.Request){
	t1 := getLastLineOfFile(timeStartFile)
	t2 := getLastLineOfFile(timeEndFile)
	now := time.Now()
	
	//t := "22:30"
	t := now.Format( "15:04" )

	//timeNow := time.Date
	timeOver := getTimeframeHours(t1, t) 
	timeFrame := getTimeframeHours(t1, t2)

	p := FloatToString(timeOver / timeFrame)
	//p := timeOver
	//p := timeFrame
	
	fmt.Printf("start: %s\n", t1)
	fmt.Printf("end: %s\n",t2)
	fmt.Printf("now: %s\n",t)
	fmt.Printf("percentage: %s\n",p)

	json.NewEncoder(w).Encode(p)
}

func getLastLineOfFile(fname string) string {
	file, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    buf := make([]byte, 62)
    stat, err := os.Stat(fname)
    start := stat.Size() - 62
    _, err = file.ReadAt(buf, start)
	
	//n := bytes.Index(buf, []byte{0})
	lines := string(buf)

	//fmt.Printf("%s\n", lines)
	l := strings.Split(lines, "\n")
	ll := l[len(l)-2]
	return ll
}

func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func FloatToString(input_num float64) string {
    return strconv.FormatFloat(input_num, 'f', 6, 64)
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