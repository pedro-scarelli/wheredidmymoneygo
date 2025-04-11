package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type TransferRequest struct {
	ToUser int `json:"toUser"`
	Amount    int `json:"amount"`
}

type LoginRequest struct {
	CPF      string `json:"cpf"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type PublicUser struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(createUserRequest *CreateUserRequest) (*User, error) {
	hashedPassword, err := HashPassword(createUserRequest.Password)

	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		CPF:       createUserRequest.CPF,
		Email:     createUserRequest.Email,
		Password:  hashedPassword,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}, nil
}
