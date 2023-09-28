package db

import (
	"database/sql"

	"github.com/Guilherme-Joviniano/go-hexagonal/application/domain"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDbAdapter struct {
	db *sql.DB
}

func NewProductDbAdapter(db *sql.DB) *ProductDbAdapter {
	return &ProductDbAdapter{
		db: db,
	}
}

func (p *ProductDbAdapter) Get(id string) (domain.ProductInterface, error) {
	var product domain.Product

	statement, err := p.db.Prepare("select id, name, price, status from products where id =?")

	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDbAdapter) Save(product domain.ProductInterface) (domain.ProductInterface, error) {
	var rows int

	p.db.QueryRow("select id from products where id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)

		if err != nil {
			return nil, err
		}

		return product, nil
	}

	_, err := p.update(product)

	if err != nil {
		return nil, err
	}

	return product, nil

}

func (p *ProductDbAdapter) create(product domain.ProductInterface) (domain.ProductInterface, error) {
	statement, err := p.db.Prepare("insert into products(id, name, price, status) values (?,?,?,?)")

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = statement.Close()

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDbAdapter) update(product domain.ProductInterface) (domain.ProductInterface, error) {
	_, err := p.db.Exec("update products set id = ?, name = ?, price = ?, status = ?", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}
