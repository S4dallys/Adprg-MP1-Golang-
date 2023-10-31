package inout

import "fmt"

const (
	ALREADY_EXISTS_ERROR  uint8 = 0
	INVALID_INPUT_ERROR         = 1
)

func Print_Error(error uint8) {
	fmt.Println()
	switch error {
		case ALREADY_EXISTS_ERROR:
			fmt.Println("Slot chosen already taken.")
		case INVALID_INPUT_ERROR:
			fmt.Println("Invalid input.")
	}
	fmt.Println()
}
