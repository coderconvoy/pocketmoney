package main

import (
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

type Password struct {
	hash, salt []byte
}

func NewPassword(pw string) (Password, error) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		return Password{}, err
	}
	dk, err := scrypt.Key([]byte(pw), salt, 16384, 8, 1, 32)
	return Password{dk, salt}, nil

}
func (p Password) Check(pw string) bool {
	dk, err := scrypt.Key([]byte(pw), p.salt, 16384, 8, 1, 32)
	if err != nil {
		return false
	}
	return string(dk) == string(p.hash)
}

type Family struct {
	Email      string
	FamilyName string
	Password   Password

	Members  []User
	Accounts []Account
}

type User struct {
	Username string
	Password string
	Parent   bool
}

type Login struct {
	Username string
	Family   string
}

type Transaction struct {
	ID                 uint64
	FromUser, DestUser string
	FromID, DestID     int
	Amount             int
	Authorized         bool
	Purpose            string
}

type Account struct {
	Username string
	ID       uint64
	Name     string
	Monthly  []int
}
