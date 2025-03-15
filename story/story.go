package story

import (
	"fmt"
	"start/enemy"
	"start/gear"
	"start/player"
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

// Gets Player, Inventory, current stats, world, world_barrier, typ of enemy
// Creates Enemy, Fights against enemy; Updates player Stats
// Returns enemyStats, new current Stats
func Fight(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int, typ int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	///////// REQUIREMENT FOR UPGRADING WORLD BARRIER ///////////////////

	var barrierRequirement = 0

	////////////////////////////////////////////////////////////////////

	var enemyName, enemyStats = enemy.CreateEnemy(world, world_barrier, typ)

	fmt.Println("Du fightest einen", enemyName+"!!!")
	fmt.Println("Er ist Level", enemyStats[0], "!!!")
	fmt.Println("Er hat", enemyStats[1], "HP!")
	fmt.Println("Du hast", hp, "HP!")

	for enemyStats[1] > 0 && hp > 0 {
		var choice int

		//Anfrage wegen Angriff
		fmt.Println("Möchtest du angreifen?")
		fmt.Println("1: Ja")
		fmt.Println("2: Nein")
		fmt.Println("3: Eigene Stats sehen")
		fmt.Println("4: Inventory sehen")
		fmt.Println("5: End Game")
		fmt.Scanln(&choice)

		//Angreifen
		switch choice {
		case 1:
			//Player attacks Enemy
			enemyStats[1] = enemyStats[1] - ((att * 100) / (100 + enemyStats[3]))
			fmt.Println("Der", enemyName, "hat noch", enemyStats[1], "HP")
			//Enemy attacks Player
			if enemyStats[1] > 0 {
				hp = hp - ((enemyStats[2] * 100) / (100 + def))
			}
			fmt.Println("Du hast noch", hp, "HP")

		case 2:
			fmt.Println("Du hast nicht angegriffen")
			//Enemy attacks Player
			hp = hp - ((enemyStats[2] * 100) / (100 + def))
			fmt.Println("Du hast noch", hp, "HP")

		case 3:
			player1.SeePlayerStats(inventory, hp, att, def, rec)
		case 4:
			gear.GiveInventoryInformation(inventory)
		case 5:
			goto end
		}
	}
	///////////////////////// Case Enemy Died ///////////////////////////
	if enemyStats[1] <= 0 {

		////////////// Setting World Barrier Upgrade Requirement ////////////
		if barrierRequirement >= 9 {
			world_barrier += 1
		}

		//////////////// Enemy Item Drop /////////////
		inventory = gear.AddDropToInventory(inventory, world_barrier)

		///////////////// Player Management ////////////////
		player1.Exp_Function(enemyStats)                                           // Giving Exp to Player
		hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) // Player will be healed with levelUp + Updating Stats + Updating current Stats

	}

	//////////////// Clearing Inventory on Player Death ////////////////
	if hp <= 0 {
		inventory = gear.FillEmptyInventory(inventory)
	}
end:
	return player1, inventory, hp, att, def, rec, world_barrier
}
