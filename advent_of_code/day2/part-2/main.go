package main

import (
	"bufio"
	"fmt"
	"os"
)

type draw rune
type strategy rune

type Player struct {
	score int
}

const (
	LOOSE strategy = 'X'
	DRAW  strategy = 'Y'
	WIN   strategy = 'Z'
)

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

// returns the draw the player should play againts the opponent to result in a loose
func (s strategy) loose(otherDraw draw) draw {
	switch otherDraw {
	case ROCK:
		return SCISSORS
	case PAPER:
		return ROCK
	case SCISSORS:
		return PAPER
	default:
		return 0
	}
}

// returns the draw the player should play againts the opponent to result in a win
func (s strategy) win(otherDraw draw) draw {
	switch otherDraw {
	case ROCK:
		return PAPER
	case PAPER:
		return SCISSORS
	case SCISSORS:
		return ROCK
	default:
		return 0
	}
}

// returns the draw the player should play againts the opponent to result in a draw
func (s strategy) draw(otherDraw draw) draw {
	return otherDraw
}

// returns the score of the draw
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

// checks if the current draw is equal to the passed draw
func (d draw) isDraw(d2 draw) bool {
	return d == d2
}

// checks if the current draw is stronger than the passed draw
func (d draw) isStrongerThan(d2 draw) bool {
	return (d == ROCK && d2 == SCISSORS) || (d == SCISSORS && d2 == PAPER) || (d == PAPER && d2 == ROCK)
}

// play a round of the game and update the score of the player
// the score is calculated based on the following rules:
// - if the player wins, he gets 6 points + the score of the draw
// - if the player looses, he gets 0 points + the score of the draw
// - if the player draws, he gets 3 points + the score of the draw
func (p *Player) play(myDraw draw, otherDraw draw) {
	if myDraw.isDraw(otherDraw) {
		p.score += SCORE_DRAW + myDraw.Score()
	} else if myDraw.isStrongerThan(otherDraw) {
		p.score += SCORE_WIN + myDraw.Score()
	} else {
		p.score += SCORE_LOST + myDraw.Score()
	}
}

// parse the input and return the internal representation of the draw
func parseDraw(d rune) draw {
	switch d {
	case 'A':
		return ROCK
	case 'B':
		return PAPER
	case 'C':
		return SCISSORS
	default:
		return 0
	}
}

// parse the input and return the internal representation of the strategy
// a strategy represents the players intention to win, loose or draw
func parseStrategy(s rune) strategy {
	switch s {
	case 'X':
		return LOOSE
	case 'Y':
		return DRAW
	case 'Z':
		return WIN
	default:
		return 0
	}
}

// drawFromStrategy returns the draw to play by the player against the opponent based on the given strategy
func (p *Player) drawFromStrategy(myInput rune, otherDraw draw) draw {
	myStrategy := parseStrategy(myInput)
	switch myStrategy {
	case DRAW:
		return myStrategy.draw(otherDraw)
	case WIN:
		return myStrategy.win(otherDraw)
	case LOOSE:
		return myStrategy.loose(otherDraw)
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

		var myInput, oppInput rune
		i, _ := fmt.Sscanf(string(line), "%c %c", &oppInput, &myInput)
		if i != 2 {
			// ignore line
			continue
		}
		otherDraw := parseDraw(oppInput)
		myDraw := me.drawFromStrategy(myInput, otherDraw)
		me.play(myDraw, otherDraw)

	}

	fmt.Println("My score:", me.score)
}
