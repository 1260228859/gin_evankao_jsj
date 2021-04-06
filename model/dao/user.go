package dao

import (
	"fmt"
	"img_tag/common/variable"
	"img_tag/model/schema"
)

// 定义增删改查文件

func CreateUserDaoFactory() *UserDao {

	return &UserDao{}
}

type UserDao struct {
}

func (u *UserDao) Find() {
	// res := variable.Db.Model(&u.schema).Where("id=?", 1)

	
	variable.Db.Find(&schema.User{})

	fmt.Println(user, 2222)
}
