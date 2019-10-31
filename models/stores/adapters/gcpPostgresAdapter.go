package adapters

import (
	"fmt"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jinzhu/gorm"
)

func CreatePostgresDBConnection(dbHost string, dbPort string, username string, password string, dbName string) *gorm.DB {
	instanceConnectionName := os.Getenv("db_instance_name")

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		instanceConnectionName,
		dbName,
		username,
		password,
	)

	fmt.Println(dsn)

	db, err := gorm.Open("cloudsqlpostgres", dsn)
	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	return db
}
