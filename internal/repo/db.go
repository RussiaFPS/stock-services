package repo

import (
	"context"
	"fmt"
	"stock-services/internal/types"
	"stock-services/pkg/postgres"
)

type Storage interface {
	Reservation(ctx context.Context, product types.Product) error
	Relief(ctx context.Context, product types.Product) error
	Remainder(ctx context.Context, nameStock string) ([]types.Result, error)
	Reboot(ctx context.Context) error
	OutStock(ctx context.Context, nameStocks string)
}

type DbRepo struct {
	db postgres.Db
}

func NewRepository(db postgres.Db) Storage {
	return &DbRepo{db}
}

func (d *DbRepo) Remainder(ctx context.Context, nameStock string) ([]types.Result, error) {
	var availability bool

	q := `select availability from stocks where name = $1`
	if err := d.db.QueryRow(ctx, q, nameStock).Scan(&availability); err != nil {
		return nil, err
	}

	if !availability {
		return nil, fmt.Errorf("stocks %s is not availability", nameStock)
	}

	q = `update stocks set availability = false where name = $1`
	_, err := d.db.Exec(ctx, q, nameStock)
	if err != nil {
		return nil, err
	}

	defer d.OutStock(ctx, nameStock)

	q = `select productsid,amount,countreserv from balance_stocks where stocksid=$1`

	rows, err := d.db.Query(ctx, q, nameStock)
	if err != nil {
		return nil, err
	}

	result := make([]types.Result, 0)
	for rows.Next() {
		var buf types.Result

		if err = rows.Scan(&buf.ProductsId, &buf.Amount, &buf.CountReserv); err != nil {
			return nil, err
		}

		result = append(result, buf)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DbRepo) Relief(ctx context.Context, product types.Product) error {
	var nameStocks string
	var availability bool

	q := `select * from stocks limit 1`
	if err := d.db.QueryRow(ctx, q).Scan(&nameStocks, &availability); err != nil {
		return err
	}

	if !availability {
		return fmt.Errorf("stocks %s is not availability", nameStocks)
	}

	q = `update stocks set availability = false where name = $1`
	_, err := d.db.Exec(ctx, q, nameStocks)
	if err != nil {
		return err
	}

	defer d.OutStock(ctx, nameStocks)

	q = `update balance_stocks set countreserv = (countreserv-$1) where ProductsId = $2 and StocksId = $3 
                                                           and (countreserv)>=$4`
	p, err := d.db.Exec(ctx, q, product.Amount, product.Id, nameStocks, product.Amount)
	if err != nil {
		return err
	}
	if p.RowsAffected() == 0 {
		return fmt.Errorf("there is not so much in reserve")
	}

	return nil
}

func (d *DbRepo) Reservation(ctx context.Context, product types.Product) error {
	var nameStocks string
	var availability bool

	q := `select * from stocks limit 1`
	if err := d.db.QueryRow(ctx, q).Scan(&nameStocks, &availability); err != nil {
		return err
	}

	if !availability {
		return fmt.Errorf("stocks %s is not availability", nameStocks)
	}

	q = `update stocks set availability = false where name = $1`
	_, err := d.db.Exec(ctx, q, nameStocks)
	if err != nil {
		return err
	}

	defer d.OutStock(ctx, nameStocks)

	q = `update balance_stocks set countreserv = (countreserv+$1) where ProductsId = $2 and StocksId = $3 
                                                           and (Amount-countreserv)>=$4`
	p, err := d.db.Exec(ctx, q, product.Amount, product.Id, nameStocks, product.Amount)
	if err != nil {
		return err
	}
	if p.RowsAffected() == 0 {
		return fmt.Errorf("how many are out of stock")
	}

	return nil
}

func (d *DbRepo) Reboot(ctx context.Context) error {
	q := `update stocks set availability = true`
	_, err := d.db.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}

func (d *DbRepo) OutStock(ctx context.Context, nameStocks string) {
	q := `update stocks set availability = true where name = $1`
	d.db.Exec(ctx, q, nameStocks)
}
