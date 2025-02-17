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
	/////////////////////// Creating Inventory ///////////////////////////////

	var inventory = gear.NewInventory()

	///////////////////////////////Test Environment//////////////////////////

	/*
		fmt.Println(inventory)

		for i := 0; i < 20; i++ {
			inventory = gear.AddDropToInventory(inventory)
		}
	*/

	////////////////////////End Test Environment///////////////////////

	story.Prologue()

	//////////////////////////// Creating Player ////////////////////////////////

	var player1 = player.BeginPlayer() // Creating Player

	var hp, att, def, rec = player1.CreateStats(inventory) // Creating current Player Stats

	////////////////////////////////// Game /////////////////////////////////////

	var choice int = 0 // Creating Variable so Player can later end game themselves
	for choice != 3 {  // Game keeps running until Player end it

		///////////////////// Setting Stats before each Run /////////////////////
		hp, att, def, rec = player1.CreateStats(inventory)
		player1.SeePlayerStats()

		////////////////////// Choosing Game Level /////////////////////////
		var game_level = game.ChooseGameLevel()

		for i := 0; i < 10 && hp > 0; i++ { // Entering Fights until Player 1. killed 10 monster; 2. is dead

			////////////////////// Setting Up Enemy ////////////////////////
			var enemyName, enemyStats = enemy.CreateEnemy(game_level)

			fmt.Println("Du fightest einen", enemyName+"!!!")
			fmt.Println("Er ist Level", enemyStats[0], "!!!")
			fmt.Println("Er hat", enemyStats[1], "HP!")
			fmt.Println("Du hast", hp, "HP!")

			////////////////////////// Fighting ////////////////////////////
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

				//////////////// Enemy Item Drop /////////////
				inventory = gear.AddDropToInventory(inventory)

				///////////////// Player Management ////////////////
				player1.Exp_Function(enemyStats) // Giving Exp to Player
				fmt.Println("Test-Hp1:", hp)
				hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) // Player will be healed with levelUp + Updating Stats + Updating current Stats
				fmt.Println("Test-Hp2:", hp)
				fmt.Println("Deine Stats sind jetzt:")
				fmt.Println("HP:", hp)
				fmt.Println("Att:", att)
				fmt.Println("Def:", def)
				//////////////////////// BUG ////////////////////////////
				// In Method Level_Management, Player is healed through level Up and through stat rec (recovery)
				// That works perfectly, so dont touch method Level_Management !!!!!!!
				// Problem: After adding an Item to the Inventory with AddDropToInventory (Line 108),
				// the new Stats from Inventory are not added to current Player Stats, rendering Inventory useless
				// We need to find a way in order to add the InventoryStats to current Player Stats without
				// interfering with current Stats - making sure past dmg stays applied before healing
				// Second Problem: I have no Idea whether to work with hp, att, def, rec or player1.Stats[]
				// Your Task: Figure out whether we have to work with the former or the latter, and if you can,
				// FIX THE PROBLEM

				// PS: Dont forget to test the code after every change BEFORE committing
				// PS After the PS: I have no Idea how this Bug even managed to slip in, must´ve been carelessness
			}

			//////////////// Clearing Inventory on Player Death ////////////////
			if hp <= 0 {
				inventory = gear.FillEmptyInventory(inventory)
			}
		}
	end: // End Game
		if choice == 3 {
			break
		}
	}
}
