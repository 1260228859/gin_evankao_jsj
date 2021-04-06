package service

import (
	"fmt"
	"img_tag/common/variable"
	"img_tag/model/schema"
)

// 定义增删改查文件

func CreateUserDaoFactory() *UserDao {

	return &UserDao{}
}

type User schema.User
type UserDao struct {
}


func (u *UserDao) Find() (*User) {
	// res := variable.Db.Model(&u.schema).Where("id=?", 1)
	
	var UserScheam User 
	variable.Db.Find(&UserScheam)

	fmt.Println(UserScheam, 2222)
	return &UserScheam
}

func Find(schema *interface{}) (*interface{}) {
	variable.Db.Find(schema)
	return schema
}
