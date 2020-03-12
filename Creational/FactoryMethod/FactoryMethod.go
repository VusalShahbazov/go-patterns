package FactoryMethod


type IProduct interface {
	getName() string
}

type AProduct struct {
	AName string
}

type BProduct struct {
	BName string
}

func (product *BProduct) getName() string  {
	return product.BName
}

func (product *AProduct) getName() string  {
	return product.AName
}


type IProductFactory interface {
	FactoryMethod() IProduct
}

type ProductFactory struct {
	Factory IProductFactory
}

type AFactory struct {}
type BFactory struct {}
func (product *AFactory) FactoryMethod() IProduct  {
	return  &AProduct{AName:"A"}
}
func (product *BFactory) FactoryMethod() IProduct  {
	return  &BProduct{BName:"B"}
}
