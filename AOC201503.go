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
	var currHouse House
	var nextHouse House
	var iter int = 0
	currHouse.x = 0
	currHouse.y = 0

	//HousesVisited contains a list of houses visited and the number of times they were visited.
	var HousesVisited = make(map[House]int)

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
		p(string(char))
		nextHouse = getNewHouse(currHouse, string(char))
		HousesVisited[nextHouse]++
		currHouse = nextHouse
	}
	fmt.Printf("Total houses visited: %d\n", (len(HousesVisited)))
}
