package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (*Service) List() []Product {
	return allProduct
}
