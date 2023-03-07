package database

import (
	"database/sql"

	"github.com/JP-Go/go-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(o *entity.Order) error {
	_, err := r.Db.Exec("Insert Into orders (id,price,tax,final_price) Values(?,?,?,?)",
		o.ID, o.Price, o.Tax, o.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (float64, error) {
	var total int

	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return float64(total), nil
}
