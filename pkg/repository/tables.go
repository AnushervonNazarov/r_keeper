package repository

import (
	"r_keeper/db"
	"r_keeper/logger"
	"r_keeper/models"
)

func GetAllTables() (tables []models.Table, err error) {
	if err = db.GetDBConn().Find(&tables).Error; err != nil {
		logger.Error.Println("[repository.GetAllTables] error getting all tables:", err.Error())
		return nil, translateError(err)
	}
	return tables, nil
}

func GetTableByID(id int) (table models.Table, err error) {

	if err = db.GetDBConn().Where("id = ?", id).First(&table).Error; err != nil {
		logger.Error.Println("[repository.GetTableByID] error getting table by id. Error is:", err.Error())
		return table, translateError(err)
	}
	return table, nil
}

func CreateTable(table models.Table) (err error) {
	if err := db.GetDBConn().Create(&table).Error; err != nil {
		logger.Error.Println("[repository.CreateTable] error creating table. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func EditTableByID(table *models.Table) (*models.Table, error) {
	if err := db.GetDBConn().Save(&table).Error; err != nil {
		logger.Error.Println("[repository.EditTableByID] error editing table. Error is:", err.Error())
		return nil, translateError(err)
	}
	return table, nil
}

func DeleteTableByID(table *models.Table) error {
	if err := db.GetDBConn().Delete(table).Error; err != nil {
		logger.Error.Println("[repository.DeleteTableByID] error deleating table. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}
