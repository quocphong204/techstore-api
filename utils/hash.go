package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword mã hóa mật khẩu trước khi lưu vào DB
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
