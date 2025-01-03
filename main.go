package main

import (
	"fmt"
	"start/enemy"
	"start/game"
	"start/gear"
	"start/player"
	"start/story"
)

/*
	func helloWelt() (string, string) {

		var x, y string
		x, y = "Hello", "World"
		return x, y

	}
*/

func main() {

	story.Prologue()
	var b = gear.NewGear("Accessoire", "Lesser Ring of Strength")
	fmt.Println(b)
	/////////////Testing Weapon///////////
	var inventory = gear.NewInventory()
	fmt.Println(inventory)

	var inventorySlot = gear.NewInventorySlot()
	inventorySlot.InputInventorySlot(b, 1)
	inventory[0] = inventorySlot
	fmt.Println(inventory[1])
	gear.GiveInventoryInformation(inventory)

	////////////////////////End Weapon Test///////////////////////

	/*
		var x, y string
		x, y = helloWelt()

		fmt.Println(x, y)
	*/

	//kreiere Spieler

	var player1 = player.BeginPlayer()
	fmt.Println("Player: ", player1)
	player1.SeePlayerStats()

	var hp, att, def, rec = player1.CreateStats(inventory)

	fmt.Println("Recovery:", rec)

	//das Game_Level auswählen
	var choice int = 0
	for choice != 3 {

		player1.SetStats()
		player1.SetStatsAccessoires(inventory)
		player1.SeePlayerStats()

		hp = player1.GetStat(2) //Healing Player before going into Level
		var game_level = game.ChooseGameLevel()

		for i := 0; i < 10 && hp > 0; i++ {
			var enemyName, enemyStats = enemy.CreateEnemy(game_level)

			fmt.Println("Du fightest einen", enemyName+"!!!")
			fmt.Println("Er ist Level", enemyStats[0], "!!!")
			fmt.Println("Er hat", enemyStats[1], "HP!")
			fmt.Println("Du hast", hp, "HP!")

			//Fight machen
			for enemyStats[1] > 0 && hp > 0 {

				//Anfrage wegen Angriff
				fmt.Println("Möchtest du angreifen?")
				fmt.Println("1: Ja")
				fmt.Println("2: Nein")
				fmt.Println("3: End Game")
				fmt.Scanln(&choice)
				fmt.Println(choice)

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

				default:
					goto end
				}
			}
			if enemyStats[1] <= 0 {
				player1.CalculateExp((enemyStats[1] + enemyStats[2] + enemyStats[3] + enemyStats[0]) / 4)
				hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) //player will be healed with LevelUp
				//Spieler will not be healed after fight
				fmt.Println("Deine Stats sind jetzt:")
				fmt.Println("HP:", hp)
				fmt.Println("Att:", att)
				fmt.Println("Def:", def)
			}
		}
	end:
		if choice == 3 {
			break
		}
	}
}
