package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/files", fileReaderHandler)
	http.ListenAndServe(":80", nil)
}

func fileReaderHandler(w http.ResponseWriter, r *http.Request) {

	fileRequest := r.Body
	fmt.Println("fileRequst: ", fileRequest)

	fileHeader := r.Header
	fmt.Println("fileHeader: ", fileHeader)

	config, _ := LoadConfig("server-config.json")
	dir := config.ServerDirectory

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	createDirectory(dir)
	createEmptyFile(u1.String())

	reader := fileRequest
	p := make([]byte, 1)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			writeBytesToFile(u1.String(), p[n-1:])
			fmt.Println(string(p[n-1:]))
			break
		}
		if err != nil {
			log.Println(err)
		}
		writeBytesToFile(u1.String(), p[:n])
		fmt.Println(string(p[:n]))
	}
	moveFile(u1.String(), dir)
}

func createDirectory(directoryPath string) {
	pathErr := os.MkdirAll(directoryPath, 0777)

	if pathErr != nil {
		log.Println("An error occurred while creating the directory! ERROR:005")
		log.Println(pathErr)
		return
	}
}

func moveFile(filename string, path string) {
	err := os.Rename(filename, path+filename)
	if err != nil {
		log.Println(err)
		log.Println("An error occurred while renaming the file directory! ERROR:004")
		return
	}
	fmt.Println("Files moved succesffuly")
}

func createEmptyFile(filename string) {
	newFile, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		log.Println("An error occurred while creating the file! ERROR:003")
		return
	}
	newFile.Close()
	fmt.Println("File created succesffuly")
}

func writeBytesToFile(filename string, bs []byte) {
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_WRONLY,
		777,
	)
	if err != nil {
		log.Println("An error occurred while re-writing the file! ERROR:006/0")
		return
	}

	bytesWritten, err := file.Write(bs)
	if err != nil {
		log.Println("An error occurred while re-writing the file! ERROR:006/1")
		return
	}

	defer file.Close()

	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
