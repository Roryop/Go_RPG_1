package text

import (
	"fmt"
	"strings"
	"time"
)

func Print(text string) {

	var splitText = strings.Split(text, "")

	for i := 0; i < len(splitText); i++ {
		fmt.Print(splitText[i])
		TextWait()
	}
	fmt.Println()
}

func Pause() {
	ShortWait()
	EmptyLine()
	ShortWait()
}

func Space(length int) {
	for i := 0; i < length; i++ {
		fmt.Println()
	}
}

func EmptyLine() {
	fmt.Println("-------------------------------------------------")
}

func LongWait() {
	time.Sleep(5 * time.Second)
}

func ShortWait() {
	time.Sleep(1 * time.Second)
}

func TextWait() {
	time.Sleep(100 * time.Millisecond)
}
