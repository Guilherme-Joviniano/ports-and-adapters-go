package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Guilherme-Joviniano/go-hexagonal/adapters/db"
	"github.com/Guilherme-Joviniano/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
		"id" string,
		"name" string,
		"status" string,
		"price" float
		);`
	statement, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func seedProduct(db *sql.DB) {
	insert := `insert into products values("abc", "test", "disabled", 0)`
	statement, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
}

func TestAdapterDbProductDbAdapter_Get(t *testing.T) {
	setUp()
	createTable(Db)
	seedProduct(Db)
	defer Db.Close()

	productDbAdapter := db.NewProductDbAdapter(Db)
	result, err := productDbAdapter.Get("abc")

	require.Nil(t, err)

	require.Equal(t, "test", result.GetName())
	require.Equal(t, "abc", result.GetID())
	require.Equal(t, "disabled", result.GetStatus())
	require.Equal(t, float32(0), result.GetPrice())
}

func TestAdapterDbProductDbAdapter_Save(t *testing.T) {
	setUp()
	createTable(Db)
	seedProduct(Db)
	defer Db.Close()

	productDbAdapter := db.NewProductDbAdapter(Db)
	product := application.NewProduct("test", 10)
	result, err := productDbAdapter.Save(product)
	
	require.Nil(t, err)

	require.Equal(t, "test", result.GetName())
	require.Equal(t, "disabled", result.GetStatus())
	require.Equal(t, float32(10), result.GetPrice())
}
