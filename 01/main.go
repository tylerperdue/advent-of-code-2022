package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/tylerperdue/advent-of-code-2022/input"
)

func main() {
	inventory, err := GetInventoryFromPuzzleInput("01/input.txt")
	if err != nil {
		panic(err)
	}

	inventory.SortElvesByCalories()

	// fmt.Printf("inventory: %+v\n", inventory)

	fmt.Printf("part one: %d\n", GetTotalCalories([]Elf{inventory.Elves[0]}))
	fmt.Printf("part two: %d\n", GetTotalCalories(inventory.Elves[:3]))
}

type Inventory struct {
	Elves []Elf
}

func (i *Inventory) SortElvesByCalories() {
	temp := i.Elves

	sort.Slice(i.Elves, func(i, j int) bool {
		return temp[i].GetTotalCalories() > temp[j].GetTotalCalories()
	})

	i.Elves = temp
}

type Elf struct {
	Foods []Food
}

func GetTotalCalories(elves []Elf) int {
	var t int

	for _, e := range elves {
		t += e.GetTotalCalories()
	}

	return t
}

func (e Elf) GetTotalCalories() int {
	var t int

	for _, food := range e.Foods {
		t += food.Calories
	}

	return t
}

type Food struct {
	Calories int
}

func GetInventoryFromPuzzleInput(filename string) (Inventory, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return Inventory{}, fmt.Errorf("[input.ReadLines]: %w", err)
	}

	var elves []Elf

	c := 0
	for _, line := range lines {
		if line == "" {
			c++
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return Inventory{}, fmt.Errorf("[strconv.Atoi]: %w", err)
		}

		if c == len(elves) {
			elves = append(elves, Elf{})
		}

		elves[c].Foods = append(elves[c].Foods, Food{Calories: calories})
	}

	return Inventory{Elves: elves}, nil
}
