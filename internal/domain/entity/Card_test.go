package entity

import (
	"testing"

	"github.com/Erick-Fernandes-dev/API_YuGiOh_GO/internal/application/validation"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	t.Run("should create valid card with correct data", func(t *testing.T) {
		// Setup
		price := createTestCardPrice(t)
		image := createTestCardImage(t)

		// Execution
		card, err := NewCard(
			"Dark Magician",
			"Monster",
			"The ultimate wizard in terms of attack and defense.",
			2500,
			2100,
			7,
			"Spellcaster",
			"Dark",
			"Dark Magician",
			[]CardImage{*image},
			[]CardPrice{*price},
		)

		// Validation
		assert.NoError(t, err)
		assertValidUUID(t, card.ID.String())
		assert.Equal(t, "Dark Magician", card.Name)
		assert.Equal(t, "Monster", card.Type)
		assert.Equal(t, 2500, card.Atk)
		assert.Equal(t, 2100, card.Def)
		assert.Equal(t, 10.00, card.CardPrices[0].TcgplayerPrice)
	})

	t.Run("should return error when required fields are missing", func(t *testing.T) {
		testCases := []struct {
			name          string
			missingField  string
			expectedError error
		}{
			{
				name:          "empty name",
				missingField:  "name",
				expectedError: validation.ErrNameIsRequired,
			},
			{
				name:          "invalid type",
				missingField:  "type",
				expectedError: validation.ErrCardTypeIsRequired,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Setup
				invalidCard := createInvalidCard(tc.missingField)

				// Execution
				_, err := NewCard(
					invalidCard.Name,
					invalidCard.Type,
					invalidCard.Desc,
					invalidCard.Atk,
					invalidCard.Def,
					invalidCard.Level,
					invalidCard.Race,
					invalidCard.Atribute,
					invalidCard.Archetype,
					invalidCard.CardImages,
					invalidCard.CardPrices,
				)

				// Validation
				assert.ErrorIs(t, err, tc.expectedError)
			})
		}
	})
}

// Helper functions
func createTestCardPrice(t *testing.T) *CardPrice {
	price, err := NewCardPrice(5.00, 10.00, 15.00, 20.00)
	assert.NoError(t, err)
	return price
}

func createTestCardImage(t *testing.T) *CardImage {
	image, err := NewCardImage(
		"https://storage.googleapis.com/ygoprodeck.com/pics/46986414.jpg",
		"https://storage.googleapis.com/ygoprodeck.com/pics_small/46986414.jpg",
	)
	assert.NoError(t, err)
	return image
}

func assertValidUUID(t *testing.T, id string) {
	_, err := uuid.Parse(id)
	assert.NoError(t, err, "ID should be a valid UUID")
}

func createInvalidCard(missingField string) *Card {
	validCard := &Card{
		Name:       "Dark Magician",
		Type:       "Monster",
		Desc:       "Description",
		Atk:        2500,
		Def:        2100,
		Level:      7,
		Race:       "Spellcaster",
		Atribute:   "Dark",
		Archetype:  "Dark Magician",
		CardImages: []CardImage{{}},
		CardPrices: []CardPrice{{}},
	}

	switch missingField {
	case "name":
		validCard.Name = ""
	case "type":
		validCard.Type = "InvalidType"
	}

	return validCard
}
