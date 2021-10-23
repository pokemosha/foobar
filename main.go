package main

import (
	"fmt"
	"log"
)

func recursion(i uint, num uint) {
	fmt.Printf("%d Foo", i)
	if i%3 == 0 {
		fmt.Print(" Bar")
	}
	fmt.Println()
	if i < num {
		recursion(i+1, num)
	}
}

func main() {
	var number uint
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal(err)
	}
	if number == 0 {
		log.Fatal("expected integer more than 0")
	}
	recursion(1, number)
}
