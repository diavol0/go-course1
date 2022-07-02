package main

import (
	"fmt"
)

func main(){

	i := 0
	for i <= 6 {
		fmt.Println(i)
		i = i + 1
	} // 0 1 2 3 4 5 6

	for i = 7; i <= 9; i++{
		fmt.Println(i)
	} // 7 8 9

	for {
		fmt.Println("loop")
		break
	} // loop

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
		continue
	}
    fmt.Println(n)
}
// 1
// 3
// 5
} // func main()
