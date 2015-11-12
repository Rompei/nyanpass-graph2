package nyanpass

import (
	"testing"
)

func TestCreateImage(t *testing.T) {
	nyanpass := NewNyanpass()
	_, err := nyanpass.GetNyanpassWithDays(7)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = nyanpass.CreateImage("test_graph7.png")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	_, err = nyanpass.GetNyanpassWithDays(30)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = nyanpass.CreateImage("test_graph30.png")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
