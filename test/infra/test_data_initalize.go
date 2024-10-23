package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name, price, discount,store) 
VALUES('AirFryer',3000.0, 22.0, 'ABC TECH'),
('Computer',1500.0, 10.0, 'ABC TECH'),
('Dishwasher',10000.0, 15.0, 'ABC TECH'),
('Lambader',2000.0, 0.0, 'Decoration Store');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductsResult, insertProductsErr := dbPool.Exec(ctx, INSERT_PRODUCTS)
	
	if insertProductsErr != nil {
		log.Error("Error inserting products: ", insertProductsErr)
		return
	}
	log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected()))
}