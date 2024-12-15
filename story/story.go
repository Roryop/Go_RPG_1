package story

import (
	"fmt"
	"strings"
	"time"
)

func Prologue() {
	EmptyLine()
	Print("The story begins")
	EmptyLine()
	LongWait()
	Print("Are you ready, Adventurer?")
	Pause()
	Print("What awaits you is a world filled with dangers at every corner.")
	ShortWait()
	Space(10)
	LongWait()
	Print("Try to endlessly grow Stronger, Strong enough to kill the Demons of this World!")
	LongWait()
	Pause()
	Space(20)
}

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
