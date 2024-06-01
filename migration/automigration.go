package migration

import (
	"github.com/bhumong/dbo-test/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	err := db.Migrator().AutoMigrate(&model.User{}, &model.Customer{})
	if err != nil {
		log.Fatal().AnErr("db.Migrator().AutoMigrate", err).Send()
		panic("panic due too failed automigrate")
	}
}
