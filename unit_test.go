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
