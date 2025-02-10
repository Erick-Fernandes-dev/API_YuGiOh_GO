package validation

import (
	"errors"

	"github.com/Erick-Fernandes-dev/API_YuGiOh_GO/internal/domain/entity"
	"github.com/Erick-Fernandes-dev/API_YuGiOh_GO/internal/infra/strategy"
)

// Validação de erro
var (
	ErrIDIsRequired        = errors.New("ID is required")
	ErrInvalidID           = errors.New("Invalid ID")
	ErrNameIsRequired      = errors.New("Name is required")
	ErrCardTypeIsRequired  = errors.New("Card Type is required")
	ErrCardLevelIsRequired = errors.New("Card Level is required")
	ErrCardPriceIsRequired = errors.New("Card Price is required")
	ErrCardPriceInvalid    = errors.New("Card Price is invalid")
)

func (c *entity.Card) Validate() error {

	validations := []strategy.ValidationStep{
		strategy.IDValidation{},
		strategy.ErrNameIsRequired{},
		strategy.ErrCardTypeIsRequired{},
		strategy.ErrCardLevelIsRequired{},
		strategy.ErrCardPriceIsRequired{},
		strategy.ErrCardPriceInvalid{},
	}

	for _, validation := range validations {
		if err := validation.Validate(c); err != nil {
			return err
		}
	}

	return nil

}
