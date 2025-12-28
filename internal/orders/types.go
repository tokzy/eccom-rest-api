package orders

import (
	"context"

	repo "github.com/tokzy/eccom-rest-api/internal/adapters/sqlc"
)

type orderItem struct {
	productID int64 `json:"productId"`
	Quantity  int32 `json:"qunatity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
