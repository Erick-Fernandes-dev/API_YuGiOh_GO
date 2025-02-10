package entity

import (
	"github.com/Erick-Fernandes-dev/API_YuGiOh_GO/pkg/entity"
)

type CardImage struct {
	ID            entity.ID `json:"id"`
	ImageURL      string    `json:"image_url"`
	ImageURLSmall string    `json:"image_url_small"`
}

func NewCardImage(imageURL, imageURLSmall string) (*CardImage, error) {
	return &CardImage{
		ID:            entity.NewID(),
		ImageURL:      imageURL,
		ImageURLSmall: imageURLSmall,
	}, nil
}
