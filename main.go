/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 10:11:12
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:31:08
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
 * @Description: 程序入口
 * @User: Steve
 * @Date: 2022-05-04 17:30:37
 * @param {*}
 * @return {*}
 */
func main() {

	routers.SetUpRouters()
}
