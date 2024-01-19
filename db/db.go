package db

import (
	"database/sql"
	"fmt"
	"go-echo-app/models"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// var DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME string

// func LoadEnv() (string, string, string, string, string) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	DB_HOST = os.Getenv("DB_HOST")
// 	DB_PORT = os.Getenv("DB_PORT")
// 	DB_USER = os.Getenv("DB_USER")
// 	DB_PASSWORD = os.Getenv("DB_PASSWORD")
// 	DB_NAME = os.Getenv("DB_NAME")

//		return DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
//	}
const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
)

func ConnectDB() {
	// DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME = LoadEnv()

	// Define the data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	// dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

	// Open a connection to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic((err))
	}
	defer db.Close()

	// Verify the connection to the database
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database!")
}

// InitDB initializes the database
func InitDB() {
	ConnectDB()
}

func CreateConnection() *sql.DB {
	return DB
}

// GetUserByID fetches a user by ID from the database
func GetUserByID(id int) (models.User, error) {
	// Implementation of fetching a user from the database
	user := models.User{}
	log.Println(&user.ID)
	err := DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user with the id %d", id)
		} else {
			log.Printf("Error getting user by ID: %v", err)
		}

		return models.User{}, err
	}
	return user, nil
}

// GetAllUsers fetches all users from the database
func GetAllUsers() ([]models.User, error) {
	// Implementation of fetching all users from the database
	rows, err := DB.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	log.Println(rows)
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			// return users, err
			log.Println("Error scanning user: ", err)
			continue
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
