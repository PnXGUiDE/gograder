package main

import (
	"./java"
	"log"
)

func main() {
	err := java.RunAllCases("Test.java", 1, "1.in", "1.out", "2.in", "2.out")
	if err != nil {
		log.Fatal(err.Error())
	}
}
