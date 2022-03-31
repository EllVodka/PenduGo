package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Mots %s contien la lettre %s. possede=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Mots %s ne doit pas contenir la lettre %s. possede=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Erreur devrai renvoyer un erreur quand on fait une chaine vide")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)

}

func TestBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("z")
	validState(t, "badGuess", g.State)

}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	validState(t, "won", g.State)

}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("v")
	g.MakeAGuess("t")
	validState(t, "lost", g.State)

}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("L'état devrai être %s, mais il est= %v", expectedState, actualState)
		return false
	}
	return true
}
