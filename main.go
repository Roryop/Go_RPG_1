package main

import (
	"fmt"
	"start/game"
	"start/gear"
	"start/menu"
	"start/player"
	"start/story"
	"start/text"
)

/*
	func helloWelt() (string, string) {

		var x, y string
		x, y = "Hello", "World"
		return x, y

	}
*/

func main() {

	///////////////////////////// Prologue ///////////////////////////////

	story.Prologue()

	/////////////////////////// Setup World //////////////////////////////////

	var world string

	/////////////////////// Creating Inventory ///////////////////////////////

	var inventory = gear.NewInventory()

	//////////////////////////// Creating Player ////////////////////////////////

	var player1 = player.InitPlayer() // Creating Player

	var player_level = player1.GetLevel()

	var hp, att, def, rec = player1.CreateStats(inventory) // Creating current Player Stats

	///////////////////////////////Test Environment//////////////////////////

	/*
		fmt.Println(inventory)

		for i := 0; i < 20; i++ {
			inventory = gear.AddDropToInventory(inventory)
		}
	*/

	////////////////////////End Test Environment///////////////////////

	////////////////////////////////// Game /////////////////////////////////////

	var choice int = 0 // Creating Variable so Player can later end game themselves

	for choice != 5 { // Game keeps running until Player end it

		////////////////////////// Refreshing Player /////////////////////////////

		hp, att, def, rec = player1.CreateStats(inventory) // Creating current Player Stats

		//////////////////////////// Menu //////////////////////////////////

		var menuLoop bool = true
		var menuChoice int = menu.Menu() // menuChoice to remember player decision

		for menuLoop {

			switch menuChoice {
			case 1:

				////////////////////// Choosing World /////////////////////////

				world = game.Chooseworld() // NEEDS PROPER WORLD_BARRIER

				goto start

			case 2:

				text.Print("Dein aktueller Schwierigkeitsgrad ist " + fmt.Sprint(player_level) + "!")

				menuChoice = menu.Menu() //In case Game is not ended or started, refresh menuChoice
			case 3:

				player1.SeePlayerStats(inventory, hp, att, def, rec)

				menuChoice = menu.Menu()
			case 4:

				gear.GiveInventoryInformation(inventory)

				menuChoice = menu.Menu()
			case 5:

				text.Print("Not yet implemented.")

				menuChoice = menu.Menu()
			case 6:
				choice = 5

				goto end
			}
		}

	start:

		////////////////////////// Menu End ///////////////////////////////

		for i := 0; i < 3 && hp > 0; i++ { // Entering Fights until Player 1. killed 10 monster; 2. is dead

			///////////////////// Choosing Place and Making Event ////////////////
			var placeArray = game.SetPlaceArray(world)
			var eventArray = game.ChoosePlace(placeArray, world)
			var event = game.EventGenerator(eventArray)
			fmt.Println(event)

			player1, inventory, hp, att, def, rec, world, player_level = game.EventExecution(event, player1, inventory, hp, att, def, rec, world, player_level)

			/*
				////////////////////// Setting Up Enemy ////////////////////////
				var enemyName, enemyStats = enemy.CreateEnemy(world, player_level, 0)

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
					if i >= 9 {
						player_level += 1
					}

					//////////////// Enemy Item Drop /////////////
					inventory = gear.AddDropToInventory(inventory, player_level)

					///////////////// Player Management ////////////////
					player1.Exp_Function(enemyStats)                                           // Giving Exp to Player
					hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) // Player will be healed with levelUp + Updating Stats + Updating current Stats

				}

				//////////////// Clearing Inventory on Player Death ////////////////
				if hp <= 0 {
					inventory = gear.FillEmptyInventory(inventory)
				}
			*/
		}
	end: // End Game
		if choice == 5 {
			break
		}
	}
}
