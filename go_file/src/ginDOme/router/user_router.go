package router

import (
	"gindome/controller"
	"github.com/gin-gonic/gin"
)

/*
 */
func userApiRoute(r *gin.Engine) {
	userRoute := r.Group("user")
	// 根据id查询用户详情信息
	userRoute.GET("/:id", controller.GetUserDetailHandler)
	// 查询用户列表信息
	userRoute.GET("/", controller.GetUserHandler)
	// 根据id修改用户详情信息
	userRoute.PATCH("/", controller.UpdateUserHandler)
	// 根据id删除用户
	userRoute.DELETE("/", controller.DeleteUserHandler)
}
