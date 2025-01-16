package dict

import (
	"dictionary/entry"
	"os"
	"strings"
)

// FromWordPairsCsv erzeugt ein neues Wörterbuch aus einer Liste von Wortpaaren.
// Ein Wortpaar ist dabei jeweils gegeben als ein String mit zwei Wörtern,
// getrennt durch ein Komma.
// Gibt es in einem Wortpaar kein oder mehrere Kommas, wird das Wortpaar ignoriert.
func FromWordPairsCsv(words []string) Dict {
	// Hinweis: Laufen Sie über words und erzeugen Sie für jedes Wortpaar ein Entry-Objekt.
	// Verwenden Sie dabei den Konstruktor entry.FromWordPairCsv().
	d := New()
	//solution:begin
	for _, w := range words {
		e := entry.FromWordPairCsv(w)
		if e.IsValid() {
			d.Add(e)
		}
	}
	//solution:end
	return d
}

// ReadFileLines liest den Inhalt einer Datei und liefert die Zeilen als String-Slice.
// Wenn die Datei nicht gelesen werden kann, wird eine leere Liste zurückgegeben.
func ReadFileLines(filename string) []string {
	// Hinweis: Lesen Sie den Inhalt der Datei mit os.ReadFile() ein.
	// Beachten Sie, dass Sie dabei eine Fehlerbehandlung benötigen.
	// Die Funktion os.ReadFile() liefert zwei Werte zurück.
	// Verwenden Sie strings.Split(), um den Inhalt der Datei in Zeilen zu teilen.
	//solution:begin
	file, err := os.ReadFile(filename)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(file), "\n")
	//solution:end
}

// FromFile erzeugt ein neues Wörterbuch aus dem Inhalt einer Datei.
// Die Datei enthält eine Liste von Wortpaaren, wobei jedes Wortpaar in einer eigenen Zeile steht.
// Ein Wortpaar ist dabei jeweils gegeben als ein String mit zwei Wörtern,
// getrennt durch ein Komma.
// Gibt es in einem Wortpaar kein oder mehrere Kommas, wird das Wortpaar ignoriert.
// Wenn die Datei nicht gelesen werden kann, wird ein leeres Wörterbuch zurückgegeben.
func FromFile(filename string) Dict {
	// Hinweis: Lesen Sie den Inhalt der Datei mit ReadFileLines() ein.
	// Verwenden Sie dann FromWordPairsCsv(), um ein Wörterbuch zu erzeugen.
	//solution:begin
	lines := ReadFileLines(filename)
	return FromWordPairsCsv(lines)
	//solution:end
}
