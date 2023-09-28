package service_test

import (
	"testing"

	"github.com/Guilherme-Joviniano/go-hexagonal/application/service"
	mock_domain "github.com/Guilherme-Joviniano/go-hexagonal/mock/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestApplicationProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_domain.NewMockProductInterface(ctrl)
	persistence := mock_domain.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("any_id")

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestApplicationProductService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	product := mock_domain.NewMockProductInterface(ctrl)

	persistence := mock_domain.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)

	require.Nil(t, err)
	require.Equal(t, result, product)
}

func TestApplicationProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	product := mock_domain.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_domain.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)

	require.Nil(t, err)
	require.Equal(t, result, product)
}

func TestApplicationProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	product := mock_domain.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_domain.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := service.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)

	require.Nil(t, err)
	require.Equal(t, result, product)
}
