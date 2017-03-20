package main

// Events trigger logic inside the simulation to
// affect the lights in the intersection.
type Event int

const (
	turnGreen Event = iota
	turnYellow
	turnRed
)
