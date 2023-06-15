package service

import (
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
)

/*
 * @description: 注册用户服务

 * @param: u *models.User 用户信息

 * @return: *pkg.Token 认证令牌

 * @return: pkg.Error 错误信息
 */
func RegisterUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 1. 检查用户是否已经存在
	_, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData == 1 {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeUserExist)
	}

	// 2. 创建新用户
	user := models.User{
		Account:  u.Account,
		Password: u.Password,
		UserProfile: models.UserProfile{
			Name: u.Account,
		},
	}
	user.HashPassword()

	// 3. 保存新用户到数据库中
	if err := repository.RegisterUser(&user); err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 4. 获取新用户的ID
	uId, err := repository.GetIDByAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 5. 生成认证令牌
	token, err := pkg.GetToken(uId, u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 6. 返回结果和错误信息
	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

/*
 * @description: 登录用户服务

 * @param: u *models.User 用户信息

 * @return: *pkg.Token 认证令牌

 * @return: pkg.Error 错误信息
 */
func LoginUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 1. 哈希加密用户密码
	u.HashPassword()

	// 2. 检查用户是否已经存在
	user, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData != 1 {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeUserNotExist)
	}

	// 3. 比较用户输入的账号和密码是否与数据库中的记录匹配
	if user.Account != u.Account || user.Password != u.Password {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeInvalidPassword)
	}

	// 4. 生成认证令牌
	token, err := pkg.GetToken(user.ID, user.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

/*
 * @description: 根据用户 ID 获取用户信息

 * @param: userId uint64 用户 ID

 * @return: *models.UserProfile 用户信息

 * @return: error 错误信息
 */
func GetUserByIdService(userId uint64) (*models.UserProfile, error) {
	// 1. 调用 repository 层获取用户信息
	data, err := repository.GetUserById(userId)
	// 2. 如果出现错误，返回错误信息
	if err != nil {
		return nil, err
	}
	// 3. 返回用户信息和 nil
	return data, nil
}

/*
 * @description: 获取用户列表服务

 * @param: page int 分页页码

 * @param: size int 分页大小

 * @return: []*models.UserProfile 用户列表

 * @return: error 错误信息
 */
func GetUserListService(page, size int) ([]*models.UserProfile, error) {
	// 1. 调用 repository 层获取用户列表
	data, err := repository.GetUserList(page, size)
	// 2. 如果出错，返回错误信息
	if err != nil {
		return nil, err // 返回值为 nil 和错误信息
	}
	// 3. 返回用户列表和 nil
	return data, nil
}
