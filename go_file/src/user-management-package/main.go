package main

import (
	"bufio"
	"fmt"
	"managementpackage/model"
	"managementpackage/utils"
	"managementpackage/service"
	"os"
	"strconv"
	"strings"
)

const filePath string = "db/user_file.json"

func main() {
	fmt.Println("客户关系管理系统")
	var userList []model.User
	var data []byte
	var err error
	data, err = utils.UserFileReading(filePath)
	if err != nil {
		userList = []model.User{
			*model.NewUser(1, "yuan", 23, "CEO"),
		}
	} else {
		userList, _ = utils.UserDataDeserialization(data)
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
			service.ListUser(&userList)
		case 2:
			service.AddUser(&userList, scanner,filePath)
		case 3:
			service.ModifyUser(&userList, scanner,filePath)
		case 4:
			service.DeleteUser(&userList, scanner,filePath)
		case 5:
			os.Exit(0)
		default:
			println("未知选项")
		}
	}

}

