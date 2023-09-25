package main

import (
	"fmt"
	"math/rand"
	"strings"
)

//func compare(line []*string, sign []string) bool {
//	return reflect.DeepEqual(line, sign)
//}

func compare(line []*string, sign []string) bool {
	var str string
	for i := 0; i < 3; i++ {
		str = *line[i]
		//fmt.Printf(str)
		if str != sign[i] {
			return false
		}
	}
	return true
}

// program do not check wrong cell select from input and will end game in 5 iterations anyway
func main() {
	board := [][]string{
		{"1", "|", "2", "|", "3"},
		{"-", "-", "-", "-", "-"},
		{"4", "|", "5", "|", "6"},
		{"-", "-", "-", "-", "-"},
		{"7", "|", "8", "|", "9"},
	}
	fmt.Printf("Board Example With Cell Numbers\n")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
	board[0][0] = " "
	board[0][2] = " "
	board[0][4] = " "
	board[2][0] = " "
	board[2][2] = " "
	board[2][4] = " "
	board[4][0] = " "
	board[4][2] = " "
	board[4][4] = " "

	line1Hor := []*string{&board[0][0], &board[0][2], &board[0][4]}
	line2Hor := []*string{&board[2][0], &board[2][2], &board[2][4]}
	line3Hor := []*string{&board[4][0], &board[4][2], &board[4][4]}
	line4Vert := []*string{&board[0][0], &board[2][0], &board[4][0]}
	line5Vert := []*string{&board[0][2], &board[2][2], &board[4][2]}
	line6Vert := []*string{&board[0][4], &board[2][4], &board[4][4]}
	line7Diag := []*string{&board[0][0], &board[2][2], &board[4][4]}
	line8Diag := []*string{&board[0][4], &board[2][2], &board[4][0]}

	lines := [8][]*string{}
	lines[0] = line1Hor
	lines[1] = line2Hor
	lines[2] = line3Hor
	lines[3] = line4Vert
	lines[4] = line5Vert
	lines[5] = line6Vert
	lines[6] = line7Diag
	lines[7] = line8Diag

	lineX := []string{"X", "X", "X"}
	lineY := []string{"O", "O", "O"}
	//board[0][0] = "X"
	//board[0][2] = "X"
	//board[0][4] = "X"

	//fmt.Printf("%t\n", compare(lines[0], lineX))
	cellsNumbers := [9][2]int{{0, 0}, {0, 2}, {0, 4}, {2, 0}, {2, 2}, {2, 4}, {4, 0}, {4, 2}, {4, 4}}
	var inp, iteration int
	winStatus := false
	xWin := false
	oWin := false

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
	for (winStatus == false) && (iteration < 5) {
		fmt.Printf("Choose number of empty cell (1-9): ")
		fmt.Scan(&inp)
		board[cellsNumbers[inp-1][0]][cellsNumbers[inp-1][1]] = "X"
		iteration++
		for i := 0; i < 8; i++ {
			if (compare(lines[i], lineX)) || (compare(lines[i], lineY)) {
				winStatus = true
				xWin = true
				break
			}
		}
		if winStatus == false && iteration != 5 {
			randomNumber := rand.Intn(9)
			for board[cellsNumbers[randomNumber][0]][cellsNumbers[randomNumber][1]] != " " {
				randomNumber = rand.Intn(9)
			}
			board[cellsNumbers[randomNumber][0]][cellsNumbers[randomNumber][1]] = "O"
		}
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], ""))
		}
		for i := 0; i < 8; i++ {
			if (compare(lines[i], lineX)) || (compare(lines[i], lineY)) {
				winStatus = true
				oWin = true
				break
			}
		}
	}
	if winStatus && xWin {
		fmt.Printf("You Win!\n")
	} else if winStatus && oWin {
		fmt.Printf("You Lose!\n")
	} else {
		fmt.Printf("Draw Game!\n")
	}

}
