package main

import(
	"testing"
)

// func unique_time

func TestUniqueTime(t *testing.T) {
	unique_time_ct = 0;

	// get fake time
	myTime := 0;
	unique_time(&myTime);

	// check that the time is 1
	if myTime != 1 {
		t.Errorf("f_time not working, set value to %d\n", myTime);
	}
}

// func init_pt && init_ft

func TestPageTableInit(t *testing.T) {
	pt := [NUM_PAGES]Page{};
	init_pt(&pt)

	// test stored values
	if pt[0] != (Page{0, -1}) && 
	   pt[1] != (Page{1, -1}) && 
	   pt[2] != (Page{2, -1}) {
		t.Errorf("Page initialization returned unexpected value %v\n", pt)
	}
}

func TestFrameTableInit(t *testing.T) {
	ft := [NUM_FRAMES]Frame{};
	init_ft(&ft)

	// test stored values
	if ft[0] != (Frame{0, -1, 0, -1, -1}) && 
	   ft[1] != (Frame{1, -1, 0, -1, -1}) && 
	   ft[2] != (Frame{2, -1, 0, -1, -1}) {
		t.Errorf("Frame initialization returned unexpected value %v\n", ft)
	}
}

// func findEmptyFrame

func TestFindEmptyFrame_empty(t *testing.T) {
	// create empty frame table
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)

	// find empty frame
	vf_index, _ := findEmptyFrame(ft)

	if vf_index != 0 {
		t.Errorf("FindEmptyFrame returned %v, expected 0", vf_index)
	}
}

func TestFindEmptyFrame_partial(t *testing.T) {
	// create frame table with 1 entry
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	ft[0] = Frame{0, 3, 0, 0, 0}

	// find empty frame
	vf_index, _ := findEmptyFrame(ft)

	if vf_index != 1 {
		t.Errorf("FindEmptyFrame returned %v, expected 1", vf_index)
	}
}

func TestFindEmptyFrame_full(t *testing.T) {
	// create full frametable
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	for i := 0; i < NUM_FRAMES; i++ {
		ft[i].pageNumber = i
	}

	// attempt finding empty frame
	vf_index, _ := findEmptyFrame(ft)

	if vf_index != -1 {
		t.Errorf("FindEmptyFrame returned %v, expected %v", 
		         vf_index, -1)
	}
}

// func accessPage

func TestAccessPage_exists(t *testing.T) {
	// mock page exists 
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	pt[2] = Page{2, 1}
	ft[1] = Frame{1, 2, 0, -1, -1}

	// run accesspage
	if !accessPage(2, pt, &ft) {
		t.Errorf("Access page did not find page, expected to find page")
	}
}

func TestAccessPage_timeUpdate(t *testing.T) {
	// mock page exists 
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	// run access page
	pt[2] = Page{2, 1}
	ft[1] = Frame{1, 2, 0, -1, -1}
	accessPage(2, pt, &ft)

	// check if time was updated
	if ft[1].lastReference != 1 {
		t.Errorf("Access page did not update page")
	}
}

func TestAccessPage_pageFault(t *testing.T) {
	// mock page exists 
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	// check that accessPage fails
	if accessPage(2, pt, &ft) {
		t.Errorf("Access page returned true and expected false")
	}
}

// func fifo

func TestFIFO(t *testing.T) {
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)

	// mock ascending allocation times so last frame will be used
	for i := 0; i < NUM_FRAMES; i++ {
		ft[i].lastAllocation = NUM_FRAMES - i;
	}

	victim := fifo(ft)

	if victim != NUM_FRAMES - 1 {
		t.Errorf("fifo algorithm returned %v, expected %v", 
		         victim, NUM_FRAMES - 1)
	}
}

// func handlePageFault

func TestHandlePageFault_empty(t *testing.T) {
	// mock page exists 
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	handlePageFault(2, &pt, &ft, fifo)

	// check that page was inserted
	if !accessPage(2, pt, &ft) {
		t.Errorf("handlePageFault did not replace %v", 2)
	}
}

func TestHandlePageFault_full_fifo(t *testing.T) {
	// mock page exists 
	pt := [NUM_PAGES]Page{}
	init_pt(&pt)
	ft := [NUM_FRAMES]Frame{}
	init_ft(&ft)
	unique_time_ct = 0

	// mock ascending allocation times so last frame will be used
	for i := 0; i < NUM_FRAMES; i++ {
		ft[i].lastAllocation = NUM_FRAMES - i
		ft[i].pageNumber = 1
	}


	handlePageFault(2, &pt, &ft, fifo)

	// check that page was inserted
	if pt[2].frameNumber != NUM_FRAMES - 1 {
		t.Errorf("Handle Pagefault used pra incorrectly")
	}
}
