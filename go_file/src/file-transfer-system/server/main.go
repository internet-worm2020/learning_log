package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer l.Close()
	fmt.Println("Listening on localhost:8888...")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
		} else {
			go handleRequest(conn)
		}
	}
}

func handleRequest(conn net.Conn) {
	// defer conn.Close()
	buf := make([]byte, 1024)
	var fileName string
	var fileSize int
	if _, err := fmt.Fscanln(conn, &fileName, &fileSize); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var serverSize int
	for {
		if fileSize==serverSize{
			break
		}
		var nn int
		if _, err := fmt.Fscanln(conn, &nn); err != nil {
			fmt.Println(err)
			return
		}
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		serverSize+=n
		if n == nn {
			data := buf[:n]
			_, err = io.WriteString(file, string(data))
			if err != nil {
				fmt.Println(err)
			}
		}else{
			conn.Write([]byte("数据丢失！"))
			conn.Close()
		}
	}
	conn.Write([]byte("上传完成"))
}
