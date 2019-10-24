package stores

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewPostgresqlConnection creates a new handle to the DB object
func CreatePostgresDBConnection(dbHost string, dbPort string, username string, password string, dbName string) *gorm.DB {
	fmt.Printf("Connection to Postgres : host=%s port=%s user=%s password=XXX dbname=%s sslmode=disable", dbHost, dbPort, username, dbName)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	return db
}

// NewPostgresqlConnection creates a new handle to the DB object using .env variables
func NewPostgresqlConnection() *gorm.DB {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	return CreatePostgresDBConnection(dbHost, dbPort, username, password, dbName) //Build connection string
}

func GetDAO(daoType string) *gorm.DB {
	switch daoType {
	case "postgres":
		return NewPostgresqlConnection()
	default:
		return NewPostgresqlConnection()
	}
}
