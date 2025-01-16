package entry

import "fmt"

func ExampleFromWordPairCsv() {
	p1 := "Apfel,apple"
	e1 := FromWordPairCsv(p1)

	fmt.Println(e1.De())
	fmt.Println(e1.En())

	// Output:
	// Apfel
	// apple
}
