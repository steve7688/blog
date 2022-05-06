/*
 * @Author: Steve
 * @Date: 2022-05-04 20:30:56
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 22:44:24
 * @FilePath: /blog/error/error.go
 * @Description:错误相关
 *
 * Copyright (c) 2022 by 深圳贤齐科技有限公司, All Rights Reserved.
 */
package error

import (
	"github.com/gin-gonic/gin"
)

//---------------错误码---------------\\
const (
	UNKNOW_ERROR                   = -1
	AUTH_MOBILE_NULL               = 1001
	AUTH_PASSWORD_NULL             = 1002
	AUTH_MOBILE_NOT_EXIST          = 1003
	AUTH_MOBILE_PASSWORD_NOT_MATCH = 1004
	AUTH_MOBILE_BLOCKED            = 1005
	AUTH_MOBILE_BLOCKED_FOREVER    = 1006
)

type MyError struct {
	Code    int
	Chinese string
	English string
}

/**
 * @description: 用错误码创建错误
 * @param {int} code
 * @return {*}
 * @user: Steve
 * @Date: 2022-05-06 22:17:22
 */
func CreateNewErrorWithCode(code int) (myErr *MyError) {

	chinese := chinese[code]
	english := english[code]

	if chinese == "" || english == "" { //没配置该错误码--返回未知错误
		return CreateNewErrorWithCode(UNKNOW_ERROR)
	}

	return &MyError{
		Code:    code,
		Chinese: chinese,
		English: english,
	}

}

/**
 * @description: 错误国际化
 * @param {*gin.Context} c
 * @return {*}
 * @user: Steve
 * @Date: 2022-05-06 22:17:14
 */
func (err *MyError) Localization(c *gin.Context) string {

	language := c.Request.Header.Get("language")

	var msg = err.Chinese //默认中文提示

	if language == "en" { //英文

		msg = err.English
	}

	return msg
}
