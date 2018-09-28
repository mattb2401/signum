package signum

import "golang.org/x/crypto/bcrypt"

func MakeHash(password string, rounds int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Attempt(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetRounds(hash string) (int, error) {
	pass := []byte(hash)
	cost, err := bcrypt.Cost(pass)
	if err != nil {
		return 0, err
	}
	return cost, nil
}
