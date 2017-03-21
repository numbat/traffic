package main

import (
	"errors"
)

// Intersection is made up of lights, of which one is active.
// Intersection is purely a passive model and is told what to
// do by the simulation.
type Intersection struct {
	lights      []Light
	activeLight int
}

func (l Intersection) SetCurrentLightColour(c Colour) error {
	if len(l.lights) == 0 {
		return errors.New("Intersection: no traffic lights defined")
	}
	l.lights[l.activeLight].colour = c
	return nil
}

func (l *Intersection) NextLight() error {
	if len(l.lights) == 0 {
		return errors.New("Intersection: no traffic lights defined")
	}
	l.activeLight = (l.activeLight + 1) % len(l.lights)
	return nil
}

func (l Intersection) String() string {
	var out, delim string
	for _, light := range l.lights {
		out += delim + light.String()
		delim = " "
	}
	return out
}
