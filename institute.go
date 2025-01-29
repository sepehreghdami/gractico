package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	teachers_count,_ := strconv.Atoi(scanner.Text())

	teachers_names := make([]string, 0)
	teachers_scores := make([][]string, 0)
	for i := 0; i < teachers_count*2; i++ {
		scanner.Scan()
		if i % 2 == 0 {
			teachers_names = append(teachers_names, scanner.Text())
		}else {
			words := strings.Fields(scanner.Text())
			teachers_scores = append(teachers_scores, words)
			}
			}


	teachers_averages := make([]int, 0)
	for _,item := range teachers_scores {
		sum := 0
		for _,score := range item {
			score_int,_ := strconv.Atoi(score)
			sum += score_int
		}
	teachers_averages = append(teachers_averages, sum/len(item))}

	for idx,item := range teachers_names {
		if teachers_averages[idx] >= 80{
			fmt.Println(item,"Excellent")
		}else if teachers_averages[idx] >= 60 && teachers_averages[idx] < 80{
			fmt.Println(item,"Very Good")
		}else if teachers_averages[idx] >= 40 && teachers_averages[idx] < 60{
			fmt.Println(item,"Good")
		}else{
			fmt.Println(item,"Fair")}}
}