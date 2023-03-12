package hash

import "golang.org/x/crypto/bcrypt"

func PswdHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func PswdVerif(hashpswd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpswd), []byte(password))
}