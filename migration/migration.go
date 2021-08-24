package migration

import (
	"myapp/config"
	"myapp/graph/model"
)

func MigrateTable() {
	db := config.GetDB()

	db.AutoMigrate(&model.User{})
}
