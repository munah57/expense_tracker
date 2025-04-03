package utils

import "golang.org/x/crypto/bcrypt"

//HASH PASSWORD

func HashPassword(password string) (string, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	//we need to convert byte to a string because it is in bye format 
	return string(hashpass), err
}

//COMPARE PASSWORD

func ComparePassword(hashPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err !=nil {
		return err
	}
	return nil
}