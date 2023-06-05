package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const filePath = "a.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}
	connection, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer connection.Close()

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	if _, err := fmt.Fprintf(connection, "%s %d\n", fileName, fileSize); err != nil {
		log.Fatalf("Error sending file info: %v", err)
	}
	fmt.Printf("开始上传文件 %s...\n", fileName)
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return
			}
			break
		}
		if _, err := fmt.Fprintln(connection, bytesRead); err != nil {
			log.Fatal(err)
			return
		}
		if _, err := connection.Write(buffer[:bytesRead]); err != nil {
			log.Fatalf("Error sending file data: %v", err)
		}
	}

	responseBuffer := make([]byte, 1024)
	bytesReceived, _ := connection.Read(responseBuffer)
	fmt.Println(string(responseBuffer[:bytesReceived]))

}
