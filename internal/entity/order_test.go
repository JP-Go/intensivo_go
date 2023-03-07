package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{
		ID:    "",
		Price: 12,
		Tax:   0.7,
	}

	assert.Error(t, order.Validate(), "invalid id")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{
		ID:    "",
		Price: 0,
		Tax:   0.7,
	}

	assert.Error(t, order.Validate(), "invalid Price")

}
func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{
		ID:    "",
		Price: 12,
		Tax:   0.0,
	}

	assert.Error(t, order.Validate(), "invalid Tax")

}

func TestWithAllValidParams(t *testing.T) {
	order := Order{
		ID:    "18927",
		Price: 12,
		Tax:   10,
	}

	assert.NoError(t, order.Validate())
	assert.Equal(t, 12.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
	assert.Equal(t, "18927", order.ID)

	order.TotalPrice()
	assert.Equal(t, 22.0, order.FinalPrice)

}
