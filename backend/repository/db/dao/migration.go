package dao

import (
	"fmt"
	"os"

	"9900project/repository/db/model"
)

// Migration 执行数据迁移
func Migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Forum{}, &model.Post{}, &model.Comment{}, &model.Course{}, &model.CourseSelect{}, &model.Material{}, &model.Assignment{}, &model.AssMark{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
