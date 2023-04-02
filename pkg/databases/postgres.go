package databases

import (
	"time"

	"github.com/Vlad1slavZhuk/grpc-postgresql/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	postgresDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := postgresDB.Ping(); err != nil {
		log.GetLoggerInstance().Fatal().Err(err).Msg("error ping")
		return nil, err
	}

	postgresDB.SetConnMaxIdleTime(10)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}
