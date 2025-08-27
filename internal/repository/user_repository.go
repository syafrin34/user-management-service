// Package repository
package repository

import (
	"database/sql"
	"user-management-service/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// // mock user db

//	var users = map[int]*entity.User{
//		1:{ID: 1, Username: "user1", Email: "user1@gmail.com",Password: "12345"},
//		2:{ID: 2, Username: "user2", Email: "user2@gmail.com",Password: "12345"},
//	}
func (u *UserRepository) GetUserByID(id int) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, username, email, password FROM users WHERE id = ?`
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, username, email, password FROM users WHERE email = ?`
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	query := `INSERT INTO products(username, email, password)VALUES(?,?,?,?)`
	res, err := u.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (u *UserRepository) GetUserByEmailAndPassword(email, password string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, username, email, password FROM users WHERE email = ? AND password = ?`
	err := u.db.QueryRow(query, email, password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
