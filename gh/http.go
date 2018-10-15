package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

func sendPost(url string) error{
	 //url := "http://restapi3.apiary.io/notes"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		//panic(err)
		return err
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func getGet(url string)string{
  r := ""
  response, err := http.Get(url)
  if err != nil {
      fmt.Printf("%s", err)
      //return err
      //os.Exit(1)
  } else {
      defer response.Body.Close()
      contents, err := ioutil.ReadAll(response.Body)
      if err != nil {
          //return err  
          fmt.Printf("%s", err)
          //os.Exit(1)
      }
      r = string(contents)
      //fmt.Printf("%s\n", string(contents))
  }
  return r
}