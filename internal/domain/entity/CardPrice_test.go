package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCardPrice(t *testing.T) {

	cp, err := NewCardPrice(1.0, 2.0, 3.0, 4.0)

	assert.Nil(t, err)
	assert.NotEmpty(t, cp)
	assert.True(t, cp.AmazonPrice == 4.0)

}
