package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	oppRock     = "A"
	oppPaper    = "B"
	oppScissors = "C"
)

const (
	youRock     = "X"
	youPaper    = "Y"
	youScissors = "Z"
)

const (
	youLose    = "X"
	youDraw    = "Y"
	youWin 	   = "Z"
)


func checkStrategyTwo(strategy []byte, yourPoints *int) {
	opponentPlayed := string(strategy[0])
	youPlayed := string(strategy[2])

	switch opponentPlayed {
		case oppRock:
			if youPlayed == youDraw {
				*yourPoints += 3 + 1
			} else if youPlayed == youWin {
				*yourPoints += 6 + 2
			} else if youPlayed == youLose {
				*yourPoints += 0 + 3
			}
		case oppPaper:
			if youPlayed == youDraw {
				*yourPoints += 3 + 2
			} else if youPlayed == youLose {
				*yourPoints += 0 + 1
			} else if youPlayed == youWin {
				*yourPoints += 6 + 3
			}
		case oppScissors:
			if youPlayed == youDraw {
				*yourPoints += 3 + 3
			} else if youPlayed == youWin {
				*yourPoints += 6 + 1
			} else if youPlayed == youLose {
				*yourPoints += 0 + 2
			}
	}
}

func checkStrategy(strategy []byte, yourPoints *int) {
	opponentPlayed := string(strategy[0])
	youPlayed := string(strategy[2])

	switch opponentPlayed {
		case oppRock:
			if youPlayed == youRock {
				*yourPoints += 3 + 1
			} else if youPlayed == youPaper {
				*yourPoints += 6 + 2
			} else if youPlayed == youScissors {
				*yourPoints += 0 + 3
			}
		case oppPaper:
			if youPlayed == youPaper {
				*yourPoints += 3 + 2
			} else if youPlayed == youRock {
				*yourPoints += 0 + 1
			} else if youPlayed == youScissors {
				*yourPoints += 6 + 3
			}
		case oppScissors:
			if youPlayed == youScissors {
				*yourPoints += 3 + 3
			} else if youPlayed == youRock {
				*yourPoints += 6 + 1
			} else if youPlayed == youPaper {
				*yourPoints += 0 + 2
			}
	}
}

func main() {
	f, err := os.Open("day2/strategy.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	reader := bufio.NewReader(f)
	yourPoints := 0
	yourPointsTwo := 0

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
		}

		checkStrategy(line,	&yourPoints)
		checkStrategyTwo(line, &yourPointsTwo)
	}

	fmt.Println("Day Two Solutions:")
	fmt.Printf("- Your Total Points %d\n", yourPoints)
	fmt.Printf("- Your Total Points Two %d\n", yourPointsTwo)
}
