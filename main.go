package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", httpHandler)
	go http.ListenAndServe(readConfig()["httpPort"], nil)

	telnetPort := readConfig()["telnetPort"]

	logMessage(fmt.Sprintf("Starting Telnet Server on Port %v", telnetPort))
	logMessage(fmt.Sprintf("Http Server on Port %v", readConfig()["httpPort"]))

	ln, err := net.Listen("tcp", telnetPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chat := CreateChat()
	for {
		conn, err := ln.Accept()
		check(err)
		chat.Connect(conn)
	}

}

func readConfig() map[string]string {
	absPath, _ := filepath.Abs("go/src/github.com/timdang/go_telnet/config.json")
	f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	check(err)
	var i map[string]string
	json.Unmarshal(b, &i)
	return i
}
