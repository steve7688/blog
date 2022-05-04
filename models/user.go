/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 11:04:22
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:32:36
 * @FilePath: /blog/models/user.go
 * @Description:用户相关
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package models

import "Blog/error"

type UserModel struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      int    `json:"se"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

/**
 * @Description: 登录方法
 * @User: Steve
 * @Date: 2022-05-04 17:18:40
 * @param {*} mobile
 * @param {string} password
 * @return {*}
 */
func Login(mobile, password string) (myErr *error.MyError, user *UserModel) {

	if mobile == "" { //手机号为空

		return error.CreateNewErrorWithCode(error.AUTH_MOBILE_NULL), nil
	}

	if password == "" { //密码为空

		return error.CreateNewErrorWithCode(error.AUTH_PASSWORD_NULL), nil
	}

	if mobile != "13888888888" { //手机号未注册

		return error.CreateNewErrorWithCode(error.AUTH_MOBILE_NOT_EXIST), nil
	}

	if password != "123456" { //手机号密码不匹配

		return error.CreateNewErrorWithCode(error.AUTH_MOBILE_PASSWORD_NOT_MATCH), nil
	}

	//登录成功
	return nil, &UserModel{
		Mobile:   mobile,
		Password: password,
		Age:      18,
		Sex:      1,
		Email:    "steve@163.com",
		Address:  "ShenZhen",
	}
}
