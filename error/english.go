/*
 * @Author: Steve
 * @Date: 2022-05-04 20:30:56
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 22:45:15
 * @FilePath: /blog/error/english.go
 * @Description:英文错误码
 *
 * Copyright (c) 2022 by 深圳贤齐科技有限公司, All Rights Reserved.
 */
package error

//---------------错误码对应的[英文]释义---------------\\

var english = map[int]string{

	//global
	UNKNOW_ERROR: "Unknown error",

	//auth
	AUTH_MOBILE_NULL:               "Mobile number cannot be null",
	AUTH_PASSWORD_NULL:             "Password cannot be null",
	AUTH_MOBILE_NOT_EXIST:          "Mobile number not exist",
	AUTH_MOBILE_PASSWORD_NOT_MATCH: "Mobile number and password not mathed",
	AUTH_MOBILE_BLOCKED:            "Mobile number is blocked,Please contact customer service center",
	AUTH_MOBILE_BLOCKED_FOREVER:    "Mobile number is blocked forever,Please contact customer service center",
}
