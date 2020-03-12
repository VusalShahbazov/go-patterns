package Builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	builderAProduct := AProductBuilder{}
	builderAProduct.SetPrice().SetCount().SetName()

	myAproduct := builderAProduct.GetProduct()

	if myAproduct.Price != 15.5 {
		t.Error("Builder does not set valid price")
	}
	if myAproduct.Name != "A" {
		t.Error("Builder does not set valid Name")
	}
	if myAproduct.Count != 1 {
		t.Error("Builder does not set valid Count")
	}

}