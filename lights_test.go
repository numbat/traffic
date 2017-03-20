package main

import (
	"fmt"
	"testing"
)

func TestSetNextLightShouldIncrement(t *testing.T) {
	intersection := Lights{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}}

	intersection.NextLight()

	if active := intersection.activeLight; active != 1 {
		t.Errorf("Expected active light to be 1, but it was %d instead.", active)
	}
}

func TestSetNextLightShouldLoop(t *testing.T) {
	intersection := Lights{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 1}

	intersection.NextLight()

	if active := intersection.activeLight; active != 0 {
		t.Errorf("Expected active light is 0, but it was %d instead.", active)
	}
}

func TestControllerString(t *testing.T) {
	expected := "northsouth=red eastwest=red"
	intersection := Lights{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 0}

	if s := intersection.String(); s != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, s))
	}
}

func TestIntersectionStringWhenLightsEmpty(t *testing.T) {
	expected := ""
	intersection := Lights{lights: []Light{}}

	if s := intersection.String(); s != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, s))
	}
}
