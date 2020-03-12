package FactoryMethod

import (
	"testing"
)
func TestAFactory(t *testing.T) {
	factory := ProductFactory{Factory:&AFactory{}}

	prod := factory.Factory.FactoryMethod()
	if prod.getName() != "A" {
		t.Error("Invalid AProduct")
	}
}
func TestBFactory(t *testing.T) {
	factory := ProductFactory{Factory:&BFactory{}}

	prod := factory.Factory.FactoryMethod()
	if prod.getName() != "B" {
		t.Error("Invalid BProduct")
	}
}
//
