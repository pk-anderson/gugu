package user

import (
	"database/sql"
	"errors"
	"fmt"
	user "gugu/interfaces/user"
	"gugu/repositories/userRepository"
	"gugu/utils"
	"regexp"
	"strings"
)

type UserService interface {
	CreateUser(username, email, password, bio string, profilePic []byte) (string, error)
	ListUsers() ([]user.User, error)
}

type userService struct {
	DB *sql.DB
}

func validations(username, email, password string) []string {
	var errs []string

	if len(username) > 4 {
		errs = append(errs, "Username must have a maximum of 4 characters")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailPattern, email); !matched {
		errs = append(errs, "Invalid email")
	}

	if len(password) > 8 {
		errs = append(errs, "Password must have at least 8 characters")
	}

	return errs
}

func (s *userService) CreateUser(username, email, password, bio string, profilePic []byte) (string, error) {
	errs := validations(username, email, password)
	if len(errs) > 0 {
		return "", fmt.Errorf(strings.Join(errs, "; "))
	}

	rep := userRepository.NewRepository(s.DB)
	user, err := rep.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user != nil {
		return "", errors.New("user with this email already exists")
	}

	userId := utils.GenerateUUID()
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	err = rep.InsertUser(userId, username, email, hashPassword, bio, profilePic)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *userService) ListUsers() ([]user.User, error) {
	rep := userRepository.NewRepository(s.DB)
	users, err := rep.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewService(DB *sql.DB) UserService {
	return &userService{
		DB: DB,
	}
}
