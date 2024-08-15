package migrate

import (
	"den/config"
	"den/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname=" + config.PostgresDB + " sslmode=" + config.PostgresSSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

// Seed user data
func seedUsers(db *gorm.DB) {
	users := []entity.User{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
		{Name: "Alice Johnson", Email: "alice@example.com"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to create user %s: %v", user.Name, err)
		} else {
			fmt.Printf("User %s created successfully\n", user.Name)
		}
	}
}

func Run_seed() {
	db := initDatabase()
	seedUsers(db)
}
