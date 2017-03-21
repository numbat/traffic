# Traffic Lights Simulator

[![Build Status](https://travis-ci.org/numbat/traffic.svg?branch=master)](https://travis-ci.org/numbat/traffic)

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
Example outout for the default 30min runtime with 2 sets of traffic 
lights:
```
2017-03-20 20:51:57.995280131 +1100 AEDT: northsouth=green eastwest=red
2017-03-20 20:56:28.586938008 +1100 AEDT: northsouth=yellow eastwest=red
2017-03-20 20:56:58.593530762 +1100 AEDT: northsouth=red eastwest=green
2017-03-20 21:01:28.609086342 +1100 AEDT: northsouth=red eastwest=yellow
2017-03-20 21:01:58.613615768 +1100 AEDT: northsouth=green eastwest=red
2017-03-20 21:06:28.628319984 +1100 AEDT: northsouth=yellow eastwest=red
2017-03-20 21:06:58.638982457 +1100 AEDT: northsouth=red eastwest=green
2017-03-20 21:11:28.65796578 +1100 AEDT: northsouth=red eastwest=yellow
2017-03-20 21:11:58.667827442 +1100 AEDT: northsouth=green eastwest=red
2017-03-20 21:16:28.686941388 +1100 AEDT: northsouth=yellow eastwest=red
2017-03-20 21:16:58.692991034 +1100 AEDT: northsouth=red eastwest=green
2017-03-20 21:21:28.706819843 +1100 AEDT: northsouth=red eastwest=yellow
```

## Tests
You can run the tests with ``go test`` when in the source code directory:
```
$ go test
```