package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xorwise/golang-todo-api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var (
		dbHost = env.DBHost
		dbPort = env.DBPort
		dbUser = env.DBUser
		dbPass = env.DBPass
		dbName = env.DBName
	)

	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	return db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close()
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(domain.Models...)
	return err
}
