package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func HashPassword(password string) (string, error) {
	cost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		cost = bcrypt.DefaultCost
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("falha ao gerar hash da senha: %w", err)
	}

	return string(bytes), nil
}

func IsPasswordIncorrect(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
