package AbstractFactory


import "testing"

func TestUi(t *testing.T) {
	bootstrap := Bootstrap{}
	grid := bootstrap.CreateGrid()
	if grid.getGrid() != "col-md-6" {
		t.Error("No bootstrap grid")
	}
	button := bootstrap.CreateButton()

	if button.getColor() != "#563d7c" {
		t.Error("No bootstrap button")
	}


	materialize := Materialize{}
	grid = materialize.CreateGrid()
	if grid.getGrid() != "s6" {
		t.Error("No materialize grid")
	}
	button = materialize.CreateButton()

	if button.getColor() != "#ee6e73" {
		t.Error("No materialize button")
	}
}