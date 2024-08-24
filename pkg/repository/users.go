package repository

import (
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, err
	}

	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return user, err
	}

	return user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return err
	}

	return nil
}

func EditUserByID(user *models.User) (*models.User, error) {
	if err := db.GetDBConn().Omit("password").Save(&user).Error; err != nil {
		logger.Error.Printf("[repository.EditUserByID] error editing user: %v\n", err)
		return nil, err
	}
	return user, nil
}

func DeleteUserByID(user *models.User) error {
	if err := db.GetDBConn().Delete(user).Error; err != nil {
		logger.Error.Printf("[repository.DeleteUserByID] error deleating user: %v\n", err)
		return err
	}
	return nil
}
