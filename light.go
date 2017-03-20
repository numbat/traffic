package main

import "fmt"

// Lights have a direction and a current colour.
type Light struct {
	direction string
	colour    Colour
}

func (l Light) String() string {
	return fmt.Sprintf("%+v=%+v", l.direction, l.colour)
}