package userRepository

import (
	"database/sql"
	"errors"
	user "gugu/interfaces/user"
	"gugu/utils"
)

type UserRepository interface {
	InsertUser(userId, username, email, password, bio string, profilePic []byte) error
	ListUsers() ([]user.User, error)
	VerifyCredentials(email, password string) (*user.User, error)
	GetUserById(id string) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func (r *userRepository) iterate(rows *sql.Rows) ([]user.User, error) {
	var users []user.User
	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.UserId,
			&user.Username,
			&user.Email,
			&user.Bio,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Status,
			&user.ProfilePic,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) InsertUser(userId, username, email, password, bio string, profilePic []byte) error {
	query := `INSERT INTO tb_users (
		user_id,
		username,
		email,
		password,
		bio,
		profile_pic
		) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.DB.Exec(query, userId, username, email, password, bio, profilePic)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) VerifyCredentials(email, password string) (*user.User, error) {
	var u user.User
	var hash string

	query := `SELECT 
	user_id,
	email,
	username,
	bio,
	profile_pic,
	created_at,
	updated_at,
	status,
	password FROM tb_users WHERE email = $1`

	err := r.DB.QueryRow(query, email).Scan(
		&u.UserId,
		&u.Email,
		&u.Username,
		&u.Bio,
		&u.ProfilePic,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Status,
		&hash,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	err = utils.CheckPassword(password, hash)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &u, nil
}

func (r *userRepository) GetUserById(id string) (*user.User, error) {
	var u user.User

	query := `SELECT 
	user_id,
	email,
	username,
	bio,
	profile_pic,
	created_at,
	updated_at,
	status
	FROM tb_users WHERE id = $1 and status = $2`

	err := r.DB.QueryRow(query, id, "active").Scan(
		&u.UserId,
		&u.Email,
		&u.Username,
		&u.Bio,
		&u.ProfilePic,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) GetUserByEmail(email string) (*user.User, error) {
	var u user.User

	query := `SELECT 
	user_id,
	email,
	username,
	bio,
	profile_pic,
	created_at,
	updated_at,
	status
	FROM tb_users WHERE email = $1 and status = $2`

	err := r.DB.QueryRow(query, email, "active").Scan(
		&u.UserId,
		&u.Email,
		&u.Username,
		&u.Bio,
		&u.ProfilePic,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) ListUsers() ([]user.User, error) {
	query := `SELECT 
		user_id, 
		username,
		email,
		bio,
		created_at,
		updated_at,
		status, 
		profile_pic
	FROM tb_users WHERE status = $1`

	rows, err := r.DB.Query(query, "active")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := r.iterate(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewRepository(DB *sql.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}
