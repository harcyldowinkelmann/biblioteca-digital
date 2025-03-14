package config

import (
	"ApiSup/pkg/mapear/constants"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_SENHA"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), /// Mostrar logs de operações do banco
	})
}

var DB *gorm.DB

func Initialize() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(constants.ERROR_CARREGAMENTO_ENV)
	}

	dbConn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	DB = dbConn

	return DB
}

func Close() error {
	sqlDB, err := DB.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
