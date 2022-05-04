/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 17:43:34
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:56:11
 * @FilePath: /blog/error/chinese.go
 * @Description:错误提示多语言--中文配置文件
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package error

//---------------错误码对应的[中文]释义---------------\\
var errs_chinese = map[int]string{

	//global
	UNKNOW_ERROR: "未知错误",

	//auth
	AUTH_MOBILE_NULL:               "登录手机号不能为空",
	AUTH_PASSWORD_NULL:             "登录密码不能为空",
	AUTH_MOBILE_NOT_EXIST:          "手机号不存在",
	AUTH_MOBILE_PASSWORD_NOT_MATCH: "手机号和密码不匹配",
	AUTH_MOBILE_BLOCKED:            "手机号已被封禁，请联系客服解封",
	AUTH_MOBILE_BLOCKED_FOREVER:    "该账号已被永久封禁，请联系客服了解详情",
}
