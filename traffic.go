package main

import (
	"bytes"
	"time"
	"fmt"
)

func printableState(lights []Light) string {
	var buffer bytes.Buffer
	buffer.WriteString(time.Now().String())
	buffer.WriteString(": ")
	for i, light := range lights {
		buffer.WriteString(light.String())
		if i < len(lights)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}

func runTrafficLights(waitGreen time.Duration, waitYellow time.Duration, totalRuntime time.Duration) {
	yellowToRedChan := make(chan bool)
	greenToYellowChan := make(chan bool)
	finishChan := make(chan bool)
	go func() {
		time.Sleep(totalRuntime)
		finishChan <- true
	}()

	lights := []Light{
		{"eastsouth", red},
		{"westnorth", red},
		{"northsouth", red},
		{"eastwest", red}}

	activeLight := 0

	go func() {
		time.Sleep(time.Second * 1)
		yellowToRedChan <- true
	}()

	for {
		select {
		case <-yellowToRedChan:
			lights[activeLight].colour = red
			activeLight = (activeLight + 1) % len(lights)
			lights[activeLight].colour = green
			fmt.Println(printableState(lights))
			go func() {
				time.Sleep(waitGreen)
				greenToYellowChan <- true
			}()
		case <-greenToYellowChan:
			lights[activeLight].colour = yellow
			fmt.Println(printableState(lights))
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