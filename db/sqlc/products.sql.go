// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: products.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  title, 
  description, 
  price
) VALUES (
  $1, $2, $3
)
RETURNING id, title, description, price, created_at
`

type CreateProductParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct, arg.Title, arg.Description, arg.Price)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, title, description, price, created_at FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT id, title, description, price, created_at FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductForUpdate, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, title, description, price, created_at FROM products
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
set title = $2,
description = $3,
price = $4
WHERE id = $1
RETURNING id, title, description, price, created_at
`

type UpdateProductParams struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}
