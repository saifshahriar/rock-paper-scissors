package main

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMessage string
	label        string
}

// Interactive terminal prompt
func promptGetSelect(pc promptContent) string {
	items := []string{"Rock \U0001faa8", "Paper 📄", "Scissors ✂", "Exit \U0001f6ab"}

	var result string
	var err error

	prompt := promptui.Select{
		Label: pc.label,
		Items: items,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | magenta }}",
			Active:   "{{ . | green }}",
			Inactive: "{{ . | cyan }}",
			Selected: "{{ . | red | bold }}",
		},
	}
	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(0)
	}
	if result == "Exit \U0001f6ab" {
		fmt.Println("Exited")
		os.Exit(0)
	}
	return result
}

// Player selects rock, paper, or scissors (in promptGetSelect())
func playerSelect() string {
	pc := promptContent{
		errorMessage: "Please select rock, paper, or scissors",
		label:        "Select",
	}
	return promptGetSelect(pc)
}

// Computer randomly selects rock, paper, or scissors
func computerSelect() string {
	items := []string{"Rock \U0001faa8", "Paper 📄", "Scissors ✂"}
	return items[rand.Intn(len(items))]
}

// Declare result
func result(player, computer string) string {
	switch {
	case player == computer:
		return "Tie! 🤞🏼"
	case player == "Rock \U0001faa8":
		if computer == "Paper 📄" {
			return "You lose. 👎🏼"
		}
		return "You Win! ✌🏼"
	case player == "Paper 📄":
		if computer == "Scissors ✂" {
			return "You lose. 👎🏼"
		}
		return "You Win! ✌🏼"
	case player == "Scissors ✂":
		if computer == "Rock \U0001faa8" {
			return "You lose. 👎🏼"
		}
		return "You Win! ✌🏼"
	}
	return ""
}

func main() {
	fmt.Println("*********************************\n*				*\n*				*\n*	Rock, Paper, Scissors	*\n*				*\n*				*\n*********************************\n\n")
	pc := promptContent{
		label: "Please select either Rock, Paper or Scissors.",
	}
	for i := 0; i < 100; i++ {
		playerChoice := promptGetSelect(pc)
		computerChoice := computerSelect()
		result := result(playerChoice, computerChoice)

		fmt.Println("\n\n\n\n\n\n\n")
		fmt.Println("Player 🧑🏼 =", playerChoice)
		fmt.Println("Computer 💻 =", computerChoice)
		fmt.Println(result)
		fmt.Println("\n\n\n\n\n\n\n")
	}
}