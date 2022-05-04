/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 11:18:15
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:55:30
 * @FilePath: /blog/error/error.go
 * @Description:错误处理
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package error

import "github.com/gin-gonic/gin"

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
 * @Description: 根据错误码创建一个错误对象
 * @User: Steve
 * @Date: 2022-05-04 17:21:26
 * @param {int} code
 * @return {*}
 */
func CreateNewErrorWithCode(code int) (myErr *MyError) {

	chinese := errs_chinese[code]
	english := errs_english[code]

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
 * @Description: 错误提示国际化
 * @User: Steve
 * @Date: 2022-05-04 17:33:16
 * @param {*gin.Context} c
 * @return {*}
 */
func (err *MyError) Localization(c *gin.Context) string {

	language := c.Request.Header.Get("language")

	var msg = err.Chinese //默认中文提示

	if language == "en" { //英文

		msg = err.English
	}

	return msg
}
