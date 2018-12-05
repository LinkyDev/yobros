package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/files", fileReaderHandler)
	http.ListenAndServe(":80", nil)
}

func fileReaderHandler(w http.ResponseWriter, r *http.Request) {

	var (
		s          string
		fileHeader string
		byteSlice  []byte
	)

	count := 0

	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "An error occurred while uploading the file! ERROR:001\n", http.StatusBadGateway)
			return
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "An error occurred while reading the file! ERROR:002\n", http.StatusInternalServerError)
			return
		}

		fileHeader = h.Filename
		log.Println(fileHeader)

		byteSlice = bs
		s = string(bs)

		count++

	}

	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/files" method="post" enctype="multipart/form-data">
		upload a file<br>
		<input type="file" name="usrfile"><br>
		<input type="submit">
		</form><br>
		<br>
		<br>
		<h1>%v</h1>`, s)

	config, _ := LoadConfig("server-config.json")
	dir := config.Database.ServerDirectory

	createDirectory(dir)
	for fileHeader != "" && count < 2 {
		createEmptyFile(fileHeader)
		writeBytesToFile(fileHeader, byteSlice)
		moveFile(fileHeader, dir)

		count++
	}

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
	err := os.Rename(filename, path)
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
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
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
