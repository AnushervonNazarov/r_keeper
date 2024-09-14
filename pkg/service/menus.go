package service

import (
	"errors"
	"fmt"
	"r_keeper/errs"
	"r_keeper/models"
	"r_keeper/pkg/repository"
)

func GetAllMenus() (menus []models.Menu, err error) {
	if menus, err = repository.GetAllMenus(); err != nil {
		return nil, err
	}
	return menus, nil
}

func GetMenuByID(id int) (menu models.Menu, err error) {
	if menu, err = repository.GetMenuByID(id); err != nil {
		return menu, err
	}
	return menu, nil
}

func CreateMenu(menu models.Menu) error {
	_, err := repository.GetMenuByID(int(menu.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if err = repository.CreateMenu(menu); err != nil {
		return err
	}
	return nil
}

func EditMenuByID(id int, menuInput models.Menu) (*models.Menu, error) {
	_, err := repository.GetMenuByID(id)
	if err != nil {
		return nil, fmt.Errorf("menu not found: %v", err)
	}

	menuInput.ID = uint(id)

	updatedMenu, err := repository.EditMenuByID(&menuInput)
	if err != nil {
		return nil, fmt.Errorf("could not update menu: %v", err)
	}

	return updatedMenu, nil
}

func DeleteMenuByID(id int) error {
	menu, err := repository.GetMenuByID(id)
	if err != nil {
		return fmt.Errorf("menu not found: %v", err)
	}

	if err := repository.DeleteMenuByID(&menu); err != nil {
		return fmt.Errorf("could not delete menu: %v", err)
	}

	return nil
}
