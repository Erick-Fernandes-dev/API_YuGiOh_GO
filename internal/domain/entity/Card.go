package entity

import (
	"API_YuGiOh/internal/infra/strategy"
	"API_YuGiOh/pkg/entity"
)

type Card struct {
	ID         entity.ID   `json:"id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Desc       string      `json:"desc"`
	Atk        int         `json:"atk"`
	Def        int         `json:"def"`
	Level      int         `json:"level"`
	Race       string      `json:"race"`
	Atribute   string      `json:"atribute"`
	Archetype  string      `json:"archetype"`
	CardImages []CardImage `json:"card_images"`
	CardPrices []CardPrice `json:"card_prices"`
}

func NewCard(name, cardType, desc string, atk, def, level int, race, atribute, archetype string, cardImages []CardImage, cardPrices []CardPrice) (*Card, error) {
	return &Card{
		ID:         entity.NewID(),
		Name:       name,
		Type:       cardType,
		Desc:       desc,
		Atk:        atk,
		Def:        def,
		Level:      level,
		Race:       race,
		Atribute:   atribute,
		Archetype:  archetype,
		CardImages: cardImages,
		CardPrices: cardPrices,
	}, nil
}

func (c *Card) Validate() error {

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
