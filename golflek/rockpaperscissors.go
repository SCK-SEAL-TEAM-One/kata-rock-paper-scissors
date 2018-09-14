package rockpaperscissors

func RockPaperScissors(playerOne, playerTwo string) string {
	if playerOne == "rock" {
		return "PlayerOneWin"
	}
	return "PlayerTwoWin"
}
