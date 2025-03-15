package game

import (
	"fmt"
	"math/rand"
	"start/gear"
	"start/player"
	"start/story"
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
// Makes player choose place
// Returns EventArray
func ChoosePlace(placeArray [3]string, world string) [3]string {

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
			eventArray = [3]string{"Robbery", "Bettler", "Muelltonne"}
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
	return eventArray
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

// Gets EventArray
// Chooses random Event from eventArray based on length of eventArray
// Returns Event
func EventGenerator(eventArray [3]string) string {
	var event string
	var eventNumber int = rand.Intn(len(eventArray))

	event = eventArray[eventNumber]

	return event
}

// Gets Event, Gets Everything needed in the game
// Executes Function of equivalent event
// Returns Everything needed in the game
func EventExecution(event string, player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) {
	switch event {
	////////////////////////////// Cyberpunk /////////////////////////////////
	case "Robbery":
		story.Robbery(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Bettler":
		story.Bettler(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Muelltonne":
		story.Muelltonne(player1, inventory, hp, att, def, rec, world, world_barrier)
	}
}
