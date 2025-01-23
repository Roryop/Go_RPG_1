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

	///////////////////////////////Test Environment//////////////////////////
	for i := 0; i < 20; i++ {
		var a = game.ItemDrop()
		fmt.Println(a)
	}

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

	////////////////////////End Test Environment///////////////////////

	//////////////////////////// Creating Player ////////////////////////////////

	var player1 = player.BeginPlayer() // Creating Player
	fmt.Println("Player: ", player1)   // Giving Out Player
	player1.SeePlayerStats()           // Giving out PlayerStats

	var hp, att, def, rec = player1.CreateStats(inventory) // Creating current Player Stats

	fmt.Println("Recovery:", rec) // Test

	////////////////////////////////// Game /////////////////////////////////////

	var choice int = 0 // Creating Variable so Player can later end game themselves
	for choice != 3 {  // Game keeps running until Player end it

		///////////////////// Setting Stats before each Run /////////////////////
		player1.SetStats()
		player1.SetStatsAccessoires(inventory)
		player1.SeePlayerStats()

		hp = player1.GetStat(2) // Healing Player before going into Level

		var game_level = game.ChooseGameLevel()

		for i := 0; i < 10 && hp > 0; i++ { // Entering Fights until Player 1. killed 10 monster; 2. is dead

			////////////////////// Setting Up Enemy ////////////////////////
			var enemyName, enemyStats = enemy.CreateEnemy(game_level)

			fmt.Println("Du fightest einen", enemyName+"!!!")
			fmt.Println("Er ist Level", enemyStats[0], "!!!")
			fmt.Println("Er hat", enemyStats[1], "HP!")
			fmt.Println("Du hast", hp, "HP!")

			////////////////////////// Fighting //////////////////////////////
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
			///////////////////////// Case Enemy Died ///////////////////////////
			if enemyStats[1] <= 0 {
				player1.Exp_Function(enemyStats)                                           // Giving Exp to Player
				hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) // Player will be healed with levelUp + Updating Stats + Updating current Stats
				// Player will not be healed after fight
				fmt.Println("Deine Stats sind jetzt:")
				fmt.Println("HP:", hp)
				fmt.Println("Att:", att)
				fmt.Println("Def:", def)
			}
		}
	end: // End Game
		if choice == 3 {
			break
		}
	}
}
