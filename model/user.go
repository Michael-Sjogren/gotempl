package model

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Con *sql.DB
}

type User struct {
	Id        int64
	Username  string
	Access    int
	CreatedAt string
}

const ()

func CheckPassword(hash []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return false
	}
	return true
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func (m *UserModel) CreateUser(newUser User, passwordHash []byte) (User, error) {
	res, err := m.Con.Exec(`
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
