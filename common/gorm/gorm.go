package gorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bsthun/gut"
	"github.com/sunthewhat/secure-docs-api/common"
	"github.com/sunthewhat/secure-docs-api/type/shared/query"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() {
	// Configure Configuration file
	lg := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Config GORM Connector
	connector := postgres.New(
		postgres.Config{
			DSN:                  *common.Config.Postgres,
			PreferSimpleProtocol: true,
		},
	)

	// Open connection
	db, connectionErr := gorm.Open(connector, &gorm.Config{
		Logger: lg,
	})

	if connectionErr != nil {
		gut.Fatal("Failed to connect to database", connectionErr)
	}

	fmt.Println("GORM Connected!")

	common.Gorm = query.Use(db)
}
