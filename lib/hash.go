package lib

import (
	"fmt"
	"log"

	"golang.org/x/crypto/sha3"
)

func Sha3Hash(input string) (string, error) {
	hash := sha3.New256()
	if n, err := hash.Write([]byte(input)); err != nil {
		return "", err
	} else {
		log.Println(n)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}