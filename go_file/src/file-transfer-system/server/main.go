package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:8888...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		} else {
			go handleRequest(conn)
		}
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	fileName, fileSize, err := readFileInfo(conn)
	if err != nil {
		log.Printf("Error reading file info: %v", err)
		return
	}
	fmt.Println(fileName)
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()
	err = receiveFile(conn, file, fileSize)
	if err != nil {
		log.Printf("Error receiving file: %v", err)
		return
	}
	conn.Write([]byte("上传完成"))
}

func readFileInfo(conn net.Conn) (string, int, error) {
	var fileName string
	var fileSize int
	if _, err := fmt.Fscanln(conn, &fileName, &fileSize); err != nil {
		return "", 0, err
	}
	return fileName, fileSize, nil
}

func receiveFile(conn net.Conn, file *os.File, fileSize int) error {
	buf := make([]byte, 1024)
	var serverFileSize int
	for serverFileSize < fileSize {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		serverFileSize += n
		data := buf[:n]
		_, err = file.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
