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

	lights := Lights{[]Light{
		{"eastsouth", red},
		{"westnorth", red}}, 0}

	go func() {
		time.Sleep(time.Second * 1)
		yellowToRedChan <- true
	}()

	for {
		select {
		case <-yellowToRedChan:
			lights.CurrentLightToRed()
			lights.NextLight()
			lights.CurrentLightToGreen()
			fmt.Println(time.Now().String() + ": " + lights.String())
			go func() {
				time.Sleep(waitGreen)
				greenToYellowChan <- true
			}()
		case <-greenToYellowChan:
			lights.CurrentLightToYellow()
			fmt.Println(time.Now().String() + ": " + lights.String())
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
