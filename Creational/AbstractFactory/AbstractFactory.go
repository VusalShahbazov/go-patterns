package AbstractFactory

type IButton interface {
	getColor() string
}

type IGrid interface {
	getGrid() string
}

type BootstrapButton struct {}
type BootstrapGrid struct {}


func (b *BootstrapGrid) getGrid() string {
	return "col-md-6"
}

func (b *BootstrapButton) getColor() string {
	return "#563d7c" // Bootstrap  color
}

type MaterializeButton struct {}
type MaterializeGrid struct {}

func (b *MaterializeGrid) getGrid() string {
	return "s6"
}

func (b *MaterializeButton) getColor() string {
	return "#ee6e73" // Materialize color
}
type Ui interface {
	CreateButton() IButton
	CreateGrid() IGrid
}
type Bootstrap struct {}
type Materialize struct {}

func (b *Bootstrap) CreateButton() IButton {
	return  &BootstrapButton{}
}

func (b *Bootstrap) CreateGrid() IGrid {
	return &BootstrapGrid{}
}
func (b *Materialize) CreateButton() IButton {
	return  &MaterializeButton{}
}

func (b *Materialize) CreateGrid() IGrid {
	return &MaterializeGrid{}
}