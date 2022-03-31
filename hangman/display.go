package hangman

import (
	"fmt"
	"strings"
)

//DrawWelcome affiche hangman en gros
func DrawWelcome() {
	fmt.Println(`
	
	_       _    _                   _             _              _   _         _                   _          
	/ /\    / /\ / /\                /\ \     _    /\ \           /\_\/\_\ _    / /\                /\ \     _  
   / / /   / / // /  \              /  \ \   /\_\ /  \ \         / / / / //\_\ / /  \              /  \ \   /\_\
  / /_/   / / // / /\ \            / /\ \ \_/ / // /\ \_\       /\ \/ \ \/ / // / /\ \            / /\ \ \_/ / /
 / /\ \__/ / // / /\ \ \          / / /\ \___/ // / /\/_/      /  \____\__/ // / /\ \ \          / / /\ \___/ / 
/ /\ \___\/ // / /  \ \ \        / / /  \/____// / / ______   / /\/________// / /  \ \ \        / / /  \/____/  
/ / /\/___/ // / /___/ /\ \      / / /    / / // / / /\_____\ / / /\/_// / // / /___/ /\ \      / / /    / / /   
/ / /   / / // / /_____/ /\ \    / / /    / / // / /  \/____ // / /    / / // / /_____/ /\ \    / / /    / / /    
/ / /   / / // /_________/\ \ \  / / /    / / // / /_____/ / // / /    / / // /_________/\ \ \  / / /    / / /     
/ / /   / / // / /_       __\ \_\/ / /    / / // / /______\/ / \/_/    / / // / /_       __\ \_\/ / /    / / /      
\/_/    \/_/ \_\___\     /____/_/\/_/     \/_/ \/___________/          \/_/ \_\___\     /____/_/\/_/     \/_/       
																												

	`)
}

//Draw affiche le pendu et l'etat de la parti,
//prend en parametre un pointeur receiver Game et un string
func Draw(_g *Game, _guess string) {
	drawTurns(_g.TurnsLeft)
	drawState(_g, _guess)

}

//drawTurns desine le pendu,
//prend en parametre le tour en int
func drawTurns(_l int) {
	var draw string
	switch _l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

//drawState affiche la lettre trouve & utilisé et affiche l'etat de la partie,
//prend en parametre un pointeur receiverGame et un string
func drawState(_g *Game, _guess string) {

	fmt.Printf("\nIl vous reste %d d'essaie\n", _g.TurnsLeft)
	fmt.Printf("Vous avez fait %d de tours\n\n", len(_g.UsedLetters))

	fmt.Print("Trouvé : ")
	DrawLetters(_g.FoundLetters)

	fmt.Print("Utilisé : ")
	DrawLetters(_g.UsedLetters)

	switch _g.State {
	case "goodGuess":
		fmt.Print("Bien vu\n")
	case "alreadyGuessed":
		fmt.Printf("La lettre %s tu la déjà utilisé ma gueule\n", strings.ToUpper(_guess))
	case "badGuess":
		fmt.Printf("T'es mauvais, la lettre %s est même pas dans le mot \n", strings.ToUpper(_guess))
	case "lost":
		fmt.Printf("Ta perdu, sah je le savais, prend pas trop le seum bg :)! Le mot était : ")
		DrawLetters(_g.Letters)
	case "won":
		fmt.Printf("Hisashiburi dana Mugiwara, Bravo ta gagner tu vien de loin bg! Le mot est : ")
		DrawLetters(_g.Letters)

	}
}

//DrawLetters affiche la lettre trouver,
//prend en parametre un string
func DrawLetters(_l []string) {
	for _, c := range _l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
