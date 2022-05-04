/*
 * @Author: steve steve7688@163.com
 * @Date: 2022-05-04 10:49:45
 * @LastEditors: steve steve7688@163.com
 * @LastEditTime: 2022-05-04 17:30:49
 * @FilePath: /blog/routers/router.go
 * @Description:路由相关
 *
 * Copyright (c) 2022 by steve steve7688@163.com, All Rights Reserved.
 */
package routers

import (
	"Blog/controller/auth"

	"github.com/gin-gonic/gin"
)

/**
 * @Description:设置路由
 * @User: Steve
 * @Date: 2022-05-04 17:23:11
 * @param {*}
 * @return {*}
 */
func SetUpRouters() {

	r := gin.Default()
	v1 := r.Group("/v1")

	//auth
	v1.POST("/login", auth.Login)

	r.Run(":8080")

}
