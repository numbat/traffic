package main

import (
	"fmt"
	"testing"
)

func TestSetNextLightWhenLightsEmptyShouldError(t *testing.T) {
	lights := []Light{}

	intersection := Intersection{lights, 0}

	if err := intersection.NextLight(); err == nil {
		t.Error("Expected error as no lights defined.")
	}
}

func TestActiveLightDefaultsToZero(t *testing.T) {
	intersection := Intersection{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}}

	if active := intersection.activeLight; active != 0 {
		t.Errorf("Expected active light to be 0, but it was %d instead.", active)
	}
}

func TestSetNextLightShouldIncrement(t *testing.T) {
	intersection := Intersection{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}}

	intersection.NextLight()

	if active := intersection.activeLight; active != 1 {
		t.Errorf("Expected active light to be 1, but it was %d instead.", active)
	}
}

func TestSetNextLightShouldLoop(t *testing.T) {
	intersection := Intersection{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 1}

	intersection.NextLight()

	if active := intersection.activeLight; active != 0 {
		t.Errorf("Expected active light is 0, but it was %d instead.", active)
	}
}

func TestSetCurrentLightColour(t *testing.T) {
	intersection := Intersection{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 0}

	intersection.SetCurrentLightColour(green)

	if c := intersection.lights[intersection.activeLight].colour; c != green {
		t.Errorf("Expected active light to be green, but it was %d instead.", c)
	}
}

func TestSetCurrentLightColourWhenLightsEmptyShouldError(t *testing.T) {
	intersection := Intersection{lights: []Light{}, activeLight: 0}

	if err := intersection.SetCurrentLightColour(green); err == nil {
		t.Error("Expected error as no lights defined.")
	}
}

func TestIntersectionString(t *testing.T) {
	expected := "northsouth=red eastwest=red"
	intersection := Intersection{lights: []Light{
		{"northsouth", red},
		{"eastwest", red}}, activeLight: 0}

	if s := intersection.String(); s != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, s))
	}
}

func TestIntersectionStringWhenLightsEmpty(t *testing.T) {
	expected := ""
	intersection := Intersection{lights: []Light{}}

	if s := intersection.String(); s != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, s))
	}
}
