package main

import (
	"fmt"
	"strings"

	"github.com/tylerperdue/advent-of-code-2022/input"
)

func main() {
	rucksacks, err := GetRucksacksFromPuzzleInput("03/input.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("rucksacks: %+v\n", rucksacks[0])

	fmt.Printf("part one: %d\n", SumOfPriorities(rucksacks))
	fmt.Printf("part two: %d\n", SumOfPrioritiesForBadgeItems(rucksacks))
}

func GetRucksacksFromPuzzleInput(filename string) ([]Rucksack, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return nil, fmt.Errorf("[input.ReadLines]: %w", err)
	}

	var rucksacks []Rucksack

	for _, line := range lines {
		rucksacks = append(rucksacks, Rucksack{
			FirstCompartment:  ToItems(strings.Split(line[:len(line)/2], "")),
			SecondCompartment: ToItems(strings.Split(line[len(line)/2:], "")),
		})
	}

	return rucksacks, nil
}

func ToItems(strings []string) []Item {
	var items []Item

	for _, s := range strings {
		items = append(items, Item(s))
	}

	return items
}

type Rucksack struct {
	FirstCompartment  []Item
	SecondCompartment []Item
}

func (r Rucksack) PriorityItem() Item {
	m := make(map[Item]bool)

	for _, item := range r.FirstCompartment {
		m[item] = true
	}

	for _, item := range r.SecondCompartment {
		if _, ok := m[item]; ok {
			return item
		}
	}

	return ""
}

type Item string

func (i Item) Priority() int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(i)) + 1
}

func SumOfPriorities(rucksacks []Rucksack) int {
	var c int

	for _, rucksack := range rucksacks {
		c += rucksack.PriorityItem().Priority()
	}

	return c
}

func SumOfPrioritiesForBadgeItems(rucksacks []Rucksack) int {
	var c int

	for i := 0; i < len(rucksacks); i += 3 {
		c += FindBadge(rucksacks[i : i+3]).Priority()
	}

	return c
}

func FindBadge(rucksacks []Rucksack) Item {
	m := make(map[Item]int)

	for _, rucksack := range rucksacks {
		t := make(map[Item]bool)

		for _, item := range rucksack.FirstCompartment {
			t[item] = true
		}

		for _, item := range rucksack.SecondCompartment {
			t[item] = true
		}

		for item := range t {
			m[item] += 1
		}
	}

	for item, c := range m {
		if c == 3 {
			return item
		}
	}

	return ""
}
