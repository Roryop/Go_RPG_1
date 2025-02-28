package story

import (
	"start/text"
)

func Prologue() {
	text.EmptyLine()
	text.Print("Deine Reise beginnt.")
	text.EmptyLine()
	text.LongWait()
	text.Print("Bist du bereit, Abenteurer?")
	text.Pause()
	text.Print("Was dich erwartet ist ein Multiversum,")
	text.Print("gefüllt mit Gefahren, die an jeder Ecke lauern.")
	text.ShortWait()
	text.Space(20)
	text.LongWait()
	text.Print("Versuche immer und immer Stärker zu werden,")
	text.Print("Stark genug um die Dämonen dieser Welt zu vernichten!")
	text.LongWait()
	text.Pause()
	text.Space(30)
}
