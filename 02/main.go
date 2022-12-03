package main

import (
	"fmt"
	"strings"

	"github.com/tylerperdue/advent-of-code-2022/input"
)

func main() {
	rounds, err := GetRoundsFromPuzzleInput("02/input.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("rounds: %+v\n", rounds)

	fmt.Printf("part one: %d\n", FinalScore(WhatIShouldPlayStrategy{}, rounds))
	fmt.Printf("part two: %d\n", FinalScore(HowTheRoundNeedsToEndStrategy{}, rounds))
}

func GetRoundsFromPuzzleInput(filename string) ([]Round, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return nil, fmt.Errorf("[input.ReadLines]: %w", err)
	}

	var rounds []Round

	for _, line := range lines {
		choices := strings.Split(line, " ")

		rounds = append(rounds, Round{
			MyOpponentsChoice: OppontentsChoice(choices[0]),
			MyChoiceOrResult:  choices[1],
		})
	}

	return rounds, nil
}

func OppontentsChoice(s string) Choice {
	switch s {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		return -1
	}
}

func FinalScore(s Strategy, rounds []Round) int {
	var t int

	for _, round := range rounds {
		t += s.Result(round) + s.MyChoice(round).Points()
	}

	return t
}

type Strategy interface {
	Result(Round) int
	MyChoice(Round) Choice
}

type Round struct {
	MyOpponentsChoice Choice
	MyChoiceOrResult  string
}

const (
	Winner int = 6
	Loser  int = 0
	Draw   int = 3
)

type Choice int

const (
	Rock Choice = iota
	Paper
	Scissors
)

func (c Choice) Points() int {
	switch c {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

type WhatIShouldPlayStrategy struct{}

func (WhatIShouldPlayStrategy) Result(r Round) int {
	switch r.MyOpponentsChoice {
	case Rock:
		if r.MyChoiceOrResult == "X" {
			return Draw
		}

		if r.MyChoiceOrResult == "Y" {
			return Winner
		}

		return Loser
	case Paper:
		if r.MyChoiceOrResult == "Y" {
			return Draw
		}

		if r.MyChoiceOrResult == "Z" {
			return Winner
		}

		return Loser
	case Scissors:
		if r.MyChoiceOrResult == "Z" {
			return Draw
		}

		if r.MyChoiceOrResult == "X" {
			return Winner
		}

		return Loser
	default:
		return -1
	}
}

func (WhatIShouldPlayStrategy) MyChoice(r Round) Choice {
	switch r.MyChoiceOrResult {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		return -1
	}
}

type HowTheRoundNeedsToEndStrategy struct{}

func (HowTheRoundNeedsToEndStrategy) Result(r Round) int {
	switch r.MyChoiceOrResult {
	case "X":
		return Loser
	case "Y":
		return Draw
	case "Z":
		return Winner
	default:
		return -1
	}
}

func (htrntes HowTheRoundNeedsToEndStrategy) MyChoice(r Round) Choice {
	switch htrntes.Result(r) {
	case Winner:
		return WinningChoice(r.MyOpponentsChoice)
	case Loser:
		return LosingChoice(r.MyOpponentsChoice)
	case Draw:
		return r.MyOpponentsChoice
	default:
		return -1
	}
}

func WinningChoice(c Choice) Choice {
	switch c {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return -1
	}
}

func LosingChoice(c Choice) Choice {
	switch c {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		return -1
	}
}
