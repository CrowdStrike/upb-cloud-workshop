package sql

import (
	"context"
	"database/sql"
	exam_api_domain "exam-api/domain"
	"fmt"

	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	sqlCreateStmt = `INSERT INTO products (id, name, manufacturer, price, stock, tags)
					VALUES ($1, $2, $3, $4, $5, $6) 
					RETURNING id, name, manufacturer, price, stock, tags`

	sqlGetByIDStmts = `SELECT id, name, manufacturer, price, stock, tags
					FROM products 
					WHERE id = ANY($1)`
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	mr := ProductRepository{
		db: db,
	}

	return &mr
}

func (p *ProductRepository) Save(product exam_api_domain.Product) (string, bool, error) {
	ctx := context.Background()
	id := product.GetHash()
	row := p.db.QueryRowContext(
		ctx,
		sqlCreateStmt,
		[]byte(id),
		[]byte(product.Name),
		[]byte(product.Manufacturer),
		product.Price,
		product.Stock,
		pq.Array(product.Tags))
	if row.Err() != nil {
		return "", false, row.Err()
	}

	return id, false, row.Err()

}

func (p *ProductRepository) Get(id string) (exam_api_domain.Product, bool, error) {
	ctx := context.Background()

	rows, err := p.db.QueryContext(ctx, sqlGetByIDStmts, pq.Array(id))
	if err != nil {
		return exam_api_domain.Product{}, false, err
	}
	if rows.Err() != nil {
		return exam_api_domain.Product{}, false, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Errorf("Failed to read rows with err=%v", err)
		}
	}()

	var products []exam_api_domain.Product
	for rows.Next() {
		var prodID string
		product := exam_api_domain.Product{}

		if err := rows.Scan(&prodID, &product.Name, &product.Manufacturer, &product.Price, &product.Stock, pq.Array(&product.Tags)); err != nil {
			return exam_api_domain.Product{}, false, err
		}

		products = append(products, product)
	}

	if len(products) != 1 {
		// no rows selected
		return exam_api_domain.Product{}, false, fmt.Errorf("no items")
	}

	return products[0], true, err
}

func (p *ProductRepository) Update(id string, diff exam_api_domain.Product) (bool, error) {
	return false, nil
}

func (p *ProductRepository) Delete(id string) (bool, error) {
	return false, nil
}
