package main

import "fmt"

func main() {
    var income int
    fmt.Scanf("%d", &income)
    //since the least income amount is 1000 coins:
    const minmum_tax int = 120
    var income_exceeding_amount int = income - 1000
    tax := 20*income_exceeding_amount/100
    total_task := minmum_tax + tax
    fmt.Println(total_task)

    
}