package userrepository

import (
	"context"
)

func (repository repository) Delete(userID int) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		`DELETE FROM tb_user
		WHERE pt_it_id = $1`,
		userID)
	if err != nil {
		return err
	}

	return nil
}
