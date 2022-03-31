package hangman

import (
	"fmt"
	"os"
	"strings"
)

type Game struct {
	State        string   //Etat du jeu
	Letters      []string //Lettre dans le mot a trouver
	FoundLetters []string //Lettre trouver
	UsedLetters  []string //Lettre Utilisé
	TurnsLeft    int      //Essaie restant
}

//New crée la partie,
//prend en valeur un int et un string,
//renvoi un pointeur Game et une erreur
func New(_turn int, _word string) (*Game, error) {
	if len(_word) < 3 {
		return nil, fmt.Errorf("le mot '%s' doit avoir 3 caractères, mais il a=%v/", _word, len(_word))
	}

	letters := strings.Split(strings.ToUpper(_word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_ "
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    _turn,
	}
	return g, nil
}

//MakeAGuess gere la proposition de l'utilisateur,
//il a un pointeur receiver de Game,
//prend en parametre un string de la lettre trouve
func (g *Game) MakeAGuess(_guess string) {
	_guess = strings.ToUpper(_guess)

	switch g.State {
	case "won", "lost":
		os.Exit(0)
	}

	if letterInWord(_guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(_guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(_guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	}

	if !letterInWord(_guess, g.UsedLetters) {
		g.State = "badGuess"
		g.LoseTurn(_guess)
		if g.TurnsLeft == 0 {
			g.State = "lost"
		}
	}
}

//RevealLetters revele les lettres trouvé par l'utilisateur,
//il a pointeur receiver de Game,
//prend en parametre un string des lettre revele
func (g *Game) RevealLetter(_guess string) {
	g.UsedLetters = append(g.UsedLetters, _guess)
	for i, l := range g.Letters {
		if l == _guess {
			g.FoundLetters[i] = _guess
		}
	}
}

//LoseTurn fonction a faire si on a trouver la mauvaise lettre,
//a pointeur receiver de Game
//prend en parametre un string
func (g *Game) LoseTurn(_guess string) {
	g.UsedLetters = append(g.UsedLetters, _guess)
	g.TurnsLeft--
}

//hasWon verifie si on a gagné
//prend en parametre un slice de letre et slice de lettre trouver,
//renvoi un bool
func hasWon(_letters []string, _foundLetteres []string) bool {
	for i := range _letters {
		if _letters[i] != _foundLetteres[i] {
			return false
		}
	}
	return true
}

//LetterInWord si c'est la bonne lettre renvoie true sinon renvoie false,
//prend en parametre un string qui est la lettre et le slice de lettre sur lequel on verifie,
//renvoie un bool
func letterInWord(_guess string, _letters []string) bool {
	for _, l := range _letters {
		if l == _guess {
			return true
		}
	}
	return false
}
