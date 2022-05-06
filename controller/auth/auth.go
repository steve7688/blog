/*
 * @Author: Steve
 * @Date: 2022-05-04 20:30:56
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 21:56:46
 * @FilePath: /blog/controller/auth/auth.go
 * @Description:登录授权相关
 *
 * Copyright (c) 2022 by 深圳贤齐科技有限公司, All Rights Reserved.
 */
package auth

import (
	model "Blog/models"
	resp "Blog/response"

	"github.com/gin-gonic/gin"
)

/**
 * @description:登录
 * @param {*gin.Context} c
 * @return {*}
 */
func Login(c *gin.Context) {

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	//调用登录方法
	err, userModel := model.Login(mobile, password)
	//返回给客户端
	resp.Response(c, err, userModel)
}
