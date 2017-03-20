package main

import (
	"bytes"
	"fmt"
)

type colour string

const (
	red    colour = "red"
	yellow colour = "yellow"
	green  colour = "green"
)

type Light struct {
	direction string
	colour    colour
}

func (l Light) String() string {
	return fmt.Sprintf("%+v=%+v", l.direction, l.colour)
}

type LightsController struct {
	lights      []Light
	activeLight int
}

func (l LightsController) String() string {
	var buffer bytes.Buffer
	for i, light := range l.lights {
		buffer.WriteString(light.String())
		if i < len(l.lights)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}
