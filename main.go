/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 10:11:12
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 22:17:02
 * @FilePath: /blog/main.go
 * @Description:程序入口文件
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package main

import (
	"Blog/routers"
)

/**
 * @description: 程序入口
 * @param {*}
 * @return {*}
 * @user: Steve
 * @Date: 2022-05-06 22:16:56
 */
func main() {

	routers.SetUpRouters()
}
