package story

import (
	"fmt"
	"start/gear"
	"start/player"
	"start/text"
)

///////////////////7//////////////////////////////////////////////////////Cyberpunk Welt////////////////////////////////////////////////////////////////////////////

//////////////////Funktionen für Slums///////////////////

// Gets player, inventory, current stats, world, world_barrier
// Executes Event Robbery
// Returns player, inventory, current  stats, world_barrier
func Robbery(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du wirst von einem Räuber überfallen!!!")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 4) ////////////einfacher Kampf (Zahl am Ende gibt den enemy aus enemy.go an)

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach switch case
func Bettler(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var giveaway int
	text.Print("Du triffst auf einen Bettler. Willst du ihm etwas geben? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&giveaway) //////////Entscheidung des Spielers ob er eas abgibt oder nicht

	switch giveaway { ////////////////////////////////////if Abfrage würde auch gehen, aber switch cases find ich schöner + man kann sie einfacher ausbauen
	case 1: //////Er gibt etwas ab -> Der Penner dankt ihm und der Spieler erlangt Karma
		text.Print("Du gibst dem Penner eins deiner Items und er dankt dir.")
		gear.SubtractFromInventory(inventory)

		player1.UpdateKarma(7)
	default:
		text.Print("Der Obdachlose wird sauer und greift dich an!") //////////////Kampfalgorithmus im Falle Nein
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 18)
	}

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets Playerstats
// Returns inventory
func Muelltonne(inventory [10]*gear.InventorySlot, player_level int) [10]*gear.InventorySlot {
	var choice int
	text.Print("Du läufst an einer Mülltonne vorbei, willst du sie durchsuchen? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice) /////////////Wahlvariable

	switch choice {
	case 1: ///////////////In der Tonne liegt ein Item
		text.Print("Du wühlst in einer Mülltonne herum. Denk mal drüber nach!")
		var item = gear.ItemDrop(player_level) ////////////Funktion aus gear, die random items droppt
		text.Print("Du hast ein Item gefunden! Es ist ein ")
		fmt.Print(item)
		gear.AddToInventory(inventory, item) ////////////Funktion aus gear, die das gedroppte Item im Inventar aufnimmt

	default: ////////////////Nix passiert
		text.Print("OK dann eben nicht")

	}
	return inventory
}

// ////////////////Funktionen für Businessviertel///////////////////

// Gets Playerstats
// Returns Playerstats nach veränderung des karmas im switch case
func Wallet(player1 *player.Player, inventory [10]*gear.InventorySlot, player_level int) ([10]*gear.InventorySlot, int) { ///////Der Funktion übergebene Werte
	var choice int

	text.Print("Du findest ein Portemonnaie. Darin ist ein Ausweis. Willst du de Besitzer suchen? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice) ///////////// 1 oder 2 als Variablenwert für den Switch Case in der Funktion für mehrere Aktionspfade innerhalb des events

	switch choice {

	case 1:
		fmt.Println("Du findest den Besitzer und er dankt dir. Das war gut für dein Karma")
		player1.UpdateKarma(5) ///////////////////////////////////////Entscheidung wirkt sich gut auf den Karma-Wert des Spielers aus

	default:
		fmt.Println("Jokes on you, du bist ein schlechter Mensch und die Geldbörse ist leer.")
		player1.UpdateKarma(-5) //////////////////////////////////////Entscheidung wirkt sich gut auf den Karma-Wert des Spielers aus
	}
	return inventory, player_level
}

// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach switch case
func Security(inventory [10]*gear.InventorySlot) [10]*gear.InventorySlot {
	text.Print("Du gerätst in eine Kontrolle der High Class Security.")
	text.Print("Sie beschlagnamen ein Teil deiner Ausrüstung. Wähle welche.")

	inventory = gear.SubtractFromInventory(inventory) //////////////////Entfernt in stück gear aus dem Inventar

	return inventory
}

// Gets Playerstats, Kampfalgoritmus
// Returns Playerstats nach switch case
func Businessman(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {
	var choice int
	text.Print("Du rämpelst versehentlich einen Businessman an und sein Handy fiel runter.")
	text.Print("Er will Geld als Entschädigung, aber  du hast keins. /Kämpfen /Weglaufen?")

	fmt.Println("1: Kämpfen")
	fmt.Println("2: Weglaufen")
	fmt.Scanln(&choice) ////////////////////Entscheidungsvariable ob kämpfen oder

	switch choice {
	case 1:
		player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 19) //////Kampfalgorithmus gegen Businessman

	default:
		fmt.Println("Du kommst ohne Geld und ohne Würde davon.") /////////////Nix passiert
	}
	return player1, inventory, hp, att, def, rec, player_level
}

//////////////////Funktionen für Außenstadt///////////////////

// Gets Player, inventory, current stats, world, player_level
// Fights a Bulle
// Returns Player, inventory, current stats, player_level
func Grenzkontrolle(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	text.Print("Du gerätst in eine Grenzkontrolle. Der Polizist ist schlecht gelaunt und greift Dich an")

	player1, inventory, hp, att, def, rec, player_level = Fight(player1, inventory, hp, att, def, rec, world, player_level, 20) /////////////////Kampfalgorithmus mit enemy 20 (Bulle)

	return player1, inventory, hp, att, def, rec, player_level
}

// Gets Player, inventory, hp
// player heals by 1/5th
// Returns hp
func Motel(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {

	var choice int
	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du stößt auf ein Motel. Willst du dort rasten und ein Paar HP zurück gewinnen? /ja /nein?")
	fmt.Println("1: Ja")
	fmt.Println("2: Nein")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		text.Print("Gute Nacht!")

		for i := 0; i < 3; i++ {
			text.LongWait() ///////////////Spieler muss 9 sek warten
		}

		hp = hp + (maxHp / 5) ////////////////////////Spieler erhält mehr hp und muss warten
		if hp > maxHp {
			hp = maxHp
		}

	default:
		text.Print("Hoffentlich wird es dir nicht zum Verhängnis. Schlaf ist wichtig!")

	}
	return hp

}

// Gets Player, inventory, hp
// Reduces hp by 1/5th
// Returns hp
func Verlaufen(player1 *player.Player, inventory [10]*gear.InventorySlot, hp int) int {

	var maxHp, _, _, _ = player1.CreateStats(inventory)

	text.Print("Du verläufst dich om einem ehemaligen Industriegebiet und brauchst zu lange, um wieder heraus zu finden. --HP")

	hp = hp - (maxHp / 5) ////////////////////////Spieler bekommt HP abgezogen
	if hp > maxHp {
		hp = maxHp
	}

	return hp
}
