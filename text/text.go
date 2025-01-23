package text

import (
	"fmt"
	"strings"
	"time"
)

//////////////////////// Custom Print Functions ///////////////////////

// Gets Text as String
// Gives Out Text character after character
// Returns Nothing
func Print(text string) {

	var splitText = strings.Split(text, "")

	for i := 0; i < len(splitText); i++ {
		fmt.Print(splitText[i])
		TextWait()
	}
	fmt.Println()
}

// Gets Nothing
// Prints a Line with a Short Wait before and After
// Returns Nothing
func Pause() {
	ShortWait()
	EmptyLine()
	ShortWait()
}

// Gets length of empty Space
// Prints an empty Line * length
// Returns Nothing
func Space(length int) {
	for i := 0; i < length; i++ {
		fmt.Println()
	}
}

// Gets Nothing
// Prints a Line
// Returns Nothing
func EmptyLine() {
	fmt.Println("-------------------------------------------------")
}

//////////////////////// Custom Wait Functions //////////////////////////

// Gets Nothing
// Sleeps for 3 Seconds
// Returns Nothing
func LongWait() {
	time.Sleep(3 * time.Second)
}

// Gets Nothing
// Sleeps for 1.5 Seconds
// Returns Nothing
func ShortWait() {
	time.Sleep(3 * time.Second / 2)
}

// Gets Nothing
// Sleeps for 75 Milliseconds
// Returns Nothing
func TextWait() {
	time.Sleep(75 * time.Millisecond)
}
