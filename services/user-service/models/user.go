package models

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

// Validate validates the User data.
func (u *User) Validate() error {
	if err := u.validateUsername(); err != nil {
		return fmt.Errorf("invalid username: %w", err)
	}
	if err := u.validateEmail(); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}
	if err := u.validatePassword(); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

// validateUsername validates the username against the specified rules.
func (u *User) validateUsername() error {
	if len(u.Username) < 6 || len(u.Username) > 20 {
		return fmt.Errorf("username must be between 6 and 20 characters long")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9._-]{6,20}$`).MatchString(u.Username) {
		return fmt.Errorf("username contains invalid characters")
	}
	return nil
}

// validateEmail validates the email address.
func (u *User) validateEmail() error {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}


// validatePassword validates the password against the specified rules.
func (u *User) validatePassword() error {
	if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

// HashPassword hashes the password using bcrypt.
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a given password with the hash.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

