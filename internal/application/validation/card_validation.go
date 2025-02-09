package validation

import "errors"

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
