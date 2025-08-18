package main

import(
	"app/src/config"
	"app/src/database"
	"app/src/router"

	"fmt"
	"log"
	"gorm.io/gorm"

)
func main(){
	// db := setupDatabase()
	r := router.SetupRouter()

	addr := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	log.Printf("Server running at http://%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func setupDatabase() *gorm.DB {
	db := database.Connect(config.DBHost, config.DBName)
	// Add any additional database setup if needed
	return db
}