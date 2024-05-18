package accessRepository

import (
	"database/sql"
	access "gugu/interfaces/access"
)

type AccessRepository interface {
	InsertAccessToken(access *access.Access) error
	GetAccessById(id string) (*access.Access, error)
	GetAccessByToken(token string) (*access.Access, error)
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

func (r *accessRepository) GetAccessById(id string) (*access.Access, error) {
	var a access.Access

	query := `SELECT
	access_id, 
	user_id,
	access_token,
	expires_at,
	revoked,
	session_id
	FROM tb_access WHERE access_id = $1 and revoked = false`

	err := r.DB.QueryRow(query, id, "active").Scan(
		&a.AccessID,
		&a.UserID,
		&a.AccessToken,
		&a.ExpiresAt,
		&a.Revoked,
		&a.SessionID,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *accessRepository) GetAccessByToken(token string) (*access.Access, error) {
	var a access.Access

	query := `SELECT
	access_id, 
	user_id,
	access_token,
	expires_at,
	revoked,
	session_id
	FROM tb_access WHERE access_token = $1 and revoked = false`

	err := r.DB.QueryRow(query, token).Scan(
		&a.AccessID,
		&a.UserID,
		&a.AccessToken,
		&a.ExpiresAt,
		&a.Revoked,
		&a.SessionID,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func NewRepository(DB *sql.DB) AccessRepository {
	return &accessRepository{
		DB: DB,
	}
}
