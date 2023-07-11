package db

import (
	"fmt"
	"github.com/gustavocortarelli/go-agency/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Repository struct {
	db     *gorm.DB
	logger logger.Interface
}

var R Repository

func (r *Repository) GetSession() *gorm.DB {
	return r.db.Session(&gorm.Session{
		Logger:               r.logger,
		FullSaveAssociations: true,
		CreateBatchSize:      500,
	})
}

func configLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Lmicroseconds), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	return newLogger
}

func OpenConnection() error {
	config := configs.GetDB()

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Database)
	log := configLogger()
	conn, err := gorm.Open(postgres.Open(strConn), &gorm.Config{Logger: log})
	R = Repository{
		db:     conn,
		logger: log,
	}
	if err != nil {
		//TODO: change it to error handler
		panic(err)
	}

	pgDB, err := R.db.DB()
	pgDB.SetMaxIdleConns(5)
	pgDB.SetMaxOpenConns(30)
	pgDB.SetConnMaxLifetime(time.Minute * 30)
	return err
}
