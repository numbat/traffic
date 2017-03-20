package main

import (
	"fmt"
	"time"
)

type Event int

const (
	turnGreen Event = iota
	turnYellow
	turnRed
)

func runTrafficLights(waitGreen time.Duration, waitYellow time.Duration, totalRuntime time.Duration) {
	eventChan := make(chan Event)
	finishChan := make(chan bool)
	go func() {
		time.Sleep(totalRuntime)
		finishChan <- true
	}()

	lights := Lights{[]Light{
		{"eastsouth", red},
		{"westnorth", red}}, 0}

	go func() {
		eventChan <- turnGreen
	}()

	for {
		select {
		case e := <-eventChan:
			switch e {
			case turnGreen:
				lights.CurrentLightToGreen()
				fmt.Println(time.Now().String() + ": " + lights.String())
				go func() {
					time.Sleep(waitGreen)
					eventChan <- turnYellow
				}()
			case turnYellow:
				lights.CurrentLightToYellow()
				fmt.Println(time.Now().String() + ": " + lights.String())
				go func() {
					time.Sleep(waitYellow)
					eventChan <- turnRed
				}()
			case turnRed:
				lights.CurrentLightToRed()
				lights.NextLight()
				go func() {
					eventChan <- turnGreen
				}()
			}
		case <-finishChan:
			return
		}
	}
}

func main() {
	runTrafficLights(
		time.Duration(time.Second*2),
		time.Duration(time.Second*1),
		time.Duration(time.Second*180))
}
