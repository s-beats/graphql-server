package usecase

import "errors"

func errorFailedUpdateData() error {
	return errors.New("failed update data")
}
