package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//pour lire la saisie utilisateure
var reader = bufio.NewReader(os.Stdin)

//ReadGuess lis ce que l'utilisateur a marquer dans le terminal,
//guess renvoie le caractère ,renvoie une err si la saisie depasse 1 lettre
func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Quel est vôtre lettre? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		//enleve les espaces
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Saisie de lettre invalide, lettre =%s, longueur=%d\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}

func Quitter() (choix string, err error) {
	valid := false
	for !valid {
		fmt.Print("Est ce que vous voulez quitter ou vous voulez rejoué? O/N")
		choix, err = reader.ReadString('\n')
		if err != nil {
			return choix, err
		}
		choix = strings.TrimSpace(strings.ToUpper(choix))
		if choix == "O" || choix == "N" {
			valid = true
		}
	}
	return choix, err
}
