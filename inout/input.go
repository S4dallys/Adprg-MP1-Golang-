package inout

import "fmt"

func Get_Int(char * int) (error) {
	fmt.Print("/> ")
	_, err := fmt.Scanln(char)

	if (err != nil) {
		var discard string
		fmt.Scanln(&discard)
		Print_Error(INVALID_INPUT_ERROR)
	}

	return err
}

func Get_String(char * string) (error) {
	fmt.Print("/> ")
	_, err := fmt.Scanln(char)

	if (err != nil) {
		var discard string
		fmt.Scanln(&discard)
		Print_Error(INVALID_INPUT_ERROR)
	}
	return err
}

func Get_Float(char * float64) (error) {
	fmt.Print("/> ")
	_, err := fmt.Scanln(char)

	if (err != nil) {
		var discard string
		fmt.Scanln(&discard)
		Print_Error(INVALID_INPUT_ERROR)
	}

	return err
}