package service

import (
	"errors"
	"fmt"
	"r_keeper/errs"
	"r_keeper/models"
	"r_keeper/pkg/repository"
)

func GetAllTables() (tables []models.Table, err error) {
	if tables, err = repository.GetAllTables(); err != nil {
		return nil, err
	}
	return tables, nil
}

func GetTableByID(id int) (table models.Table, err error) {
	if table, err = repository.GetTableByID(id); err != nil {
		return table, err
	}
	return table, nil
}

func CreateTable(table models.Table) error {
	_, err := repository.GetTableByID(int(table.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if err = repository.CreateTable(table); err != nil {
		return err
	}
	return nil
}

func EditTableByID(id int, tableInput models.Table) (*models.Table, error) {
	_, err := repository.GetTableByID(id)
	if err != nil {
		return nil, fmt.Errorf("table not found: %v", err)
	}

	tableInput.ID = id

	updatedtable, err := repository.EditTableByID(&tableInput)
	if err != nil {
		return nil, fmt.Errorf("could not update table: %v", err)
	}

	return updatedtable, nil
}

func DeleteTableByID(id int) error {
	table, err := repository.GetTableByID(id)
	if err != nil {
		return fmt.Errorf("table not found: %v", err)
	}

	if err := repository.DeleteTableByID(&table); err != nil {
		return fmt.Errorf("could not delete table: %v", err)
	}

	return nil
}
