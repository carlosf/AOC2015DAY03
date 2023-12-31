// Main package for Advent of Code 2015, Day 3.
package main

import (
	"fmt"
	"io/ioutil"
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
	currHouse.x = 0
	currHouse.y = 0

	// Read the input file.
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	currHouse = getNewHouse(currHouse, "v")
}
