package main

import (
	"fmt"
)

const NUM_PAGES int = 9
const NUM_FRAMES int = 5

var unique_time_ct int = 0

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

func main() {
	fmt.Println("hi")
}

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

func unique_time(val *int) int {
	unique_time_ct++
	*val = unique_time_ct
	return *val
}

