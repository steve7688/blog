/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 11:01:49
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:33:50
 * @FilePath: /blog/controller/auth/auth.go
 * @Description:登录相关
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package auth

import (
	model "Blog/models"
	resp "Blog/response"

	"github.com/gin-gonic/gin"
)

/**
 * @Description:登录入口
 * @User: Steve
 * @Date: 2022-05-04 17:27:36
 * @param {*gin.Context} c
 * @return {*}
 */
func Login(c *gin.Context) {

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	err, userModel := model.Login(mobile, password)

	resp.Response(c, err, userModel)
}
