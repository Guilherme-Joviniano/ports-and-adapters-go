package factories

import (
	dbSQL "github.com/Guilherme-Joviniano/go-hexagonal/adapters/db/sql"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/infra/db/sqlite/connections"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/service"
)

func MakeProductService() *service.ProductService {
	db := connections.GetSqliteInstance()
	productAdapter := dbSQL.NewProductDbAdapter(db.Connection)
	return service.NewProductService(productAdapter)
}
