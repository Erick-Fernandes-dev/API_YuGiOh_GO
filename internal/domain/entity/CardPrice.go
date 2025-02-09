package entity

type CardPrice struct {
	CardmarketPrice float64 `json:"cardmarket_price"`
	TcgplayerPrice  float64 `json:"tcgplayer_price"`
	EbayPrice       float64 `json:"ebay_price"`
	AmazonPrice     float64 `json:"amazon_price"`
}

func NewCardPrice(cardmarketPrice, tcgplayerPrice, ebayPrice, amazonPrice float64) (*CardPrice, error) {

	return &CardPrice{
		CardmarketPrice: cardmarketPrice,
		TcgplayerPrice:  tcgplayerPrice,
		EbayPrice:       ebayPrice,
		AmazonPrice:     amazonPrice,
	}, nil
}
