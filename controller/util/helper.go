package util

import "rest-playground/service/dto"

//TODO: implement
func ToModel(addProductRequest dto.ProductCreate) dto.ProductCreate {
	return dto.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}