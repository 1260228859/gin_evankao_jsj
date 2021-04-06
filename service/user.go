package service

import (
	"fmt"
	"img_tag/common/variable"
	"img_tag/model/schema"
)

// 定义增删改查文件

func User() *UserService {

	return &UserService{}
}

type UserSchema schema.User
type UserService struct {
}


func (u *UserService) Find() (*UserSchema) {
	// res := variable.Db.Model(&u.schema).Where("id=?", 1)
	
	var UserScheam UserSchema 
	variable.Db.Find(&UserScheam)

	fmt.Println(UserScheam, 2222)
	return &UserScheam
}

func Find(schema *interface{}) (*interface{}) {
	variable.Db.Find(schema)
	return schema
}
