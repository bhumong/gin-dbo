package db

import (
	"time"

	"github.com/bhumong/dbo-test/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var newLogger = logger.New(
	&log.Logger,
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Error,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  false,
	},
)

func Init() {
	var err error

	c := config.GetConfig()

	dsn := c.GetString("db.dsn")
	log.Debug().Msg(dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal().AnErr("failed to connect database", err).Send()
	}

	stdDb, err := db.DB()
	if err != nil {
		panic("failed to get *sql.DB")
	}

	if err = stdDb.Ping(); err != nil {
		panic("failed to ping db")
	}
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
