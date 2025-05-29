package accountrepository

import (
	"context"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (repository repository) Delete(accountID string) error {
	ctx := context.Background()

	result, err := repository.db.Exec(
		ctx,
		`DELETE FROM tb_account
		WHERE pk_st_id = $1`,
		accountID,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrAccountNotFound
	}

	return nil
}