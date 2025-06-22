package main

import (
	"cell_automaton/model"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ruleNumber, length, count int
	var err error

	args := os.Args

	if len(args) >= 2 {
		ruleNumber, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	} else {
		ruleNumber = 90
	}

	if len(args) >= 3 {
		length, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
	} else {
		length = 90
	}

	if len(args) >= 4 {
		count, err = strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
	} else {
		count = 45
	}

	rule := model.NewRule(uint(ruleNumber))
	gen := model.NewGen(uint(length))
	fmt.Println(gen)

	for i := 0; i < count; i++ {
		gen = gen.Next(*rule)
		fmt.Println(gen)
	}
}
