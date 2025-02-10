package strategy

import (
	"github.com/Erick-Fernandes-dev/API_YuGiOh_GO/internal/application/validation"
	C "github.com/Erick-Fernandes-dev/API_YuGiOh_GO/internal/domain/entity"
	pkg "github.com/Erick-Fernandes-dev/API_YuGiOh_GO/pkg/entity"
)

type ValidationStep interface {
	Validate(c *C.Card) error
}

// Structs de validação
type IDValidation struct{}
type ErrNameIsRequired struct{}
type ErrCardTypeIsRequired struct{}
type ErrCardLevelIsRequired struct{}
type ErrCardPriceIsRequired struct{}
type ErrCardPriceInvalid struct{}

// Implementações concretas

func (v IDValidation) Validate(c *C.Card) error {

	if c.ID.String() == "" {
		return validation.ErrIDIsRequired
	}

	if _, err := pkg.ParseID(c.ID.String()); err != nil {
		return validation.ErrInvalidID
	}

	return nil
}

func (v ErrNameIsRequired) Validate(c *C.Card) error {
	if c.Name == "" {
		return validation.ErrNameIsRequired
	}
	return nil
}

func (v ErrCardTypeIsRequired) Validate(c *C.Card) error {
	if c.Type == "" {
		return validation.ErrCardTypeIsRequired
	}
	return nil
}

func (v ErrCardLevelIsRequired) Validate(c *C.Card) error {
	if c.Level == 0 {
		return validation.ErrCardLevelIsRequired
	}
	return nil
}

func (v ErrCardPriceIsRequired) Validate(c *C.Card) error {
	if c.CardPrices == nil {
		return validation.ErrCardPriceIsRequired
	}
	return nil
}

func (v ErrCardPriceInvalid) Validate(c *C.Card) error {
	for _, cp := range c.CardPrices {
		if cp.TcgplayerPrice == 0 {
			return validation.ErrCardPriceInvalid
		}
	}
	return nil
}
