package service

import (
	"r_keeper/models"
	"r_keeper/pkg/repository"
	"r_keeper/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return "", err
	}

	accessToken, err = GenerateToken(uint(user.ID), user.Username)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
