package controller

import (
	"errors"
	"log"
	"net/http"
	"stock-services/internal/types"
	"unicode/utf8"
)

func (c *Controller) Reservation(r *http.Request, params *types.Params, response *types.Response) error {
	for _, v := range params.Products {
		if utf8.RuneCountInString(v.Id) < 32 || v.Amount < 1 {
			log.Printf("err reservation when Id = %s, Amount= %d", v.Id, v.Amount)
			return errors.New("invalid params")
		}
		err := c.db.Reservation(c.ctx, v)
		if err != nil {
			log.Printf("Db err reservation when Id = %s, Amount= %d", v.Id, v.Amount)
			return err
		}

		log.Printf("successfully reservation when Id = %s, Amount= %d", v.Id, v.Amount)
	}

	response.Message = "done"
	return nil
}

func (c *Controller) Relief(r *http.Request, params *types.Params, response *types.Response) error {
	for _, v := range params.Products {
		if utf8.RuneCountInString(v.Id) < 32 || v.Amount < 1 {
			log.Printf("err relief when Id = %s, Amount= %d", v.Id, v.Amount)
			return errors.New("invalid params")
		}
		err := c.db.Relief(c.ctx, v)
		if err != nil {
			log.Printf("Db err relief when Id = %s, Amount= %d", v.Id, v.Amount)
			return err
		}

		log.Printf("successfully relief when Id = %s, Amount= %d", v.Id, v.Amount)
	}

	response.Message = "done"
	return nil
}

func (c *Controller) Remainder(r *http.Request, params *types.Stock, response *types.ResponseResult) error {
	if utf8.RuneCountInString(params.Name) < 1 {
		log.Printf("err remainder when name stock = %s", params.Name)
		return errors.New("invalid params")
	}
	res, err := c.db.Remainder(c.ctx, params.Name)
	if err != nil {
		log.Printf("Db err remainder when name stock = %s", params.Name)
		return err
	}

	log.Printf("successfully remainder when stock name = %s", params.Name)
	response.Message = res
	return nil
}
