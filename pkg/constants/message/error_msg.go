package general

import "fmt"

//* Body parser
const ErrBodyParser = "Please check your input!"
const ErrBodyParserInd = "Tolong periksa input anda!"

//* Integer
func ErrValueID(field string) string {
	return fmt.Sprintf("Make sure '%s' valid!", field)
}
func ErrValueIDInd(field string) string {
	return fmt.Sprintf("Pastikan value dari '%s' adalah benar!", field)
}

func ErrValuePageInt(field string) string {
	return fmt.Sprintf("Make sure '%s' is integer!", field)
}
func ErrValuePageIntInd(field string) string {
	return fmt.Sprintf("Pastikan value dari '%s' adalah bilangan bulat!", field)
}
