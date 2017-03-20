# Traffic Lights Simulator

Simulates a set of traffic lights at an intersection. The traffic lights 
are designated (N, S) and (E, W) like a compass. When switching from 
green to red, the yellow light must be displayed for 30 seconds prior to 
it switching to red. The lights will change automatically every 5 minutes.

Assumptions:
- The lights are initialised in a state where they are all red, and the
  first light will transition to green as part of starting the simulation.
- When a light switches from yellow to red, the next traffic light will
  immediatedly go green (there is no waiting time where all lights are
  red).

## Requirements
- [Go](https://golang.org/dl/)

## Installation
Clone this repository into ``$GOPATH/src/github.com/numbat``:
```
$ git clone https://github.com/numbat/traffic.git
```
Build the application with ``go install`` when in the source code 
directory:
```
$ go install
```

## Running the Simulation
The simulation can be run after building by going to ``$GOPATH`` and 
executing ``bin/traffic``:
```
$ bin/traffic
```

## Tests
You can run the tests with ``go test`` when in the source code directory:
```
$ go test
```