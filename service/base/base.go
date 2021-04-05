package base

import "img_tag/common/variable"

func CreateBaseDaoFactory() *BaseDao {

	return &BaseDao{}
}

type BaseDao struct {
}

func (d *BaseDao) Find() {
	variable.Db.Find()
}

func (d *BaseDao) FindOne() {

}

func (d *BaseDao) Update() {

}

func (d *BaseDao) UpdateMany() {

}

func (d *BaseDao) Delete() {

}
