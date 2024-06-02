package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func StringToUint(str string, ok bool) (uint, bool) {
	i, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, false
	}

	return uint(i), true
}

func Divide(a, b int) float64 {
	if b == 0 {
		return 0
	}

	return float64(a) / float64(b)
}

func GetBearerToken(authHeader string) string {
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return ""
	}

	return authHeaderParts[1]
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyHashedPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("incorrect password, %s, hashed: %s, payload: %s", err.Error(), hashedPassword, password)
	}

	return nil
}

func ToStringSLice[T any](input []T) []string {
	var s []string

	for _, v := range input {
		s = append(s, fmt.Sprintf("%v", v))
	}

	return s
}

func GetPageCount(total, pageSize int) int {
	if pageSize == 0 {
		return 0
	}

	return int(math.Ceil(Divide(total, pageSize)))
}

func DollarsFromCents(cents int) float64 {
	return float64(cents) / 100.0
}
