package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Store interface {
	GetById(id int)(*User, error)
}

type User struct {
	ID   string
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type UserRepository interface {
	GetByID(id int) (*User, error)
}

type InMemoryRepository struct {
	users []User
}

func (r *InMemoryRepository) GetByID(id int) (*User, error) {
	return nil, nil
}


func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepository := NewPostgresUserRepository(db)
	userService := NewUserService(userRepository)

	inMemoryRepository := &InMemoryRepository{}

	app := &application {
		// store: userRepository
		store: inMemoryRepository
	}


	user, err := userService.getUserById(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User: %+v\n", user)
	}
}
