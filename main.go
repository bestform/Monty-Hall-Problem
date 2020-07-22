package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stats struct {
	WinChose *[3]int
	PlayerChose *[3]int
}

type Door struct {
	Win bool
	Chosen bool
	Eliminated bool
	WasFirstChoice bool
}

const numberOfDoors = 3
const numberOfRuns = 1000000

type Doors [numberOfDoors]Door

func NewDoors() *Doors {
	doors := new(Doors)

	return doors
}

func (doors *Doors) init() {
	winningDoor := rand.Int() % numberOfDoors
	chosenDoor := rand.Int() % numberOfDoors
	for i, _ := range doors {
		if i == winningDoor {
			doors[i].Win = true
		}
		if i == chosenDoor {
			doors[i].Chosen = true
			doors[i].WasFirstChoice = true
		}
	}
}

func (doors *Doors) eliminate() {
	eliminated := 0
	for i, door := range doors {
		if door.Win || door.Chosen {
			continue
		}
		doors[i].Eliminated = true
		eliminated++
		if eliminated == numberOfDoors - 2 {
			break
		}
	}
}

func (doors *Doors) switchChoice() {
	var oldChoice int
	var newChoice int

	for i, _ := range doors {
		if doors[i].Chosen {
			oldChoice = i
		}
		doors[i].Chosen = false
	}

	for i, door := range doors {
		if door.Eliminated {
			continue
		}
		if i == oldChoice {
			continue
		}
		newChoice = i
		break
	}

	doors[newChoice].Chosen = true
}

func (doors *Doors) won() bool {
	for _, door := range doors {
		if door.Win && door.Chosen {
			return true
		}
	}

	return false
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	total := numberOfRuns
	wonWithSwitch := 0
	wonWithoutSwitch := 0

	// with switch
	for i := 0; i < total; i++ {
		doors := NewDoors()
		doors.init()
		doors.eliminate()
		doors.switchChoice()
		if doors.won() {
			wonWithSwitch++
		}
	}
	// without switch
	for i := 0; i < total; i++ {
		doors := NewDoors()
		doors.init()
		doors.eliminate()
		if doors.won() {
			wonWithoutSwitch++
		}
	}

	fmt.Println("Won with switch:", float64(wonWithSwitch) / float64(total) * 100, "%")
	fmt.Println("Won without switch:", float64(wonWithoutSwitch) / float64(total) * 100, "%")
}
