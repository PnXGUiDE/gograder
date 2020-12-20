package main

import (
	"log"

	"github.com/PnXGUiDE/gograder/model"
)

func main() {
	var grader model.Grader
	grader = model.JavaGrader{}
	err := grader.RunAllCases("Test.java", 1, "1.in", "1.out", "2.in", "2.out")
	if err != nil {
		log.Fatal(err.Error())
	}
}
