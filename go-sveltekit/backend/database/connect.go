package database

import (
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"strukt/config"
	"strukt/model"
)

// ConnectDB connect to db
func ConnectPostgres(_port string, _user string, _pass string, _name string) {
	var err error
	p := config.Config(_port)
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=0.0.0.0 port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config(_user),
		config.Config(_pass),
		config.Config(_name),
	)

	Postgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	Postgres.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
