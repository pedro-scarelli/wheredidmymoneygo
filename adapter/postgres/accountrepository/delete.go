package accountrepository

import (
	"context"
)

func (repository repository) Delete(accountID string) error {
	ctx := context.Background()

	_, err := repository.db.Exec(
		ctx,
		`DELETE FROM tb_account
		WHERE pk_st_id = $1`,
		accountID)
	if err != nil {
		return err
	}

	return nil
}
