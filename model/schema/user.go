package schema

import (
	// "fmt"
	// "img_tag/pkg/variable"

	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string  `gorm:"column:user_name;size:64;index;default:'';not null;"` // 用户名
	RealName string  `gorm:"column:real_name;size:64;index;default:'';not null;"` // 真实姓名
	Password string  `gorm:"column:password;size:40;default:'';not null;"`        // 密码(sha1(md5(明文))加密)
	Email    *string `gorm:"column:email;size:255;index;"`                        // 邮箱
	Phone    *string `gorm:"column:phone;size:20;index;"`                         // 手机号
	Status   int     `gorm:"column:status;index;default:0;not null;"`             // 状态(1:启用 2:停用)
	Creator  string  `gorm:"column:creator;size:36;"`                             // 创建者
}

// func Init() {
// variable.Db.AutoMigrate(&User{})

// var user User
// variable.Db.Find(&user, 1)
// fmt.Println(user, 11111)
// }
