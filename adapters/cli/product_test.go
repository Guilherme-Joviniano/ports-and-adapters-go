package cli_test

import (
	"fmt"
	"testing"
	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/cli"
	mock_domain "github.com/Guilherme-Joviniano/go-hexagonal/mock/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "any_name"
	productId := "any_id"
	productPrice := float32(49.98)
	productStatus := "enabled"

	productMock := mock_domain.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_domain.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// Create flow
	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s ", productId, productName, productPrice, productStatus)
	result, err := cli.Run(serviceMock, "create", productName, "", productPrice)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	// Enable Flow
	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been enabled", productId, productName)
	result, err = cli.Run(serviceMock, "enable", "", productId, productPrice)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	// Disable Flow
	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been disabled", productId, productName)
	result, err = cli.Run(serviceMock, "disable", "", productId, productPrice)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	// Get Flow
	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been found with the price %f and status %s ", productId, productName, productPrice, productStatus)
	result, err = cli.Run(serviceMock, "default", "", productId, productPrice)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)
}
