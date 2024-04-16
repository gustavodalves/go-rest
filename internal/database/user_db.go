package database

import (
	"database/sql"

	"github.com/gustavodalves/go-api/internal/entity"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDb(db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (ud *UserDB) Insert(user entity.User) error {
	_, err := ud.db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (ud *UserDB) GetById(id uint64) (*entity.User, error) {
	var user entity.User

	err := ud.db.QueryRow("SELECT email, password FROM users WHERE id = ?", id).Scan(&user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ud *UserDB) GetAll() ([]*entity.User, error) {
	var users []*entity.User

	rows, err := ud.db.Query("SELECT email, password FROM users")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entity.User

		if err := rows.Scan(&user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
