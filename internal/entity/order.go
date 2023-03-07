package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o *Order) Validate() error {

	if o.ID == "" {
		return errors.New("Id is required")
	}

	if o.Price <= 0 {
		return errors.New("Invalid Price")
	}

	if o.Tax <= 0 {
		return errors.New("Invalid Tax")
	}

	return nil
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {

	o := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := o.Validate()
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Order) TotalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}
