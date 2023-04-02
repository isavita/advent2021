package main

import (
	"fmt"
)

func play(player1Pos, player2Pos, player1Score, player2Score, depth int) (int64, int64) {
	if player1Score >= 21 || player2Score >= 21 {
		if player1Score > player2Score {
			return 1 << depth, 0
		} else {
			return 0, 1 << depth
		}
	}

	player1Wins, player2Wins := int64(0), int64(0)

	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {
				rolls := roll1 + roll2 + roll3
				newPlayer1Pos := (player1Pos+rolls-1)%10 + 1
				newPlayer2Pos := (player2Pos+rolls-1)%10 + 1
				newPlayer1Score := player1Score + newPlayer1Pos
				newPlayer2Score := player2Score + newPlayer2Pos

				wins1, wins2 := play(newPlayer1Pos, newPlayer2Pos, newPlayer1Score, newPlayer2Score, depth+1)
				player1Wins += wins1
				player2Wins += wins2
			}
		}
	}

	return player1Wins, player2Wins
}

func main() {
	player1Pos, player2Pos := 4, 8 // Replace with your given starting positions

	player1Wins, player2Wins := play(player1Pos, player2Pos, player1Pos, player2Pos, 0)
	fmt.Printf("Player 1 wins in %d universes.\n", player1Wins)
	fmt.Printf("Player 2 wins in %d universes.\n", player2Wins)
}
