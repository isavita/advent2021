# Prompt 1
Q: 
"""
What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?

Your puzzle answer was 27027.

--- Part Two ---
On the other hand, it might be wise to try a different strategy: let the giant squid win.

You aren't sure how many bingo boards a giant squid could play at once, so rather than waste time counting its arms, the safe thing to do is to figure out which board will win last and choose that one. That way, no matter which boards it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after 13 is eventually called and its middle column is completely marked. If you were to keep playing until this point, the second board would have a sum of unmarked numbers equal to 148 for a final score of 148 * 13 = 1924.

Figure out which board will win last. Once it wins, what would its final score be?
"""
I: Please, write me a go program that solves Part Two and read the input from "day4_part2/input.txt".
A:
Sure, here's a Go program that solves Part Two of the puzzle using the input from "day4_part2/input.txt:
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Board struct {
    Rows     [][]int
    Columns  [][]int
    Diagonal []int
}

func NewBoard(rows [][]int) *Board {
    n := len(rows)
    columns := make([][]int, n)
    for i := range columns {
        columns[i] = make([]int, n)
    }
    for i, row := range rows {
        for j, num := range row {
            columns[j][i] = num
        }
    }
    diagonal := make([]int, n)
    for i := range rows {
        diagonal[i] = rows[i][i]
    }
    return &Board{
        Rows:     rows,
        Columns:  columns,
        Diagonal: diagonal,
    }
}

func (b *Board) MarkNumber(num int) bool {
    marked := false
    for i, row := range b.Rows {
        for j, val := range row {
            if val == num {
                b.Rows[i][j] = -1
                b.Columns[j][i] = -1
                if i == j {
                    b.Diagonal[i] = -1
                }
                marked = true
            }
        }
        if allMarked(row) {
            return true
        }
    }
    for _, column := range b.Columns {
        if allMarked(column) {
            return true
        }
    }
    if allMarked(b.Diagonal) {
        return true
    }
    return marked
}

func allMarked(nums []int) bool {
    for _, num := range nums {
        if num != -1 {
            return false
        }
    }
    return true
}

func (b *Board) Score(num int) int {
    sum := 0
    for _, row := range b.Rows {
        for _, num := range row {
            if num != -1 {
                sum += num
            }
        }
    }
    return sum * num
}

func main() {
    file, err := os.Open("day4_part2/input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var boards []*Board
    for scanner.Scan() {
        rows := make([][]int, 5)
        for i := range rows {
            scanner.Scan()
            line := scanner.Text()
            nums := strings.Split(line, " ")
            row := make([]int, 5)
            for j, numStr := range nums {
                num, _ := strconv.Atoi(numStr)
                row[j] = num
            }
            rows[i] = row
        }
        board := NewBoard(rows)
        boards = append(boards, board)
    }

    numsStr := "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"
    nums := strings.Split(numsStr, ",")
    for _, numStr := range nums {
        num, _ := strconv.Atoi(numStr)
        for _, board := range boards {
            if board.MarkNumber(num) {
                goto next
            }
        }
    next:
    }

    lastBoard := boards[0]
    for _, board := range boards {
        if board != lastBoard {
            for _, numStr := range
```
# Prompt 2
Q: Could you finish the program above and read the numbers for the bingo from the first line of the input instead of having them hardcoded?
A:
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Board struct {
    Rows     [][]int
    Columns  [][]int
    Diagonal []int
}

func NewBoard(rows [][]int) *Board {
    n := len(rows)
    columns := make([][]int, n)
    for i := range columns {
        columns[i] = make([]int, n)
    }
    for i, row := range rows {
        for j, num := range row {
            columns[j][i] = num
        }
    }
    diagonal := make([]int, n)
    for i := range rows {
        diagonal[i] = rows[i][i]
    }
    return &Board{
        Rows:     rows,
        Columns:  columns,
        Diagonal: diagonal,
    }
}

func (b *Board) MarkNumber(num int) bool {
    marked := false
    for i, row := range b.Rows {
        for j, val := range row {
            if val == num {
                b.Rows[i][j] = -1
                b.Columns[j][i] = -1
                if i == j {
                    b.Diagonal[i] = -1
                }
                marked = true
            }
        }
        if allMarked(row) {
            return true
        }
    }
    for _, column := range b.Columns {
        if allMarked(column) {
            return true
        }
    }
    if allMarked(b.Diagonal) {
        return true
    }
    return marked
}

func allMarked(nums []int) bool {
    for _, num := range nums {
        if num != -1 {
            return false
        }
    }
    return true
}

func (b *Board) Score(num int) int {
    sum := 0
    for _, row := range b.Rows {
        for _, num := range row {
            if num != -1 {
                sum += num
            }
        }
    }
    return sum * num
}

func main() {
    file, err := os.Open("day4_part2/input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var boards []*Board
    scanner.Scan()
    numsStr := scanner.Text()
    nums := strings.Split(numsStr, ",")
    for scanner.Scan() {
        rows := make([][]int, 5)
        for i := range rows {
            scanner.Scan()
            line := scanner.Text()
            nums := strings.Split(line, " ")
            row := make([]int, 5)
            for j, numStr := range nums {
                num, _ := strconv.Atoi(numStr)
                row[j] = num
            }
            rows[i] = row
        }
        board := NewBoard(rows)
        boards = append(boards, board)
    }

    for _, numStr := range nums {
        num, _ := strconv.Atoi(numStr)
        for _, board := range boards {
            if board.MarkNumber(num) {
                goto next
            }
        }
    next:
    }

    lastBoard := boards[0]
    for _, board := range boards {
        if board != lastBoard {
            lastBoard = board
        }
    }
    score := lastBoard.Score(0)
    fmt.Println("Final score:", score)
}
NOTE: we got error
```
panic: runtime error: index out of range [5] with length 5

goroutine 1 [running]:
main.main()
	.../advent2021/day4_part2_rev/main.go:110 +0x6b8
exit status 2
```
# Prompt 3
Q: Looking at the code, it seems like the program does not account for padding spaces infron of some of the boards. Could you fix that?
Example:
On line 3-5 we have
```
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
```
I: I would like to trim the empty spaces at the beginning of each line and replace double spaces with single spaces before to parse the given row.
A:
