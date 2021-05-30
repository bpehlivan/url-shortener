package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandStringBytes(n int) string {
	randString := make([]byte, n)
	for i := range randString {
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}

func GetConnection() (conn *gorm.DB) {
	dbHost := viper.Get("POSTGRES_DB_HOST")
	dbPort := viper.Get("POSTGRES_DB_PORT")
	dbUser := viper.Get("POSTGRES_DB_USER")
	dbPassword := viper.Get("POSTGRES_DB_PASSWORD")
	dbName := viper.Get("POSTGRES_DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errContent := "Error on DB connection: \n"
		errContent += err.Error()
		panic(errContent)
	}
	return
}

func CloseConnection(conn *gorm.DB)  {
	db, _ := conn.DB()
	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}