package main

import(
	"fmt"
)

func main(){

	var source , result string
	var time   int
	fmt.Scan(&source , &time)

	i := 0
	for i < time{
		result = result + source
		i++
	}

	fmt.Println(result)
}
