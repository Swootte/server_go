package utils

import (
	"fmt"
	"os"
	"server/graph/model"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPincode(user model.User, pinCode string) (*model.User, error) {
	err := bcrypt.CompareHashAndPassword([]byte(*user.PinCode), []byte(pinCode))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	decryped, err := Decrypt(user.Keypair.SecretKey, os.Getenv("SERVER_SECRET"))
	if err != nil {
		return nil, err
	}
	user.Keypair.SecretKey = decryped
	return &user, nil
}

func VerifyPincodeWithEnterprise(enterprise model.Enterprise, user model.User, pinCode string) (*model.Enterprise, error) {
	err := bcrypt.CompareHashAndPassword([]byte(*user.PinCode), []byte(pinCode))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}
	decryped, err := Decrypt(*enterprise.WalletSecretKey, os.Getenv("SERVER_SECRET"))
	if err != nil {
		return nil, err
	}
	enterprise.WalletSecretKey = &decryped

	return &enterprise, nil

}
