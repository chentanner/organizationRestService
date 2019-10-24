package main

import (
	"fmt"
	"os"
    "strconv"

	"organizationRestService/api/handlers"
	"organizationRestService/models"
	"organizationRestService/router"
	"organizationRestService/services"
	"organizationRestService/models/stores"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Println(e)
	}

	dbType := os.Getenv("db_type") //Get database type from .env file, we did not specify any port so this should return an empty string when tested locally
	if dbType == "" {
		fmt.Println("No environment variable found for db_type, defaulting to postgres")
		dbType = "postgres" //default database is postgres
	}

	autoMigrateEnv := os.Getenv("AUTOMIGRATE") //Get auto migration setting from .env file.
	if autoMigrateEnv == "" {
		fmt.Println("No environment variable found for AUTOMIGRATE, defaulting to false")
		autoMigrateEnv = "false" //default to not auto migrate
	}
	autoMigrate, err := strconv.ParseBool(autoMigrateEnv)
	if err != nil {
		autoMigrate = false
	}
	fmt.Println(autoMigrateEnv,autoMigrate)

	// Dependency injection of components
	db := stores.GetDAO(dbType)
	if autoMigrate {
		fmt.Println("Auto migrating")
		db.AutoMigrate(&models.Organization{})
	}

	organizationStore := stores.NewOrganizationStore(db)

	serviceValidator := services.NewServiceValidator()

	organizationService := services.NewOrganizationService(organizationStore, serviceValidator)
	organizationHandlers := handlers.NewOrganizationHandler(organizationService)

	r := router.BuildRouter(organizationHandlers)

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //default port
	}

	err = r.Start(":"+port)
	if err != nil {
		panic(err)
	}
}
