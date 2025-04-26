package core

import (
	"fmt"
)

func GetOptionInput() string {
	var resInput string
	fmt.Scanln(&resInput)

	return resInput
}

func CheckInput() {
	byInput := GetOptionInput()
	switch byInput {
	case "1":
		PersonersVS()
	case "2":
		RobotVS()
	}
}
