package main

import (
	"fmt"
)

func main() {
	months := [...]string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	Q2 := months[3:6]
	Q1 := months[:3]
	Q3 := months[6:9]
	Q4 := months[9:]
	summer := months[5:8]

	fmt.Println("Q2:", Q2)
	fmt.Println("Q1:", Q1)
	fmt.Println("Q3:", Q3)
	fmt.Println("Q4:", Q4)
	fmt.Println("Summer:", summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println("Months array address:", &months)
}
