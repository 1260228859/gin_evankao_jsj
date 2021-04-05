package model

import (
	// "fmt"
	"img_tag/common/variable"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Phone string
	Name  string
}

// func Init() {
// 	variable.Db.AutoMigrate(&User{})

// 	var user User
// 	variable.Db.Find(&user, 1)
// 	fmt.Println(user, 11111)
// }
