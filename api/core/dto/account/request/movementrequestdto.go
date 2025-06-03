package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	enum "github.com/pedro-scarelli/wheredidmymoneygo/core/domain/enum"
	"io"
	"time"
)

type MovementRequestDTO struct {
	AccountID   string            `json:"-"`
	Value       float64           `json:"value"`
	Type        enum.MovementType `json:"type"`
	Recurrence  int               `json:"recurrence"`
	DueDate     time.Time         `json:"dueDate"`
	Description string            `json:"description"`
}

func (d *MovementRequestDTO) Validate() error {
	switch d.Type {
	case enum.DEBITO, enum.CREDITO:
	default:
		return fmt.Errorf("tipo inválido: %s. Valores permitidos: %s ou %s",
			d.Type,
			enum.DEBITO,
			enum.CREDITO)
	}

	if d.Recurrence < 1 || d.Recurrence > 12 {
		return errors.New("ocorrência deve ser entre 1 e 12")
	}

	return nil
}

func FromJSONCreateMovementRequestDTO(body io.Reader) (*MovementRequestDTO, error) {
	createMovementRequestDTO := MovementRequestDTO{}
	if err := json.NewDecoder(body).Decode(&createMovementRequestDTO); err != nil {
		return nil, err
	}

	if err := createMovementRequestDTO.Validate(); err != nil {
		return nil, err
	}

	return &createMovementRequestDTO, nil
}
