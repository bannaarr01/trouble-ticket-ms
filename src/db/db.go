package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"trouble-ticket-ms/src/config"
)

type DB struct {
	*gorm.DB
}

func Init() *DB {
	cfg := config.New()

	// MySQL DSN format
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Get the underlying *sql.DB object
	sqlDB, err := dbConn.DB()

	if err != nil {
		log.Panic(err.Error())
	}

	MaxOpenCon, err := strconv.Atoi(cfg.DB.MaxOpenCon)

	if err != nil {
		log.Printf("Error converting Max Open Con value to int: %v. fallback to default", err)
		MaxOpenCon = 10
	}

	MaxIdleCon, err := strconv.Atoi(cfg.DB.MaxIdleCon)

	if err != nil {
		log.Printf("Error converting Max Idle Con value to int: %v. fallback to default", err)
		MaxIdleCon = 5
	}

	sqlDB.SetMaxOpenConns(MaxOpenCon)
	sqlDB.SetMaxIdleConns(MaxIdleCon)

	log.Printf("Successfully connected to %s database on %s:%s", cfg.DB.Name, cfg.DB.Host, cfg.DB.Port)

	return &DB{dbConn}
}
