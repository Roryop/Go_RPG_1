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

	for isInLoop { // Choosing World until valid World has been chosen
		fmt.Println("In welche Welt willst du?")                                        // This example is for User deciding world
		fmt.Println("Tutorial; Cyberpunk; Middleage; Armageddon; Prehistory; Wildwest") ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////Gerade hier
		fmt.Scanln(&world)

		world = strings.ToLower(world) // Enabling any sort of Upper/Lowercase by setting world LowerCase

		if world == "cyberpunk" || world == "middleage" || world == "armageddon" || world == "prehistory" || world == "wildwest" || world == "tutorial" {
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
			eventArray = [3]string{"Baerangriff", "Bonfire", "GesunderPilz"}
			isInLoop = false
		case place == "castle" && world == "middleage":
			eventArray = [3]string{"Hexenjagd", "Burgritter", "Kerker"}
			isInLoop = false
		case place == "village" && world == "middleage":
			eventArray = [3]string{"BetrunkenerDorfbewohner", "Rotzbuben", "Wirtshaus"}
			isInLoop = false

		case place == "ground zero" && world == "armageddon":
			eventArray = [3]string{"Mutantenratte", "DreikoepfigeSchlange", "LaufendeMakrowelle"}
			isInLoop = false
		case place == "settlement" && world == "armageddon":
			eventArray = [3]string{"PluendererGroup", "Stammarzt", "VerletzterBewohner"}
			isInLoop = false
		case place == "ghost town" && world == "armageddon":
			eventArray = [3]string{"PluendererSingle"}
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
func EventExecution(event string, player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, world_barrier int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, string, int) {
	switch event {
	////////////////////////////// Cyberpunk /////////////////////////////////

	//////////////////// Funktionen für Slums ////////////////////////
	case "Robbery":
		player1, inventory, hp, att, def, rec, world_barrier = story.Robbery(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Bettler":
		player1, inventory, hp, att, def, rec, world_barrier = story.Bettler(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Muelltonne":
		inventory, world_barrier = story.Muelltonne(inventory, world_barrier)

	////////////////////////////// Mittelalter /////////////////////////////////

	////////////////////// Funktionen für Wald ////////////////////////////
	case "Baerangriff":
		player1, inventory, hp, att, def, rec, world_barrier = story.Baerangriff(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Bonfire":
		hp = story.Bonfire(player1, inventory, hp)
	case "GesunderPilz":
		hp = story.GesunderPilz(player1, inventory, hp)

	////////////////////// Funktionen für Burg ////////////////////////////
	case "Hexenjagd":
		story.Hexenjagd(player1)
	case "Burgritter":
		player1, inventory, hp, att, def, rec, world_barrier = story.Burgritter(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Kerker":
		player1, inventory, hp, att = story.Kerker(player1, inventory, hp, att)

	////////////////////// Funktionen für Dorf ////////////////////////////
	case "BetrunkenerDorfbewohner":
		player1, inventory, hp, att, def, rec, world_barrier = story.BetrunkenerDorfbewohner(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Rotzbuben":
		story.Rotzbuben(player1)
	case "Wirtshaus":
		hp = story.Wirtshaus(player1, inventory, hp)

	////////////////////////////// Armageddon /////////////////////////////////

	////////////////////// Funktionen für Ground Zero ////////////////////////////
	case "Mutantenratte":
		player1, inventory, hp, att, def, rec, world_barrier = story.Mutantenratte(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "DreikoepfigeSchlange":
		player1, inventory, hp, att, def, rec, world_barrier = story.DreikoepfigeSchlange(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "LaufendeMakrowelle":
		player1, inventory, hp, att, def, rec, world_barrier = story.LaufendeMakrowelle(player1, inventory, hp, att, def, rec, world, world_barrier)

	////////////////////// Funktionen für Settlement ////////////////////////////
	case "PluendererGroup":
		player1, inventory, hp, att, def, rec, world_barrier = story.PluendererGroup(player1, inventory, hp, att, def, rec, world, world_barrier)
	case "Stammarzt":
		hp = story.Stammarzt(player1, inventory, hp)
	case "VerletzterBewohner":
		story.VerletzterBewohner()

	}

	// Final Return Statement
	return player1, inventory, hp, att, def, rec, world, world_barrier
}
