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
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type LoginRequest struct {
	CPF      string `json:"cpf"`
	Password string `json:"password"`
}

type Account struct {
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

type PublicAccount struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(createAccountRequest *CreateAccountRequest) (*Account, error) {
	hashedPassword, err := HashPassword(createAccountRequest.Password)

	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName: createAccountRequest.FirstName,
		LastName:  createAccountRequest.LastName,
		CPF:       createAccountRequest.CPF,
		Email:     createAccountRequest.Email,
		Password:  hashedPassword,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}, nil
}
