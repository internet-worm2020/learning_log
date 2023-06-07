package users

import (
	"fmt"
	"gindome/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserDetailHandler(c *gin.Context) {
	var userIdStr string =c.Param("id")
	var userIdInt uint64
	var err error
	userIdInt,err=strconv.ParseUint(userIdStr, 10, 64)
	if err!=nil{
		fmt.Println("a",err)
	}
	fmt.Printf("userIdInt: %v\n", userIdInt)
	controller.ResponseSuccess(c, "ok")
}

func getUserHandler(c *gin.Context){
	page, size := controller.GetPageInfo(c)
	data, err := getUserList(page, size)
	if err != nil {
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, data)
}

func createUserHandler(c *gin.Context){
	u:=new(User)
	if err := c.ShouldBindJSON(&u); err != nil {
		fmt.Println(err)
		controller.ResponseError(c, controller.CodeInvalidParam)
		return
	}
	if err := createUser(u); err != nil {
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, nil)
}