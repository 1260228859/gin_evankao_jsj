package migrate

import (
	"img_tag/model/schema"
	"img_tag/pkg/variable"
)

func InitMysqlAutoMigrate() {
	variable.Db.AutoMigrate(
		&schema.User{},
	)
}
