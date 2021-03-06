// @Author: steve
// @Date: 2022-05-04 15:34:33
// @LastEditors: steve
// @LastEditTime: 2022-05-04 17:39:19
// @FilePath: /blog/response/response.go
// @Description:本文件专门处理response到客户端
// Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.

package response

import (
	"Blog/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description:返回json给客户端
 * @param {*gin.Context} c
 * @param {*error.MyError} err
 * @param {interface{}} data
 * @return {*}
 * @user: Steve
 * @Date: 2022-05-06 22:16:27
 */
func Response(c *gin.Context, err *error.MyError, data interface{}) {

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"code": err.Code,
			"msg":  err.Localization(c), //错误提示做国际化
			"data": nil,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  nil,
			"data": data,
		})
	}

}
