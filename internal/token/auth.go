package token

import "golang.org/x/crypto/bcrypt"

func ComparePasswordWithHash(input, hash string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	if err != nil {
		return err
	}
	return
}
