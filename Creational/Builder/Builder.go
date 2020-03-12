package Builder

type Builder interface {
	SetName() Builder
	SetCount() Builder
	SetPrice() Builder
	GetProduct() Product
}

type Product struct {
	Name string
	Price float64
	Count uint
}

type AProduct struct{
	P Product
}
type AProductBuilder struct {
	Builder
	P Product
}
func (builder *AProductBuilder) GetProduct() Product {
	return builder.P
}
func (builder *AProductBuilder) SetPrice() Builder {
	builder.P.Price = 15.5
	return builder
}
func (builder *AProductBuilder) SetCount() Builder {
	builder.P.Count = 1
	return builder
}
func (builder *AProductBuilder) SetName() Builder {
	builder.P.Name = "A"
	return builder
}


