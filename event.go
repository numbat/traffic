package main

import "time"

// Events trigger logic inside the simulation to affect the lights in the
// intersection. Event provides all the necessary parameters for affecting
// the intersection and lights in the simulation.
type Event struct {
	changeColourTo    Colour
	eventDuration     time.Duration
	nextEvent         *Event
	changeActiveLight bool
	printLightStates  bool
}

var goGreen Event
var goYellow Event
var goRed Event

func init() {
	goGreen = Event{
		changeColourTo:    green,
		eventDuration:     time.Duration(time.Second * 270),
		nextEvent:         &goYellow,
		changeActiveLight: false,
		printLightStates:  true}
	goYellow = Event{
		changeColourTo:    yellow,
		eventDuration:     time.Duration(time.Second * 30),
		nextEvent:         &goRed,
		changeActiveLight: false,
		printLightStates:  true}
	goRed = Event{
		changeColourTo:    red,
		eventDuration:     time.Duration(time.Second * 0),
		nextEvent:         &goGreen,
		changeActiveLight: true,
		printLightStates:  false}
}
