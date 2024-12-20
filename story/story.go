package story

import (
	"start/text"
)

func Prologue() {
	text.EmptyLine()
	text.Print("The Journey begins")
	text.EmptyLine()
	text.LongWait()
	text.Print("Are you ready, Adventurer?")
	text.Pause()
	text.Print("What awaits you is a world filled with dangers at every corner.")
	text.ShortWait()
	text.Space(20)
	text.LongWait()
	text.Print("Try to endlessly grow Stronger, Strong enough to kill the Demons of this World!")
	text.LongWait()
	text.Pause()
	text.Space(30)
}
