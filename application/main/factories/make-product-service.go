package factories

import (
	"database/sql"
	"sync"

	dbSQL "github.com/Guilherme-Joviniano/go-hexagonal/adapters/db/sql"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/infra/db/sqlite/connections"
	"github.com/Guilherme-Joviniano/go-hexagonal/application/service"
)

var wg sync.WaitGroup

func MakeProductService() *service.ProductService {
	dbChannel := make(chan *sql.DB)
	wg.Add(1)
	go connections.GetSqliteInstance(dbChannel)
	productAdapter := dbSQL.NewProductDbAdapter(<-dbChannel)
	wg.Wait()
	close(dbChannel)
	return service.NewProductService(productAdapter)
}
