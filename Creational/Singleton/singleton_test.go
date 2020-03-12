package Singleton

import "testing"

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance();
	if instance1 == nil {
		t.Error("Unexpected nil")
	}
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Error("Two instance")
	}

	instance1.SetNumber(4)

	if instance2.Number != 4 {
		t.Error("Unexpected number")
	}
}
