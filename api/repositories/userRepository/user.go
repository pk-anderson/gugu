package userRepository

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) InsertUser(uuid, username, email, password string) error {
	_, err := repo.DB.Exec("INSERT INTO tb_users (id, username, email, password) VALUES ($1, $2, $3, $4)", uuid, username, email, password)
	if err != nil {
		return err
	}
	return nil
}
