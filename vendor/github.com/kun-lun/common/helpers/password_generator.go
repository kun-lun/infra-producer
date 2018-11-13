package helpers

import (
	"crypto/rand"
	"math/big"
)

type PasswordGenerator struct{}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

const DefaultPasswordLength = 20

type passwordParams struct {
	Length int `yaml:"length"`
}

var supportedPasswordParams = []string{
	"length",
}

func NewPasswordGenerator() PasswordGenerator {
	return PasswordGenerator{}
}

func (PasswordGenerator) Generate(length int) (string, error) {
	if length == 0 {
		length = DefaultPasswordLength
	}

	lengthLetterRunes := big.NewInt(int64(len(letterRunes)))
	passwordRunes := make([]rune, length)

	for i := range passwordRunes {
		index, err := rand.Int(rand.Reader, lengthLetterRunes)
		if err != nil {
			return "", err
		}

		passwordRunes[i] = letterRunes[index.Int64()]
	}

	return string(passwordRunes), nil
}
