package application_test

import (
	"testing"

	"github.com/Guilherme-Joviniano/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestApplication_Product_Enable(t *testing.T) {
	product := application.NewProduct("Cup of Coffe", 10.99)
	err := product.Enabled()
	require.Nil(t, err)
	product.Price = 0
	err = product.Enabled()
	require.Equal(t, "Product price must be greater than zero", err.Error())
}

func TestApplication_Product_Disable(t *testing.T) {
	product := application.NewProduct("Cup of Coffe", 0)
	err := product.Disable()
	require.Nil(t, err)
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Product price must be zero to be disabled", err.Error())
}

func TestApplication_Product_IsValid(t *testing.T) {
	product := application.NewProduct("Cup of Coffe", 10)
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "ANY_INVALID_STATUS"
	_, err = product.IsValid()
	product.Status = application.DISABLED
	require.Equal(t, "the status must be setted as the constants in the package", err.Error())

	product.Price = -10
	_, err = product.IsValid()
	product.Price = 10
	require.Equal(t, "the price must be greater than zero", err.Error())

	product.ID = "INVALID_UUID"
}
