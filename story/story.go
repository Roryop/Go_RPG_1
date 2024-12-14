package story

import (
	"fmt"
	"time"
)

func Prologue() {
	EmptyLine()
	fmt.Println("The story begins")
	EmptyLine()
	LongWait()
	fmt.Println("Are you ready, Adventurer?")
	ShortWait()
	EmptyLine()
	ShortWait()
	fmt.Println("What awaits you is a world filled with dangers at every corner.")
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
