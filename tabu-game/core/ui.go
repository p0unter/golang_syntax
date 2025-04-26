package core

import "fmt"

const authorName string = "Pounter"
const madeTime string = "26/04/2025"

var gameOptions = []string{
	"1\tPerson vs Person",
	"2\tPerson vs AI",
}

func GetMainUI() {
	fmt.Printf("- Tabu -\n%v presents. %v Allright reserved\n", authorName, madeTime)
}

func GetOptions() {
	fmt.Printf("\n- - - - - - -\n")
	fmt.Printf("Options (by Number)\n")
	var content string
	for index, option := range gameOptions {
		content += option
		if index != len(gameOptions) {
			content += "\n"
		}
	}
	fmt.Printf("%v- - - - - - -\n", content)
}
