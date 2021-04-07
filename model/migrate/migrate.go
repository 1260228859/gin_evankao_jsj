package migrate

import (
	"img_tag/common/variable"
	"img_tag/model/schema"
)

func InitMysqlAutoMigrate() {
	variable.Db.AutoMigrate(
		&schema.User{},
	)
}
