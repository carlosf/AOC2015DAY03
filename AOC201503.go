// Main package for Advent of Code 2015, Day 3.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(s string) {
	fmt.Println(s)
}

// House is a struct to hold the x and y coordinates of a house.
type House struct {
	x int
	y int
}

// GetNewHouse returns a new house based on the current house and the move.
func getNewHouse(h House, m string) House {
	var newX, newY int
	switch m {
	case "^":
		newX = h.x
		newY = h.y + 1
	case "v":
		newX = h.x
		newY = h.y - 1
	case ">":
		newX = h.x + 1
		newY = h.y
	case "<":
		newX = h.x - 1
		newY = h.y
	default:
		panic("Invalid move")
	}
	return House{newX, newY}

}

func main() {
	p("Starting...")
	var currHouseSanta House
	var currHouseRoboSanta House
	var nextHouseSanta House
	var nextHouseRoboSanta House
	var iter int = 0
	currHouseSanta.x = 0
	currHouseSanta.y = 0
	currHouseRoboSanta.x = 0
	currHouseRoboSanta.y = 0

	//HousesVisited contains a list of houses visited and the number of times they were visited.
	var HousesVisitedSanta = make(map[House]int)
	var HousesVisitedRoboSanta = make(map[House]int)

	//Set start position of Santa & RoboSanta
	HousesVisitedSanta[currHouseSanta]++
	HousesVisitedRoboSanta[currHouseRoboSanta]++

	//Open input.txt file
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}
		iter++
		//Santa move
		if iter%2 != 0 {
			//	p(string(char))
			nextHouseSanta = getNewHouse(currHouseSanta, string(char))
			HousesVisitedSanta[nextHouseSanta]++
			currHouseSanta = nextHouseSanta
		} else { //RoboSanta move
			//	p(string(char))
			nextHouseRoboSanta = getNewHouse(currHouseRoboSanta, string(char))
			HousesVisitedRoboSanta[nextHouseRoboSanta]++
			currHouseRoboSanta = nextHouseRoboSanta
		}

	}
	//merge the two maps
	for k, v := range HousesVisitedRoboSanta {
		HousesVisitedSanta[k] += v
	}
	fmt.Printf("Total houses visited: %d\n", (len(HousesVisitedSanta)))
}
