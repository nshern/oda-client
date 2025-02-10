package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
)

func establishConnection() *ftp.ServerConn {
	ftp_adr := "oda.ft.dk:21"

	connection, err := ftp.Dial(ftp_adr)
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func getFolderPaths(connection *ftp.ServerConn) []string {

	var folderPaths []string

	connection.ChangeDir("ODAXML/Referat/samling/")

	currentDir, err := connection.CurrentDir()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	fileList, err := connection.List(currentDir)
	if err != nil {
		log.Fatalf("Error listing files: %v", err)
	}

	for _, k := range fileList {
		path := currentDir + "/" + k.Name
		folderPaths = append(folderPaths, path)
	}

	return folderPaths
}

func getFilePaths(folderPaths []string, connection *ftp.ServerConn) []string {

	var filePaths []string

	for _, folderPath := range folderPaths {
		connection.ChangeDir(folderPath)
		currentDir, _ := connection.CurrentDir()
		files, _ := connection.List(currentDir)
		for _, file := range files {
            filePath := (currentDir + "/" + file.Name)
            filePaths = append(filePaths,filePath)
		}
	}

    return filePaths
}

func main() {
	connection := establishConnection()
	folderPaths := getFolderPaths(connection)
    filePaths := getFilePaths(folderPaths, connection)

	for _, v := range filePaths {
		fmt.Println(v)
	}

}
