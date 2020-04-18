package main

import (
	"fmt"
	"errors"
)

// Both must be at least len >= 3 to run tests
const NUM_PAGES int = 9
const NUM_FRAMES int = 5

var unique_time_ct int = 0

/*** DATATYPES ***/

type Page struct {
	pageNumber int
	frameNumber int
}

type Frame struct {
	frameNumber int
	pageNumber int
	faultCt int
	lastAllocation int
	lastReference int
}

type PageReplacementAlgorithm func(ft [NUM_FRAMES]Frame) int

/*** MAIN ***/

func main() {
	fmt.Println("hi")
}

/*** INITIALIZATION ***/

func init_pt(pg *[NUM_PAGES]Page) {
	for i := 0; i < len(pg); i++ {
		pg[i] = Page{i, -1}
	}
}

func init_ft(ft *[NUM_FRAMES]Frame) {
	for i := 0; i < len(ft); i++ {
		ft[i] = Frame{i, -1, 0, -1, -1}
	}
}

/*** PAGE REPLACEMENT ALGORITHMS ***/

func fifo(ft [NUM_FRAMES]Frame) int {
	victim := 0
	for i := 1; i < NUM_FRAMES; i++ {
		if ft[i].lastAllocation < ft[victim].lastAllocation {
			victim = i
		}
	}
	return victim
}

/*** UTILITY FUNCTIONS ***/

func findEmptyFrame(ft [NUM_FRAMES]Frame) (int, error) {
	for i := 0; i < NUM_FRAMES; i++ {
		if ft[i].pageNumber == -1 {
			return i, nil
		}
	}
	return -1, errors.New("No empty frames available")
}

func accessPage(pageIndex int, pt [NUM_PAGES]Page, ft *[NUM_FRAMES]Frame) bool {
	// check that page exists in memory
	if pt[pageIndex].frameNumber == -1 {
		return false;
	}
	// update time (also simulates reference)
	unique_time(&ft[pt[pageIndex].frameNumber].lastReference)
	return true
}

func unique_time(val *int) int {
	unique_time_ct++
	*val = unique_time_ct
	return *val
}

