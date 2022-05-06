/*
 * @Author: Steve
 * @Date: 2022-05-04 20:30:56
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 22:05:53
 * @FilePath: /blog/routers/router.go
 * @Description:路由相关
 *
 * Copyright (c) 2022 by 深圳贤齐科技有限公司, All Rights Reserved.
 */
package routers

import (
	"Blog/controller/auth"

	"github.com/gin-gonic/gin"
)

/**
 * @Description: 设置路由
 * @Author: Steve
 * @Date: 2022-05-06 22:05:51
 * @Param:
 * @Return:
 */
func SetUpRouters() {

	r := gin.Default()
	v1 := r.Group("/v1")

	//auth
	v1.POST("/login", auth.Login)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
