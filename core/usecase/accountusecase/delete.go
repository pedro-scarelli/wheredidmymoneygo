package accountusecase

func (usecase usecase) Delete(accountID string) error {
	err := usecase.repository.Delete(accountID)
	if err != nil {
		return err
	}

	return nil
}
