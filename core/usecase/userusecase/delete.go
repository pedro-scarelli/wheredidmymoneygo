package userusecase

func (usecase usecase) Delete(userID int) error {
	err := usecase.repository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
