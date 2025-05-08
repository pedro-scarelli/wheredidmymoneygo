package accountrepository

import (
	"context"
)

func (repository repository) Delete(accountID int) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		`DELETE FROM tb_account
		WHERE pt_it_id = $1`,
		accountID)
	if err != nil {
		return err
	}

	return nil
}
