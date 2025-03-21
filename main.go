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

				text.Print("Du hast keine Karte und kriegst auch keine")
				text.ShortWait()
				fmt.Println("HAHA")
				text.LongWait()

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

		}
	end: // End Game
		if choice == 5 {
			break
		}
	}

	story.Ende(player1)

}
