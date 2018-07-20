package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Fprintf(w, "Received\n")
		decoder := json.NewDecoder(req.Body)
		var t map[string]string
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		logMessage(t["message"])
	case "GET":
		b, err := ioutil.ReadFile(readConfig()["logFileLocation"])
		check(err)
		str := string(b)
		fmt.Fprintf(w, str)
	default:
		fmt.Fprintf(w, "Method Not Supported\n")
	}
}
