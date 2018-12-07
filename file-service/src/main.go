package main

//"C:/Users/omerd/workspace/yobros/file-service/src/DBhandlers"

import (
	"log"
	"net/http"

	ufh "../src/pkg/dbh"
)

var config Config

func main() {
	config, _ = LoadConfig("server-config.json")

	http.HandleFunc("/files", fileReaderHandler)
	http.ListenAndServe(":80", nil)
}

func fileReaderHandler(w http.ResponseWriter, req *http.Request) {

	dir := config.ServerDirectory

	log.Println("dir: ", dir)

	err := ufh.CreateDirectory(dir)
	log.Println(err)

	filename, err := ufh.WriteBytesToFile(req)
	log.Println(err)

	log.Println(filename)

	ufh.MoveFile(filename, dir)

}
