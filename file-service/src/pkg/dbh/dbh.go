package dbh

import (
	"io"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

//CreateDirectory creates a new directory path, if directory already exists does nothing - returns error if not equals to nil
func CreateDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, 0777)

	if err != nil {
		return err
	}
	return nil
}

//MoveFile moves file of the "filename string" to the new path "path string" - returns error if not equals to nil
func MoveFile(filename string, path string) error {
	err := os.Rename(filename, path+filename)
	if err != nil {
		return err
	}
	return nil
}

//WriteBytesToFile writes the information about a file to a newly created file and returns error if has any - returns error if not equals to nil
func WriteBytesToFile(fileRequest *http.Request) (name string, err error) {

	filename := uuid.Must(uuid.NewV4()).String()
	log.Println("filename in: ", filename)

	newFile, err := os.Create(filename)

	if err != nil {
		return "", err
	}

	openedFile, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_WRONLY,
		777,
	)

	if err != nil {
		return "", err
	}

	defer openedFile.Close()
	defer newFile.Close()

	byteChunk := make([]byte, 10)
	//buffWriter := bufio.NewWriter(newFile)

	for {
		byteLen, err := fileRequest.Body.Read(byteChunk)

		if err == io.EOF {
			return filename, nil
		}

		bytesWritten, err := newFile.Write(byteChunk[:byteLen])
		defer log.Printf("%d", bytesWritten)
		log.Printf("%s", byteChunk[:byteLen])
	}

}
