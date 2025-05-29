package accountrepository

import (
	"context"
	"fmt"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (repository repository) Update(updateAccountRequestDto *dto.UpdateAccountRequestDTO) error {
	ctx := context.Background()
	query := `
		UPDATE tb_account
		SET st_first_name = $1,
		    st_last_name  = $2,
		    st_password   = $3
		WHERE pk_st_id = $4;
	    `

	_, err := repository.db.Exec(
		ctx,
		query,
		updateAccountRequestDto.FirstName,
		updateAccountRequestDto.LastName,
		updateAccountRequestDto.Password,
		updateAccountRequestDto.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar conta no banco: %w", err)
	}

	return nil
}
