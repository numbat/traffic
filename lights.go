package main

import (
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