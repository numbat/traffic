package main

import (
	"fmt"
	"testing"
)

func TestControllerString(t *testing.T) {
	expected := "northsouth=red eastwest=red"
	intersection := LightsController{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 0}

	if s := intersection.String(); s != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, s))
	}
}
