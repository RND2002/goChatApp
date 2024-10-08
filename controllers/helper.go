package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Use DefaultCost for better readability
	if err != nil {
		return "", err // Return the error if it occurs
	}
	return string(bytes), nil
}

// CompareHashedPassword compares a plain password with a hashed password
func CompareHashedPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // If err is nil, the passwords match
}
