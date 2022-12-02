package main

import (
	"bufio"
	"fmt"
	"os"
)

type draw rune

type Player struct {
	score int
}

const (
	ROCK     draw = 'R'
	PAPER    draw = 'P'
	SCISSORS draw = 'S'
)

const (
	SCORE_ROCK = 1 + iota
	SCORE_PAPER
	SCORE_SCISSORS
)

const (
	SCORE_LOST = 0
	SCORE_DRAW = 3
	SCORE_WIN  = 6
)

func (d draw) Score() int {
	switch d {
	case ROCK:
		return SCORE_ROCK
	case PAPER:
		return SCORE_PAPER
	case SCISSORS:
		return SCORE_SCISSORS
	default:
		return 0
	}
}

func (d draw) isDraw(d2 draw) bool {
	return d == d2
}

func (d draw) isStrongerThan(d2 draw) bool {
	return (d == ROCK && d2 == SCISSORS) || (d == SCISSORS && d2 == PAPER) || (d == PAPER && d2 == ROCK)
}

func (p *Player) play(myDraw draw, otherDraw draw) {
	if myDraw.isDraw(otherDraw) {
		p.score += SCORE_DRAW + myDraw.Score()
	} else if myDraw.isStrongerThan(otherDraw) {
		p.score += SCORE_WIN + myDraw.Score()
	} else {
		p.score += SCORE_LOST + myDraw.Score()
	}
}

func parseDraw(d rune) draw {
	switch d {
	case 'A', 'X':
		return ROCK
	case 'B', 'Y':
		return PAPER
	case 'C', 'Z':
		return SCISSORS
	default:
		return 0
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	filePath := os.Args[1]

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var reader = bufio.NewReader(f)
	var me = Player{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		var myDraw, otherDraw rune
		i, _ := fmt.Sscanf(string(line), "%c %c", &otherDraw, &myDraw)
		if i != 2 {
			// ignore line
			continue
		}
		me.play(parseDraw(myDraw), parseDraw(otherDraw))

	}

	fmt.Println("My score:", me.score)
}
