package service

import (
	"bufio"
	"fmt"
	"managementpackage/model"
	"managementpackage/utils"
	"strconv"
	"strings"
)

func ListUser(userList *[]model.User) {
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

func AddUser(userList *[]model.User, scanner *bufio.Scanner,filePath string) {
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
	var id int = (*userList)[len(*userList)-1].Id + 1
	var user model.User = *model.NewUser(id, name, age, career)
	*userList = append(*userList, user)
	var dataSerialization string
	dataSerialization, _ = utils.UserDataSerialization(*userList)
	err = utils.UserFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}

func ModifyUser(userList *[]model.User, scanner *bufio.Scanner,filePath string) {
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
	var user *model.User
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
	user.ModifyUser(model.NewUser(user.Id, name, age, career))
	var dataSerialization string
	dataSerialization, _ = utils.UserDataSerialization(*userList)
	err = utils.UserFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}

func DeleteUser(userList *[]model.User, scanner *bufio.Scanner,filePath string) {
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
	dataSerialization, _ = utils.UserDataSerialization(*userList)

	err = utils.UserFileWrite(filePath, dataSerialization)
	if err != nil {
		fmt.Println("数据写入错误")
		return
	}
}