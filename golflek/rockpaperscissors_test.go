package rockpaperscissors_test

import (
	. "rockpaperscissors"
	"testing"
)

func Test_RockPaperScissors_Input_PlayerOne_Rock_PlayerTwo_Scissor_Should_Be_PlayerOneWin(t *testing.T) {
	expected := "PlayerOneWin"
	playerOne := "rock"
	playerTwo := "scissors"

	actual := RockPaperScissors(playerOne, playerTwo)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}
func Test_RockPaperScissors_Input_PlayerOne_scissors_PlayerTwo_rock_Should_Be_PlayerTwoWin(t *testing.T) {
	expected := "PlayerTwoWin"
	playerOne := "scissors"
	playerTwo := "rock"

	actual := RockPaperScissors(playerOne, playerTwo)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}
