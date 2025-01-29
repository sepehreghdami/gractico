package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){


	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rockets_count,_ := strconv.Atoi(scanner.Text())


	rockets_names := make([]string,rockets_count)
	rockets_fuel_consumption := make([][]string,rockets_count)
	for i := 0; i < rockets_count; i++{
		scanner.Scan()
		line := strings.Fields(scanner.Text())									
		rockets_names[i] = line[0]
		rockets_fuel_consumption[i] = line[1:]
		
	}

	
	for _,fuel_cons := range rockets_fuel_consumption{
		for idx, item := range fuel_cons{
			if idx != len(fuel_cons)-1{
			this_fuel_cons,_ := strconv.Atoi(item)
			next_fuel_cons,_ := strconv.Atoi(fuel_cons[idx+1])
			difference := next_fuel_cons - this_fuel_cons
			fmt.Println("Difference: ",difference)
			count := 0
			for i, it := range fuel_cons[idx+1:]{

				it_int,_ := strconv.Atoi(it)
				fmt.Println("It: ",it_int)	
				differences := it_int - this_fuel_cons
				times := i + 1
				fmt.Println("Times: ",times)
				fmt.Println("Differences: ",differences)
				if differences == times*difference{
					count++
				}else{
					break
				}
			
			}
			if count > 2{
				fmt.Println("hi")
			}
	}}}


}