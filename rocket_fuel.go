package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"slices"
)

func main(){


	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rockets_count,_ := strconv.Atoi(scanner.Text())
	rockets_names := make([]string,rockets_count)
	rockets_fuel_consumption := make([][]string,rockets_count)
	how_many_patterns := make([]int,0)
	for i := 0; i < rockets_count; i++{
		scanner.Scan()
		line := strings.Fields(scanner.Text())									
		rockets_names[i] = line[0]
		rockets_fuel_consumption[i] = line[1:]
	}
	// fmt.Println(rockets_fuel_consumption)
	for _,fuel_cons := range rockets_fuel_consumption{
		subsets := make([][]int,0)
		for i := len(fuel_cons)-1; i>=3; i--{
			for idx, item := range fuel_cons[:i]{
				if idx != len(fuel_cons)-1{
				this_fuel_cons,_ := strconv.Atoi(item)
				next_fuel_cons,_ := strconv.Atoi(fuel_cons[idx+1])
				difference := next_fuel_cons - this_fuel_cons
				// fmt.Println("Difference: ",difference)
				count := 1
				subset := make([]int,0)
				subset = append(subset,this_fuel_cons)
				for i, it := range fuel_cons[idx+1:]{

					it_int,_ := strconv.Atoi(it)
					// fmt.Println("It: ",it_int)	
					differences := it_int - this_fuel_cons
					times := i + 1
					// fmt.Println("Times: ",times)
					// fmt.Println("Differences: ",differences)
					if differences == times*difference{
						count++
						subset = append(subset,it_int)

					}else{
						// fmt.Println("else occured")
						break
					}
					// fmt.Println("count is: ",count)
				
				}
				if count > 2{
					does_exist := false
					for _, s := range subsets{
						if slices.Equal(s, subset){
							does_exist = true
							break
						}
					} 
					if !does_exist{
						subsets = append(subsets , subset)
					}
				}

				
	}}}
	how_many_patterns = append(how_many_patterns, len(subsets))}

for n, c := range how_many_patterns{
	fmt.Println(rockets_names[n],c)
}}

