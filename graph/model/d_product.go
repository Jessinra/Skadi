package model

import (
	"strings"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
)

func NewProduct(in *domain.Product) *Product {
	return &Product{
		ID:          in.ID,
		CreatedAt:   in.CreatedAt,
		Name:        in.Name,
		Description: in.Description,
		ImagesURLs:  strings.Split(in.ImagesURLs, ","),
		Weight:      in.Weight,
		Dimensions:  in.Dimensions,
		Categories:  strings.Split(in.Categories, ","),
		Locations:   NewProductLocations(in.Locations),
		Prices:      NewProductPrices(in.Prices),
	}
}

func NewProducts(in []domain.Product) []Product {
	out := []Product{}
	for i := range in {
		out = append(out, *NewProduct(&in[i]))
	}

	return out
}

func NewProductLocation(in *domain.ProductLocation) *ProductLocation {
	return &ProductLocation{
		ID:        in.ID,
		CreatedAt: in.CreatedAt,
		UserID:    in.UserID,
		ProductID: in.ProductID,
		Text:      in.Text,
		Country:   in.Country,
		Province:  in.Province,
		City:      in.City,
		Area:      in.Area,
		Street:    in.Street,
		Building:  in.Building,
		Store:     in.Store,
		Longitude: in.Longitude,
		Latitude:  in.Latitude,
	}
}

func NewProductLocations(in []domain.ProductLocation) []ProductLocation {
	out := []ProductLocation{}
	for i := range in {
		out = append(out, *NewProductLocation(&in[i]))
	}

	return out
}

func NewProductPrice(in *domain.ProductPrice) *ProductPrice {
	return &ProductPrice{
		ID:               in.ID,
		CreatedAt:        in.CreatedAt,
		UserID:           in.UserID,
		ProductID:        in.ProductID,
		Currency:         in.Currency,
		Price:            in.Price,
		IsPriceEstimated: in.IsPriceEstimated,
	}
}

func NewProductPrices(in []domain.ProductPrice) []ProductPrice {
	out := []ProductPrice{}
	for i := range in {
		out = append(out, *NewProductPrice(&in[i]))
	}

	return out
}
