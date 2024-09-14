package repository

import (
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllMenus() (menus []models.Menu, err error) {
	if err = db.GetDBConn().Preload("Items").Find(&menus).Error; err != nil {
		logger.Error.Println("[repository.GetAllMenus] error getting all menus:", err.Error())
		return nil, translateError(err)
	}
	return menus, nil
}

func GetMenuByID(id int) (menu models.Menu, err error) {

	if err = db.GetDBConn().Where("id = ?", id).First(&menu).Error; err != nil {
		logger.Error.Println("[repository.GetAllMenus] error getting menu by id. Error is:", err.Error())
		return menu, translateError(err)
	}
	return menu, nil
}

func CreateMenu(menu models.Menu) (err error) {
	if err := db.GetDBConn().Create(&menu).Error; err != nil {
		logger.Error.Println("[repository.CreateMenu] error creating menu. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func EditMenuByID(menu *models.Menu) (*models.Menu, error) {
	if err := db.GetDBConn().Save(&menu).Error; err != nil {
		logger.Error.Println("[repository.EditMenuByID] error editing menu. Error is:", err.Error())
		return nil, translateError(err)
	}
	return menu, nil
}

func DeleteMenuByID(menu *models.Menu) error {
	if err := db.GetDBConn().Delete(menu).Error; err != nil {
		logger.Error.Println("[repository.DeleteMenuByID] error deleating menu. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}
