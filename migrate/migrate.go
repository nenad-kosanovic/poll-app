package main

import (
	"fmt"
	"gin-api-mysql-crud/initializers"
	"gin-api-mysql-crud/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Statement.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	initializers.DB.AutoMigrate(models.Poll{})
	fmt.Println("? Migration complete")
}
