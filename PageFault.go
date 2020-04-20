package main

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Both must be at least len >= 3 to run tests
const NUM_PAGES int = 9
const NUM_FRAMES int = 5
const NUM_REFERENCES int = 80

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
	dist := randomDistribution()

	testAlgorithm(fifo, dist, "\nFIFO")
	testAlgorithm(lru, dist, "\nLRU")
}

/*** DEMO ***/

func testAlgorithm(pra PageReplacementAlgorithm, dist [NUM_REFERENCES]int, 
	               name string) {
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	for i := 0; i < NUM_REFERENCES; i++ {
		//TODO Remove the -1 from this statement, it is due to incompatability with class data
		simulate(dist[i], &pt, &ft, pra)
	}

	printResults(ft, name)
}

func printResults(ft [NUM_FRAMES]Frame, alg string) {
	fmt.Printf("%s:\n", alg)
	
	for i := 0; i < NUM_FRAMES; i++ {
		fmt.Printf("Frame %d: %d page faults\n", i, ft[i].faultCt)
	}
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

func lru(ft [NUM_FRAMES]Frame) int {
	victim := 0
	for i := 1; i < NUM_FRAMES; i++ {
		if ft[i].lastReference < ft[victim].lastReference {
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

func handlePageFault(pageIndex int, pt *[NUM_PAGES]Page, ft *[NUM_FRAMES]Frame, 
	                 pra PageReplacementAlgorithm) {
	// Find frame to insert page into
	victim, _ := findEmptyFrame(*ft)
	if victim == -1 {
		victim = pra(*ft)
		pt[ft[victim].pageNumber].frameNumber = -1
	}

	// Replace page
	ft[victim].faultCt++
	ft[victim].pageNumber = pageIndex
	unique_time(&ft[victim].lastAllocation)
	pt[pageIndex].frameNumber = victim
}

func simulate(ref int, pt *[NUM_PAGES]Page, ft *[NUM_FRAMES]Frame, 
	          pra PageReplacementAlgorithm) error {
	// access page and handle pagefault if it occurs
	if !accessPage(ref, *pt, ft) {
		handlePageFault(ref, pt, ft, pra)

		// check for error
		if !accessPage(ref, *pt, ft) {
			return errors.New("handlePageFault failed in simulate")
		}
	}

	return nil
}

func randomDistribution() [NUM_REFERENCES]int {
	dist := [NUM_REFERENCES]int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < NUM_REFERENCES; i++ {
		dist[i] = rand.Int() % NUM_PAGES
	}
	return dist
}

func unique_time(val *int) int {
	unique_time_ct++
	*val = unique_time_ct
	return *val
}

