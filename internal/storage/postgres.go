package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gene-qxsi/CRM-M/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

var (
	driver      = "pgx"
	storagePath = "postgresql://postgres:admin@localhost:5432/crm"
)

func New() *Storage {
	const op = "internal.storage.postgres.New"

	db, err := sql.Open(driver, storagePath)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil
	}

	return &Storage{db: db}
}

// TODO: в обработки ошибок, добавить запись хедеров
func (s *Storage) CreateUser(user models.User) (int, error) {
	const op = "internal.storage.postgres.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(name, age, password, role) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return 0, err
	}
	var id int
	err = stmt.QueryRow(user.Name, user.Age, user.Password, user.Role).Scan(&id)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetUser(id int) (*models.User, error) {
	op := "internal.storage.postgres.GetUser"

	row := s.db.QueryRow("SELECT id, name, age, password, role FROM users WHERE id = $1", id)
	if err := row.Err(); err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil, err
	}
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Password, &user.Role)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil, err
	}

	return &user, nil
}

func (s *Storage) GetUsers() ([]models.User, error) {
	op := "internal.storage.postgres.GetUsers"

	rows, err := s.db.Query("SELECT id, name, age, password, role FROM users")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil, err
	}
	var users []models.User
	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Password, &user.Role)
		if err != nil {
			log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *Storage) DeleteUser(id int) error {
	const op = "internal.storage.postgres.DeleteUser"

	stmt, err := s.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return err
	}

	lines, err := result.RowsAffected()
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return err
	}

	if lines <= 0 {
		log.Printf("❌ ERROR: %s. PATH: %s\n", "ни одного пользователя не было удалено", op)
		return fmt.Errorf("ни одного пользователя не было удалено")
	}

	return nil
}

func (s *Storage) GetUserByNameAndPassword(name, password string) (*models.User, error) {
	const op = "internal.storage.postgres.GetUserByNameAndPassword"

	row := s.db.QueryRow("SELECT id, name, age, password, role FROM users WHERE name = $1 AND password = $2", name, password)
	if err := row.Err(); err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil, err
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Password, &user.Role)
	if err != nil {
		log.Printf("❌ ERROR: %s. PATH: %s\n", err, op)
		return nil, err
	}

	return &user, nil
}
