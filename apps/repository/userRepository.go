package repository

import (
	"database/sql"
	"depogunabangunan/apps/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, fullname FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.FullName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, fullname FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.FullName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (username, fullname) VALUES ($1, $2)"
	_, err := r.db.Exec(query, user.Username, user.FullName)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	query := "UPDATE users SET username = $1, fullname = $2 WHERE id = $3"
	_, err := r.db.Exec(query, user.Username, user.FullName, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
