package main

import (
	"fmt"
	"time"
)

func runTrafficLights(waitGreen time.Duration, waitYellow time.Duration, totalRuntime time.Duration) {
	yellowToRedChan := make(chan bool)
	greenToYellowChan := make(chan bool)
	finishChan := make(chan bool)
	go func() {
		time.Sleep(totalRuntime)
		finishChan <- true
	}()

	controller := LightsController{[]Light{
		{"eastsouth", red},
		{"westnorth", red}}, 0}

	go func() {
		time.Sleep(time.Second * 1)
		yellowToRedChan <- true
	}()

	for {
		select {
		case <-yellowToRedChan:
			controller.lights[controller.activeLight].colour = red
			controller.activeLight = (controller.activeLight + 1) % len(controller.lights)
			controller.lights[controller.activeLight].colour = green
			fmt.Println(time.Now().String() + ": " + controller.String())
			go func() {
				time.Sleep(waitGreen)
				greenToYellowChan <- true
			}()
		case <-greenToYellowChan:
			controller.lights[controller.activeLight].colour = yellow
			fmt.Println(time.Now().String() + ": " + controller.String())
			go func() {
				time.Sleep(waitYellow)
				yellowToRedChan <- true
			}()
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
