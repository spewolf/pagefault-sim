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

	if pt[0] != (Page{0, -1}) && 
	   pt[1] != (Page{1, -1}) && 
	   pt[2] != (Page{2, -1}) {
		t.Errorf("Page initialization returned unexpected value %v\n", pt)
	}
}

func TestFrameTableInit(t *testing.T) {
	ft := [NUM_FRAMES]Frame{};
	init_ft(&ft)

	if ft[0] != (Frame{0, -1, 0, -1, -1}) && 
	   ft[1] != (Frame{1, -1, 0, -1, -1}) && 
	   ft[2] != (Frame{2, -1, 0, -1, -1}) {
		t.Errorf("Frame initialization returned unexpected value %v\n", ft)
	}
}
