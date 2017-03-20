package main

import (
	"bytes"
	"fmt"
)

// There are three traffic light colours, as defined in the
// constants: red, yellow and green.
type Colour string

const (
	red    Colour = "red"
	yellow Colour = "yellow"
	green  Colour = "green"
)

// Lights have a direction and a current colour.
type Light struct {
	direction string
	colour    Colour
}

func (l Light) String() string {
	return fmt.Sprintf("%+v=%+v", l.direction, l.colour)
}

// Intersection is made up of lights, of which one is active.
// Intersection is purely a passive model and is told what to
// do by the simulation.
type Intersection struct {
	lights      []Light
	activeLight int
}

func (l Intersection) CurrentLightToGreen() {
	l.lights[l.activeLight].colour = green
}

func (l Intersection) CurrentLightToYellow() {
	l.lights[l.activeLight].colour = yellow
}

func (l Intersection) CurrentLightToRed() {
	l.lights[l.activeLight].colour = red
}

func (l *Intersection) NextLight() {
	l.activeLight = (l.activeLight + 1) % len(l.lights)
}

func (l Intersection) String() string {
	var buffer bytes.Buffer
	for i, light := range l.lights {
		buffer.WriteString(light.String())
		if i < len(l.lights)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
