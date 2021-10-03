package services

var productService *ProductService

func InitServices() error {
	svc, err := NewProductService()
	if err != nil {
		return err
	}

	productService = svc
	return nil
}

func InitServiceDependencies() error {
	return productService.InitDependencies()
}

func ValidateServices() error {
	return productService.Validate()
}

func GetProductService() *ProductService {
	return productService
}
