/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 17:44:19
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:56:05
 * @FilePath: /blog/error/english.go
 * @Description:错误提示多语言--英文配置文件
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package error

//---------------错误码对应的[英文]释义---------------\\
var errs_english = map[int]string{

	//global
	UNKNOW_ERROR: "Unknow error",

	//auth
	AUTH_MOBILE_NULL:               "Mobile number cannot be null",
	AUTH_PASSWORD_NULL:             "Password cannot be null",
	AUTH_MOBILE_NOT_EXIST:          "Mobile number not exist",
	AUTH_MOBILE_PASSWORD_NOT_MATCH: "Mobile number and password not mathed",
	AUTH_MOBILE_BLOCKED:            "Mobile number is blocked,Please contact customer service center",
	AUTH_MOBILE_BLOCKED_FOREVER:    "Mobile number is blocked forever,Please contact customer service center",
}
