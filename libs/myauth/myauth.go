package myauth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hasPassword(password string) (string, error) {
	// Mã hóa mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}
func checkPasswordHash(password, hash string) bool {
	// Kiểm tra mật khẩu
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
