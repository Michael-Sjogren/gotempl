package model

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func NewUserRepo(con *sql.DB) UserRepo {
	return UserRepo{
		con: con,
	}
}

type UserRepo struct {
	con *sql.DB
}

type User struct {
	Id        int64
	Username  string
	Access    int
	CreatedAt string
}

func CheckPassword(hash []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return false
	}
	return true
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+bcrypt.MinCost)
}

func (m *UserRepo) CreateUser(newUser User, passwordHash []byte) (User, error) {
	res, err := m.con.Exec(`
		INSERT INTO users (username,password_hash,access) VALUES (?,?,?);
	`, newUser.Username, passwordHash, newUser.Access)

	if err != nil {
		return newUser, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	newUser.Id = id
	return newUser, nil
}

func (h *UserRepo) GetAll() ([]User, error) {
	var users []User
	rows, err := h.con.Query(`SELECT id,username,access,created_at FROM users;`)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.Access,
			&user.CreatedAt,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	return users, nil
}
