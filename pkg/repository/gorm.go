package repository

import (
	"errors"
	"r_keeper/errs"
)

func translateError(err error) error {
	if errors.Is(err, errs.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}
