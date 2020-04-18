package main

import(
	"testing"
)

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