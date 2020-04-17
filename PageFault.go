package main

import (
	"fmt"
)

var unique_time_ct int = 0;

func main() {
	fmt.Println("hi");
}

func unique_time(val *int) int {
	unique_time_ct++;
	*val = unique_time_ct;
	return *val;
}

