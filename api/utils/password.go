package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Gerar o hash da senha usando bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Retornar o hash como uma string
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	// Comparar a senha fornecida com o hash armazenado
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
