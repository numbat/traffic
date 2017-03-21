package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"errors"
)

var simulation = runSimulation

// Traffic light simulation - run for 30min and publish
// all results to standard output with the provided
// intersection.
func main() {
	intersection := Intersection{[]Light{
		{"northsouth", red},
		{"eastwest", red}}, 0}

	simulation(intersection, time.Duration(time.Minute*30), os.Stdout)
}

// Run the simulation - pass in the duration to run, and
// the waiting times for the green and yellow lights.
func runSimulation(intersection Intersection, runDuration time.Duration, out io.Writer) error {
	if len(intersection.lights) == 0 {
		return errors.New("runSimulation: cannot run with no traffic lights")
	}

	eventChan := make(chan Event)
	finishChan := make(chan bool)
	go func() {
		time.Sleep(runDuration)
		finishChan <- true
	}()

	go func() {
		eventChan <- goGreen
	}()

	for {
		select {
		case e := <-eventChan:
			intersection.SetCurrentLightColour(e.changeColourTo)
			if e.printLightStates {
				fmt.Fprintln(out, time.Now().String()+": "+intersection.String())
			}
			if e.changeActiveLight {
				intersection.NextLight()
			}
			go func() {
				time.Sleep(e.eventDuration)
				eventChan <- *e.nextEvent
			}()
		case <-finishChan:
			return nil
		}
	}
}
