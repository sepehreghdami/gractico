package main

import(
	"fmt"	
)

func main() {
	var cases_count,hospit_count float32 
	fmt.Scanf("%f\n%f",&cases_count, &hospit_count)
	hospit_case_ratio := hospit_count/cases_count
	var white_threshold, red_threshold float32 = 0.2, 0.6
	fmt.Println(hospit_case_ratio)
	if hospit_case_ratio <= white_threshold {
		fmt.Println("White")
	} else if hospit_case_ratio >= red_threshold {
		fmt.Println("Red")
	} else {
		fmt.Println("Yellow")
	}
	
}