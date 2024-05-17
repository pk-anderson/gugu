package accessRepository

import (
	"database/sql"
	access "gugu/interfaces/access"
)

type AccessRepository interface {
	InsertAccessToken(access *access.Access) error
}

type accessRepository struct {
	DB *sql.DB
}

func (r *accessRepository) InsertAccessToken(access *access.Access) error {
	query := `INSERT INTO tb_access (user_id, access_token, expires_at, revoked, session_id) 
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(
		query,
		access.UserID,
		access.AccessToken,
		access.ExpiresAt,
		access.Revoked,
		access.SessionID)
	return err
}

func NewRepository(DB *sql.DB) AccessRepository {
	return &accessRepository{
		DB: DB,
	}
}
