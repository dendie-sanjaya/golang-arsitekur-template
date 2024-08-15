package migrate

import (
	"den/config"
	"den/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func db_and_table() {
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " sslmode=" + config.PostgresSSLMode
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connection established", config.PostgresHost, config.PostgressPort, "ssl mode", config.PostgresSSLMode)
	}

	var exists bool
	err = gormDB.Raw("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = ?)", config.PostgresDB).Scan(&exists).Error
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database exist", config.PostgresDB, "is", exists)
	}

	if !exists {
		err = gormDB.Exec("CREATE DATABASE " + config.PostgresDB).Error
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("Database created " + config.PostgresDB + " successfully")
		}
	}

	dsn = "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname= " + config.PostgresDB + " sslmode=" + config.PostgresSSLMode
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	gormDB.AutoMigrate(&entity.User{})
	log.Println("Migrate table user")
}

func Run_db() {
	db_and_table()
}
