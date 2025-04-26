package main

import "tabu/core"

func main() {
	core.GetMainUI()
	core.GetOptions()

	for {
		core.CheckInput()
	}
}
