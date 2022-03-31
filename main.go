package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"training.go/Hangman/hangman"
	"training.go/Hangman/hangman/dictionary"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Le dictioniare n'a pas reussi a charger: %v\n", err)
		os.Exit(1)
	}

	g, _ := hangman.New(8, dictionary.PickWord())

	guess := ""
	for {
		CallClear()
		if guess == "" {
			hangman.DrawWelcome()
		}

		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			reponse, err := hangman.Quitter()
			if err != nil {
				fmt.Printf("Votre reponse cause une erreur: %v\n", err)
			}
			switch reponse {
			case "O":
				main()
			case "N":
				os.Exit(0)
			}
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("y'a une erreur dans le terminal : %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
