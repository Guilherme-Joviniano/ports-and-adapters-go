package cli

import (
	"fmt"

	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
)

func Run(service domain.ProductServiceInterface, action, productName, productId string, price float32) (string, error) {
	var result string = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s ", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been enabled", res.GetID(), res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been disabled", res.GetID(), res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been found with the price %f and status %s ", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
