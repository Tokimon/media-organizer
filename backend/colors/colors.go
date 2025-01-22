package colors

import "fmt"

// RESET escape code
const RESET = "\033[0m"

// Red applies the red color to a string
func Red(str string) string {
	return fmt.Sprintf("\033[31m%s%s", str, RESET)
}

// Green applies the green color to a string
func Green(str string) string {
	return fmt.Sprintf("\033[32m%s%s", str, RESET)
}

// Yellow applies the yellow color to a string
func Yellow(str string) string {
	return fmt.Sprintf("\033[33m%s%s", str, RESET)
}

// Blue applies the blue color to a string
func Blue(str string) string {
	return fmt.Sprintf("\033[34m%s%s", str, RESET)
}

// Purple applies the purple color to a string
func Purple(str string) string {
	return fmt.Sprintf("\033[35m%s%s", str, RESET)
}

// Cyan applies the cyan color to a string
func Cyan(str string) string {
	return fmt.Sprintf("\033[36m%s%s", str, RESET)
}

// Gray applies the gray color to a string
func Gray(str string) string {
	return fmt.Sprintf("\033[38;2;125;125;125m%s%s", str, RESET)
}

// White applies the white color to a string
func White(str string) string {
	return fmt.Sprintf("\033[97m%s%s", str, RESET)
}

func Strikethrough(str string) string {
	return fmt.Sprintf("\x1b[9m%s%s\x1b[29m", str, RESET)
}
