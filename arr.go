package main

import("fmt")

func printA(array []int){
	for _, value := range array {
		fmt.Println("values===", value)
	}

	array[0]=5
}

func main(){
	array := []int{1,232,4242,11111111,10000000000}

	for _, value := range array{
		fmt.Println("values==",value)
	}
}