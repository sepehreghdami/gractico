package main

import "fmt"


func main() {
	var first_name_one, first_name_two, second_name_one, second_name_two string
	
	fmt.Scanf("%s %s %s %s", &first_name_one, &first_name_two, &second_name_one, &second_name_two)
	first_name := first_name_one + " " + first_name_two
	second_name := second_name_one + " " + second_name_two
	if first_name == second_name {
		fmt.Println("together")
	} else {
		fmt.Println("lost")
	}
}