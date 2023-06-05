package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"fmt"
	"encoding/json"
	"managementpackage/model"
)

// 序列化数据
// data 需要序列化的数据
func UserDataSerialization(data []model.User) (string, error) {
	var bytes []byte
	var err error
	bytes, err = json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 反序列化数据
// data 需要反序列化的数据
func UserDataDeserialization(data []byte) ([]model.User, error) {
	var newUser []model.User
	var err error = json.Unmarshal(data, &newUser)
	if err != nil {
		return make([]model.User, 0), err
	}
	return newUser, nil
}


// 写入用户信息文件
// path 文件路径
// data 写入内容
func UserFileWrite(path string, data string) error {
	var file *os.File
	var err error
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {

		var err error = file.Close()
		if err != nil {
			if errors.Is(err, os.ErrClosed) {
				// 文件已经被关闭
				fmt.Println("文件已经被关闭")
			} else if errors.Is(err, os.ErrNotExist) {
				// 文件已经被删除
				fmt.Println("文件已经被删除")
			} else if errors.Is(err, os.ErrInvalid) {
				// 文件描述符已经被其他进程或线程使用
				fmt.Println("文件描述符已经被其他进程或线程使用")
			} else {
				// 其他错误
				fmt.Println("其他错误")
			}
		}
	}(file)

	var writer *bufio.Writer = bufio.NewWriterSize(file, 1024*1024) // 设置缓冲区大小为1MB
	_, err = writer.Write([]byte(data))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// 读取用户信息文件
// path 文件路径
func UserFileReading(path string) ([]byte, error) {
	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			if errors.Is(err, os.ErrClosed) {
				fmt.Println("文件已经被关闭")
			} else if errors.Is(err, os.ErrNotExist) {
				fmt.Println("文件已经被删除")
			} else if errors.Is(err, os.ErrInvalid) {
				fmt.Println("文件描述符已经被其他进程或线程使用")
			} else {
				fmt.Println("其他错误")
			}
		}
	}(file)
	var content []byte
	var buf [4096]byte
	var reader *bufio.Reader = bufio.NewReader(file)
	for {
		var n int
		var err error
		n, err = reader.Read(buf[:])
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		content = append(content, buf[:n]...)
	}
	return content, nil
}