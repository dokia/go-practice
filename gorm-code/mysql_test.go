package gorm_code

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	err := Insert()

	assert.NoError(t, err)
}

func TestSingleQuery(t *testing.T) {
	product, err := SingleQuery()

	if assert.NoError(t, err) {
		assert.NotEmpty(t, product)
	}
}

func TestMultiQuery(t *testing.T) {
	products, err := MultiQuery()

	if assert.NoError(t, err) {
		assert.NotEmpty(t, products)
	}
}