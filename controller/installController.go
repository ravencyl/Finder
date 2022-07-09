package controller

import (
	"Finder/base"
	"Finder/models"
)

func Install_DB() {
	db := base.Db_sqlite_open()
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Attachment{})
}
