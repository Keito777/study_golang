package main

import "fmt"

func Shadow() {
	age := 10
	fmt.Println(age) // 10

	if true {
		age := 20
		fmt.Println(age) // 20

		age += 30
		fmt.Println(age) // 50
	}

	fmt.Println(age) // 10
}
