package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const filePath string = "user_file.json"

type User struct {
	Id     int
	Name   string
	Age    int
	Career string
}

func NewUser(id int, name string, age int, career string) *User {
	return &User{
		Id:     id,
		Name:   name,
		Age:    age,
		Career: career,
	}
}
func (u *User) ModifyUser(modifyUser *User) {
	if modifyUser.Name != "" {
		u.Name = modifyUser.Name
	}
	if modifyUser.Age != 0 {
		u.Age = modifyUser.Age
	}
	if modifyUser.Career != "" {
		u.Career = modifyUser.Career
	}
}

func main() {
	fmt.Println("客户关系管理系统")
	var userList []User
	var data []byte
	var err error
	data, err = userFileReading(filePath)
	if err != nil {
		userList = []User{
			*NewUser(1, "yuan", 23, "CEO"),
		}
	} else {
		userList, _ = userDataDeserialization(data)
	}

	var scanner *bufio.Scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))

	for {
		fmt.Println(`
1 查看用户
2 添加用户
3 修改用户
4 删除用户
5 退出
    `)
		scanner.Scan()
		var option int
		option, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		switch option {
		case 1:
			listUser(&userList)
		case 2:
			addUser(&userList, scanner)
		case 3:
			modifyUser(&userList, scanner)
		case 4:
			deleteUser(&userList, scanner)
		case 5:
			os.Exit(0)
		default:
			println("未知选项")
		}
	}

}

// 写入用户信息文件
// path 文件路径
// data 写入内容
func userFileWrite(path string, data string) error {
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
func userFileReading(path string) ([]byte, error) {
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

// 序列化数据
// data 需要序列化的数据
func userDataSerialization(data []User) (string, error) {
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
func userDataDeserialization(data []byte) ([]User, error) {
	var newUser []User
	var err error = json.Unmarshal(data, &newUser)
	if err != nil {
		return make([]User, 0), err
	}
	return newUser, nil
}

func listUser(userList *[]User) {
	fmt.Println("查看用户")
	if len(*userList) == 0 {
		fmt.Println("没有用户信息")
	}
	fmt.Printf("|%-4s | %-10s | %-4s | %-10s |\n", "id", "name", "age", "career")

	for _, value := range *userList {
		fmt.Println("---------------------------------------")
		fmt.Printf("|%-4d | %-10s | %-4d | %-10s |\n", value.Id, value.Name, value.Age, value.Career)
	}
}

func addUser(userList *[]User, scanner *bufio.Scanner) {
	println("添加用户")
	var name, ageStr, career string

	fmt.Print("请输入姓名：")
	scanner.Scan()
	name = strings.TrimSpace(scanner.Text())
	fmt.Print("请输入年龄：")
	scanner.Scan()
	ageStr = strings.TrimSpace(scanner.Text())
	var age int
	var err error
	age, err = strconv.Atoi(ageStr)
	if err != nil || age < 0 {
		fmt.Println("年龄输入有误")
		return
	}
	fmt.Print("请输入职业：")
	scanner.Scan()
	career = strings.TrimSpace(scanner.Text())
	var id int
	id = (*userList)[len(*userList)-1].Id + 1
	var user User = *NewUser(id, name, age, career)
	*userList = append(*userList, user)
	var dataSerialization string
	dataSerialization, _ = userDataSerialization(*userList)
	err = userFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}

func modifyUser(userList *[]User, scanner *bufio.Scanner) {
	println("修改用户")
	if len(*userList) == 0 {
		fmt.Println("没有用户信息")
		return
	}
	fmt.Print("请输入要修改用户的id：")
	scanner.Scan()
	var id int
	var err error
	id, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || id <= 0 || id > len(*userList) {
		fmt.Println("输入的id无效")
		return
	}
	var name string
	var age int
	var career string
	var user *User
	for i, v := range *userList {
		if v.Id == id {
			user = &((*userList)[i])
			break
		}
	}
	fmt.Printf("请输入名称，回车不修改 \"%s\" ：\n", user.Name)
	scanner.Scan()

	var nameInput string = strings.TrimSpace(scanner.Text())
	if nameInput != "" {
		name = nameInput
	}
	fmt.Printf("请输入年龄，回车不修改 \"%d\" ：\n", user.Age)
	scanner.Scan()

	var ageStrInput string = strings.TrimSpace(scanner.Text())

	if ageStrInput != "" {
		age, err = strconv.Atoi(ageStrInput)
		if err != nil || age < 0 {
			fmt.Println("年龄输入有误")
			return
		}
	}
	fmt.Printf("请输入职业，回车不修改 \"%s\" ：\n", user.Career)
	scanner.Scan()

	var careerInput string = strings.TrimSpace(scanner.Text())
	if careerInput != "" {
		career = careerInput
	}
	user.ModifyUser(NewUser(user.Id, name, age, career))
	var dataSerialization string
	dataSerialization, _ = userDataSerialization(*userList)
	err = userFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}

func deleteUser(userList *[]User, scanner *bufio.Scanner) {
	fmt.Println("删除用户")
	if len(*userList) == 0 {
		fmt.Println("没有用户信息")
		return
	}
	fmt.Print("请输入要删除用户的id：")
	scanner.Scan()
	var id int
	var err error
	id, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || id <= 0 {
		fmt.Println("输入的id无效")
		return
	}
	for i, v := range *userList {
		if v.Id == id {
			id = i
		}
	}
	*userList = append((*userList)[:id], (*userList)[id+1:]...)
	var dataSerialization string
	dataSerialization, _ = userDataSerialization(*userList)

	err = userFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}
