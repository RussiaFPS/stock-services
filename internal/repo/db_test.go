package repo

import (
	"context"
	"stock-services/internal/types"
	"stock-services/pkg/postgres"
	"testing"
)

func Test_db_Reservation(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "stock-services")
	t.Setenv("pst_port", "5430")

	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name    string
		product types.Product
	}{
		{
			name:    "test1",
			product: types.Product{Id: "4795435a-a814-4203-8dc6-4bf093bec52e", Amount: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			err := d.Reservation(context.Background(), tt.product)
			if err != nil {
				t.Error("Reservation error")
				return
			}
		})
	}
}

func Test_db_Relief(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "stock-services")
	t.Setenv("pst_port", "5430")

	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name    string
		product types.Product
	}{
		{
			name:    "test1",
			product: types.Product{Id: "4795435a-a814-4203-8dc6-4bf093bec52e", Amount: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			err := d.Relief(context.Background(), tt.product)
			if err != nil {
				t.Error("Relief error")
				return
			}
		})
	}
}

func Test_db_Remainder(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "stock-services")
	t.Setenv("pst_port", "5430")

	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name      string
		nameStock string
	}{
		{
			name:      "test1",
			nameStock: "t1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			res, err := d.Remainder(context.Background(), tt.nameStock)
			if err != nil {
				t.Error("Relief error: ", res)
				return
			}
		})
	}
}
