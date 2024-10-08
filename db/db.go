// package db

// import (
// 	"fmt"
// 	"log"

// 	"github.com/RND2002/goChatApp/models"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// type Config struct {
// 	Host     string
// 	Port     string
// 	Password string
// 	User     string
// 	DBName   string
// 	SSLMode  string
// }

// // func NewConnection(config Config) (*gorm.DB, error) {
// // 	//fmt.Printf("%s %s %s %s %s %s", config.Host, config.Port, config.Password, config.DBName, config.SSLMode, config.User)
// // 	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=%s user=%s",
// // 		config.Host, config.Port, config.Password, config.DBName, config.SSLMode, config.User)

// // 	log.Println("Connecting to db...")

// // 	var db *gorm.DB
// // 	var err error
// // 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

// // 	if err != nil {
// // 		fmt.Println(err)
// // 	}

// // 	DB = db
// // 	fmt.Println("Database connected successfully")
// // 	return db, nil

// // }
// func NewConnection(config *Config) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

// 	log.Println("Connecting to db with DSN:", dsn)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Println("Failed to connect to database:", err)
// 		return nil, err
// 	}

// 	log.Println("Database connected successfully")
// 	return db, nil
// }

// func InitializeDB() {
// 	// err := godotenv.Load(".env")
// 	// if err != nil {
// 	// 	fmt.Println("Error loading environment data")
// 	// }

// 	// config := &Config{
// 	// 	Host:     os.Getenv("DB_HOST"),
// 	// 	Port:     os.Getenv("DB_PORT"),
// 	// 	Password: os.Getenv("DB_PASSWORD"),
// 	// 	User:     os.Getenv("DB_USER"),
// 	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
// 	// 	DBName:   os.Getenv("DB_NAME"),
// 	// }

// 	config := &Config{
// 		Host:     "localhost",
// 		Port:     "5432",
// 		Password: "root",
// 		User:     "root",
// 		SSLMode:  "disable",
// 		DBName:   "go-chat-app",
// 	}

// 	db, err := NewConnection(config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		err = models.AutoMigrate(db)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package db

import (
	"fmt"
	"log"

	"github.com/RND2002/goChatApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// Config holds the database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection establishes a connection to the PostgreSQL database
func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	log.Println("Connecting to database with DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}

// InitializeDB initializes the database connection and performs auto-migration
func InitializeDB() {
	// Directly set your database configuration
	config := &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "root",
		DBName:   "go-chat-app",
		SSLMode:  "disable", // Change to "require" if SSL is needed
	}

	// Establish the database connection
	var err error
	DB, err = NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	// Perform auto-migration
	err = models.AutoMigrate(DB)
	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}
}
