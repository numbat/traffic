package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestRunInitialisationState(t *testing.T) {
	goGreen.eventDuration = time.Duration(time.Second * 10)

	output := new(bytes.Buffer)
	intersection := Intersection{[]Light{
		{"northsouth", red},
		{"eastwest", red}}, 0}

	runSimulation(intersection, time.Millisecond*50, output)

	var expected = "northsouth=green eastwest=red\n"

	result := output.String()[len(output.String())-len(expected):]

	if result != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, result))
	}
}

func TestRunTransitionStateToYellow(t *testing.T) {
	goGreen.eventDuration = time.Duration(time.Second * 0)
	goYellow.eventDuration = time.Duration(time.Second * 10)

	output := new(bytes.Buffer)
	intersection := Intersection{[]Light{
		{"northsouth", red},
		{"eastwest", red}}, 0}

	runSimulation(intersection, time.Millisecond*50, output)

	var expected = "northsouth=yellow eastwest=red\n"

	result := output.String()[len(output.String())-len(expected):]

	if result != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, result))
	}
}

func TestRunTransitionStateGreenToYellowToRed(t *testing.T) {
	goGreen.eventDuration = time.Duration(time.Second * 0)
	goYellow.eventDuration = time.Duration(time.Second * 0)
	goRed.eventDuration = time.Duration(time.Second * 10)
	goRed.printLightStates = true

	output := new(bytes.Buffer)
	intersection := Intersection{[]Light{
		{"northsouth", red},
		{"eastwest", red}}, 0}

	runSimulation(intersection, time.Millisecond*50, output)

	var expected = "northsouth=red eastwest=red\n"

	result := output.String()[len(output.String())-len(expected):]

	if result != expected {
		t.Error(fmt.Sprintf("Expected output '%v' but got '%v'", expected, result))
	}
}

func TestRunNoLights(t *testing.T) {
	output := new(bytes.Buffer)
	intersection := Intersection{[]Light{}, 0}

	if err := runSimulation(intersection, time.Millisecond*50, output); err == nil {
		t.Error("Expected error as no lights defined.")
	}
}
