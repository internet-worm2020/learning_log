package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var path string ="/home/gao/test/dome/learning_log/go_file/src/file-transfer-system/client/a.txt"
func main() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	fmt.Println(fileInfo.Name(),fileInfo.Size())
    if err != nil {
        fmt.Println(err)
        return
    }
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()
    if _, err := fmt.Fprintln(conn, fileName,fileSize); err != nil {
        fmt.Println(err)
        return
    }
	fmt.Printf("开始上传文件 %s...\n", fileName)
	buffer := make([]byte, 1024)
	for {
        n, err := file.Read(buffer)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
		if _, err := fmt.Fprintln(conn, n); err != nil {
			fmt.Println(err)
			return
		}
		if _, err := conn.Write(buffer[:n]); err != nil {
            fmt.Println(err)
            return
        }
    }
	var xc []byte=make([]byte, 1024)
	c,_:=conn.Read(xc)
	fmt.Println(string(xc[:c]))

}
