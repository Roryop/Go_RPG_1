package story

import (
	"fmt"
	"start/enemy"
	"start/gear"
	"start/player"
	"start/text"
)

func Robbery(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) ([4]int, int, int, int, int) {
	text.Print("Du wirst von einem hinterhältigen Räuber angeriffen")

	var enemyName, enemyStats = enemy.CreateEnemy(world, world_barrier, 4)

	fmt.Println("Du fightest einen", enemyName+"!!!")
	fmt.Println("Er ist Level", enemyStats[0], "!!!")
	fmt.Println("Er hat", enemyStats[1], "HP!")
	fmt.Println("Du hast", hp, "HP!")

	for enemyStats[1] > 0 && hp > 0 {
		var choice = 1

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
end:
	return enemyStats, hp, att, def, rec
}
