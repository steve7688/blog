/*
 * @Author: Steve
 * @Date: 2022-05-06 23:06:52
 * @LastEditors: Steve
 * @LastEditTime: 2022-05-06 23:24:42
 * @FilePath: /blog/db/db.go
 * @Description:数据库相关
 *
 * Copyright (c) 2022 by 深圳贤齐科技有限公司, All Rights Reserved.
 */
package db

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	fmt.Println("init db")
}
