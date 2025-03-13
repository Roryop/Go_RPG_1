package game

import (
	"fmt"
	"math/rand"
	"start/enemy"
	"start/gear"
	"start/player"
	"strings"
)

// Gets world_barrier
// Player Chooses World
// Returns world
func Chooseworld(world_barrier int) string {

	var world string
	var isInLoop bool = true
	for isInLoop {
		fmt.Println("In welche Welt willst du?")                                        // This example is for User deciding world
		fmt.Println("Tutorial; Cyberpunk; Middleage; Armageddon; Prehistory; Wildwest") ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////Gerade hier
		fmt.Scanln(&world)

		world = strings.ToLower(world) // Enabling any sort of Upper/Lowercase by setting world LowerCase

		switch {
		case world == "cyberpunk" && world_barrier >= 1:
			isInLoop = false
		case world == "middleage" && world_barrier >= 2:
			isInLoop = false
		case world == "armageddon" && world_barrier >= 3:
			isInLoop = false
		case world == "prehistory" && world_barrier >= 4:
			isInLoop = false
		case world == "wildwest" && world_barrier >= 5:
			isInLoop = false
		case world == "tutorial":
			isInLoop = false
		}

	}
	fmt.Println("Du hast", strings.ToUpper(world), "ausgewählt")
	return world
}

// Gets placeArray
func ChoosePlace(placeArray [3]string, world string) (string, [3]string) {

	var place string
	var eventArray [3]string
	var isInLoop bool = true
	for isInLoop {
		fmt.Println("Welchen Ort möchtest du besuchen?")
		fmt.Println(placeArray)
		fmt.Scanln(&place)

		place = strings.ToLower(place)

		switch {
		case place == "slums" && world == "cyberpunk":
			eventArray = [3]string{"event1", "event2", "event3"}
			isInLoop = false
		case place == "busimess" && world == "cyberpunk":
			eventArray = [3]string{}
			isInLoop = false
		case place == "outer city" && world == "cyberpunk":
			eventArray = [3]string{}
			isInLoop = false

		case place == "forest" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "castle" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "village" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false

		case place == "ground zero" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "settlement" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "ghost town" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false

		case place == "cave" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "jungle" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "mountain" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false

		case place == "pub" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "sheriff" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "goldmine" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		}
	}
	fmt.Println("Du hast", strings.ToUpper(place), "ausgewählt")
	return place, eventArray
}

// Gets EventArray
// Chooses random Event from eventArray based on length of eventArray
// Returns Event
func EventGenerator(eventArray [3]string) string {
	var event string
	var eventNumber int = rand.Intn(len(eventArray))

	event = eventArray[eventNumber]

	return event
}

// Gets world and world_barrier
// Decides placeArray based on world
// Returns placeArray
func SetPlaceArray(world string) [3]string {
	var placeArray [3]string

	switch world {
	case "cyberpunk":

		placeArray = [3]string{"slums", "busimess", "outer city"}
	case "middleage":

		placeArray = [3]string{"forest", "castle", "village"}
	case "armageddon":

		placeArray = [3]string{"ground zero", "settlement", "ghost town"}
	case "prehistory":

		placeArray = [3]string{"cave", "jungle", "mountain"}
	case "wildwest":

		placeArray = [3]string{"pub", "sherrif", "goldmine"}
	default:
		/*

			case 3001:

				enemy_level = 3001
			default:
		*/
		placeArray = [3]string{"Place1", "Place2", "Place3"}

	}
	return placeArray
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
