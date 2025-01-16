package entry

import "strings"

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {
	// Hinweis: Teilen Sie den String pair in zwei Wörter auf.
	// Dafür können Sie die Funktion strings.Split() verwenden.
	//solution:begin
	words := strings.Split(pair, ",")
	if len(words) == 2 {
		return New(words[0], words[1])
	}
	//solution:end
	return Empty()
}
