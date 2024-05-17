package userRepository

import (
	"database/sql"
	user "gugu/interfaces/user"
)

type UserRepository interface {
	InsertUser(userId, username, email, password, bio string, profilePic []byte) error
	ListUsers() ([]user.User, error)
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
	_, err := r.DB.Exec(
		`INSERT INTO tb_users (
			user_id,
			username,
			email,
			password,
			bio,
			profile_pic
			) 
			VALUES ($1, $2, $3, $4, $5, $6)`, userId, username, email, password, bio, profilePic)
	if err != nil {
		return err
	}
	return nil
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
	FROM tb_users`

	rows, err := r.DB.Query(query)
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
