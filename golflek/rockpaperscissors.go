package rockpaperscissors

func RockPaperScissors(playerOne, playerTwo string) string {
	if playerOne == "rock" && playerTwo == "scissors" {
		return "PlayerOneWin"
	}
	if playerOne == "rock" && playerTwo == "paper" {
		return "PlayerTwoWin"
	}
	if playerOne == "rock" && playerTwo == "rock" {
		return "Draw"
	}

	if playerOne == "scissors" && playerTwo == "paper" {
		return "PlayerOneWin"
	}
	if playerOne == "scissors" && playerTwo == "rock" {
		return "PlayerTwoWin"
	}
	if playerOne == "scissors" && playerTwo == "scissors" {
		return "Draw"
	}

	if playerOne == "paper" && playerTwo == "rock" {
		return "PlayerOneWin"
	}
	if playerOne == "paper" && playerTwo == "scissors" {
		return "PlayerTwoWin"
	}
	if playerOne == "paper" && playerTwo == "paper" {
		return "Draw"
	}

	return "PlayerTwoWin"
}
