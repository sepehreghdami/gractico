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
    line := scanner.Text()
	line_array := strings.Fields(line)
	p,_ := strconv.Atoi(line_array[0])
	q,_ := strconv.Atoi(line_array[1])

	for i := 1; i<q; i++ {
		remainder := i%p
		if remainder == 0 {
			hope_str := "Hope "
			repeat_count := i/p
			result := strings.Repeat(hope_str,repeat_count)
			fmt.Println(result)
		}else {
			fmt.Println(i)
		}
	}
	

	
}  