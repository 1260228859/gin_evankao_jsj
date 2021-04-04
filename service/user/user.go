package user

import "img_tag/service/base"

func CreateUserDaoFactory() *UsersDao {

	return &UsersDao{}
}

type UsersDao struct {
	base.BaseDao
}

func (d *UsersDao) Find() {

}
