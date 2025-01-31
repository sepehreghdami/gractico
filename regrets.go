package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	// "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input_count, energy_left int
	var input [][]string

	fmt.Scanf("%d %d", &input_count, &energy_left)

	scanner.Scan()

	for i := 0; i < input_count; i++ {
		scanner.Scan()
		inp := []string{scanner.Text()}
		input = append(input, inp)
	}
	fmt.Println(input)
	// var disappeard_opps []str
	// for idx, element := range input{
	// 	var earlier_inputs []string = input[:idx]
	// 	for _,el := range earlier_inputs{
	// 		if el == element {
	// 			disappeard_opps = append(disappeard_opps, el)
	// 			break

	// 		}
	// 	}
	// }
}