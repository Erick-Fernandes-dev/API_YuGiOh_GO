package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {

	cardImage, err := NewCardImage("https://storage.googleapis.com/ygoprodeck.com/pics/46986414.jpg", "https://storage.googleapis.com/ygoprodeck.com/pics_small/46986414.jpg")
	if err != nil {
		t.Fatalf("failed to create card image: %v", err)
	}

	assert.Nil(t, err)
	assert.NotEmpty(t, cardImage.ID)

	assert.Equal(t, cardImage.ImageURL, "https://storage.googleapis.com/ygoprodeck.com/pics/46986414.jpg")
	assert.Equal(t, cardImage.ImageURLSmall, "https://storage.googleapis.com/ygoprodeck.com/pics_small/46986414.jpg")

}
