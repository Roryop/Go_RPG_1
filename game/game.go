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
func Chooseworld() string {

	var world string
	var isInLoop bool = true

	for isInLoop { // Choosing World until valid World has been chosen
		fmt.Println("In welche Welt willst du?")                                             // This example is for User deciding world
		fmt.Println("Tutorial; Cyberpunk; Mittelalter; Armageddon; Prähistorie; Wildwesten") ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////Gerade hier
		fmt.Scanln(&world)

		world = strings.ToLower(world) // Enabling any sort of Upper/Lowercase by setting world LowerCase

		if world == "cyberpunk" || world == "mittelalter" || world == "armageddon" || world == "prähistorie" || world == "wildwesten" || world == "tutorial" {
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
		case world == "tutorial":
			eventArray = [3]string{"Ork", "DemoKarma", "DemoAuswahl"}
			isInLoop = false

		case place == "elendsviertel" && world == "cyberpunk":
			eventArray = [3]string{"Robbery", "Bettler", "Muelltonne"}
			isInLoop = false
		case place == "busimess" && world == "cyberpunk":
			eventArray = [3]string{"Wallet", "Security", "Businessman"}
			isInLoop = false
		case place == "außenstadt" && world == "cyberpunk":
			eventArray = [3]string{"Grenzkontrolle", "Motel", "Verlaufen"}
			isInLoop = false

		case place == "wald" && world == "mittelalter":
			eventArray = [3]string{"Baerangriff", "Bonfire", "GesunderPilz"}
			isInLoop = false
		case place == "burg" && world == "mittelalter":
			eventArray = [3]string{"Hexenjagd", "Burgritter", "Kerker"}
			isInLoop = false
		case place == "dorf" && world == "mittelalter":
			eventArray = [3]string{"BetrunkenerDorfbewohner", "Rotzbuben", "Wirtshaus"}
			isInLoop = false

		case place == "zero" && world == "armageddon":
			eventArray = [3]string{"Mutantenratte", "DreikoepfigeSchlange", "LaufendeMakrowelle"}
			isInLoop = false
		case place == "siedlung" && world == "armageddon":
			eventArray = [3]string{"PluendererGroup", "Stammarzt", "VerletzterBewohner"}
			isInLoop = false
		case place == "geisterstadt" && world == "armageddon":
			eventArray = [3]string{"PluendererSingle"}
			isInLoop = false

		case place == "hoehle" && world == "prähistorie":
			eventArray = [3]string{"HoehlenmenschenWaffe", "HoehlenmenschenHunger", "HoehlenmenschenWerkzeug"}
			isInLoop = false
		case place == "jungel" && world == "prähistorie":
			eventArray = [3]string{"Triceratops", "Teich", "EatWoman"}
			isInLoop = false
		case place == "berg" && world == "prähistorie":
			eventArray = [3]string{"MountainPath", "Pterodactylus", "Bergsteigerziegen"}
			isInLoop = false

		case place == "bar" && world == "wildwesten":
			eventArray = [3]string{"Whiskey", "Schlaegerei", "Barueberfall"}
			isInLoop = false
		case place == "sheriff" && world == "wildwesten":
			eventArray = [3]string{"Jailbreak", "SheriffQuest", "SheriffBerauben"}
			isInLoop = false
		case place == "goldmine" && world == "wildwesten":
			eventArray = [3]string{"Koyothoehle", "AufDieFresseKriegen", "Banditenlager"}
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

		placeArray = [3]string{"elendsviertel", "busimess", "außenstadt"}
	case "mittelalter":

		placeArray = [3]string{"wald", "burg", "dorf"}
	case "armageddon":

		placeArray = [3]string{"zero", "siedlung", "geisterstadt"}
	case "prähistorie":

		placeArray = [3]string{"hoele", "jungel", "berg"}
	case "wildwesten":

		placeArray = [3]string{"bar", "sheriff", "goldmine"}
	default:
		/*

			case 3001:

				enemy_level = 3001
			default:
		*/
		placeArray = [3]string{"Raum", "Raum", "Raum"}

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
func EventExecution(event string, player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, string, int) {
	switch event {
	////////////////////////////// Tutorial //////////////////////////////////
	case "Ork":
		player1, inventory, hp, att, def, rec, player_level = story.Ork(player1, inventory, hp, att, def, rec, world, player_level)
	case "DemoKarma":
		story.DemoKarma(player1)
	case "DemoAuswahl":
		inventory = story.DemoAuswahl(player1, inventory, player_level)

	////////////////////////////// Cyberpunk /////////////////////////////////

	//////////////////// Funktionen für Slums ////////////////////////
	case "Robbery":
		player1, inventory, hp, att, def, rec, player_level = story.Robbery(player1, inventory, hp, att, def, rec, world, player_level)
	case "Bettler":
		player1, inventory, hp, att, def, rec, player_level = story.Bettler(player1, inventory, hp, att, def, rec, world, player_level)
	case "Muelltonne":
		inventory = story.Muelltonne(inventory, player_level)

		//////////////////// Funktionen für Außenstadt ////////////////////////
	case "Grenzkontrolle":
		player1, inventory, hp, att, def, rec, player_level = story.Grenzkontrolle(player1, inventory, hp, att, def, rec, world, player_level)
	case "Motel":
		hp = story.Motel(player1, inventory, hp)
	case "Verlaufen":
		hp = story.Verlaufen(player1, inventory, hp)

	//////////////////// Funktionen für Businessviertel ////////////////////////
	case "Wallet":
		inventory, player_level = story.Wallet(player1, inventory, player_level)
	case "Security":
		inventory = story.Security(inventory)
	case "Businessman":
		player1, inventory, hp, att, def, rec, player_level = story.Businessman(player1, inventory, hp, att, def, rec, world, player_level)

	////////////////////////////// Mittelalter /////////////////////////////////

	////////////////////// Funktionen für Wald ////////////////////////////
	case "Baerangriff":
		player1, inventory, hp, att, def, rec, player_level = story.Baerangriff(player1, inventory, hp, att, def, rec, world, player_level)
	case "Bonfire":
		hp = story.Bonfire(player1, inventory, hp)
	case "GesunderPilz":
		hp = story.GesunderPilz(player1, inventory, hp)

	////////////////////// Funktionen für Burg ////////////////////////////
	case "Hexenjagd":
		story.Hexenjagd(player1)
	case "Burgritter":
		player1, inventory, hp, att, def, rec, player_level = story.Burgritter(player1, inventory, hp, att, def, rec, world, player_level)
	case "Kerker":
		player1, inventory, hp, att = story.Kerker(player1, inventory, hp, att, def, rec)

	////////////////////// Funktionen für Dorf ////////////////////////////
	case "BetrunkenerDorfbewohner":
		player1, inventory, hp, att, def, rec, player_level = story.BetrunkenerDorfbewohner(player1, inventory, hp, att, def, rec, world, player_level)
	case "Rotzbuben":
		story.Rotzbuben(player1)
	case "Wirtshaus":
		hp = story.Wirtshaus(player1, inventory, hp)

	////////////////////////////// Armageddon /////////////////////////////////

	////////////////////// Funktionen für Ground Zero ////////////////////////////
	case "Mutantenratte":
		player1, inventory, hp, att, def, rec, player_level = story.Mutantenratte(player1, inventory, hp, att, def, rec, world, player_level)
	case "DreikoepfigeSchlange":
		player1, inventory, hp, att, def, rec, player_level = story.DreikoepfigeSchlange(player1, inventory, hp, att, def, rec, world, player_level)
	case "LaufendeMakrowelle":
		player1, inventory, hp, att, def, rec, player_level = story.LaufendeMakrowelle(player1, inventory, hp, att, def, rec, world, player_level)

	////////////////////// Funktionen für Settlement ////////////////////////////
	case "PluendererGroup":
		player1, inventory, hp, att, def, rec, player_level = story.PluendererGroup(player1, inventory, hp, att, def, rec, world, player_level)
	case "Stammarzt":
		hp = story.Stammarzt(player1, inventory, hp, att, def, rec)
	case "VerletzterBewohner":
		story.VerletzterBewohner(player1)

	////////////////////// Funktionen für Geisterstadt ////////////////////////////
	case "PluendererSingle":
		player1, inventory, hp, att, def, rec, player_level = story.PluendererSingle(player1, inventory, hp, att, def, rec, world, player_level)
	case "GeschäftPlündern":
		inventory = story.GeschäftPlündern(player1, inventory, player_level)
	case "Hilfeschreie":
		player1, inventory, hp, att, def, rec, player_level = story.Hilfeschreie(player1, inventory, hp, att, def, rec, world, player_level)

	////////////////////////////// Wild West /////////////////////////////////

	/////////////////////// Funktionen für Bar ///////////////////////////////
	case "Whiskey":
		story.Whiskey(player1)
	case "Schlaegerei":
		player1, inventory, hp, att, def, rec, player_level = story.Schlaegerei(player1, inventory, hp, att, def, rec, world, player_level)
	case "Barueberfall":
		player1, inventory, hp, att, def, rec, player_level = story.Barueberfall(player1, inventory, hp, att, def, rec, world, player_level)

	////////////////// Funktionen für Sheriff Büro ///////////////////////
	case "Jailbreak":
		player1, inventory, hp, att, def, rec, player_level = story.Jailbreak(player1, inventory, hp, att, def, rec, world, player_level)
	case "SheriffQuest":
		inventory = story.SheriffQuest(player1, inventory, player_level)
	case "SheriffBerauben":
		inventory = story.SheriffBerauben(player1, inventory, player_level)

	////////////////// Funktionen für Goldmine ///////////////////////////////
	case "Koyothoehle":
		player1, inventory, hp, att, def, rec, player_level = story.Koyothoehle(player1, inventory, hp, att, def, rec, world, player_level)
	case "AufDieFresseKriegen":
		hp = story.AufDieFresseKriegen(player1, inventory, hp)
	case "Banditenlager":
		player1, inventory, hp, att, def, rec, player_level = story.Banditenlager(player1, inventory, hp, att, def, rec, world, player_level)

	////////////////////////////// Prehistory /////////////////////////////////

	////////////////////// Funktionen für Höhle ////////////////////////////
	case "HoelenmenschenWaffe":
		inventory = story.HoelenmenschenWaffe(player1, inventory)
	case "HoelenmenschenHunger":
		hp = story.HoelenmenschenHunger(player1, hp)
	case "HoelenmenschenWerkzeug":
		story.HoelenmenschenWerkzeug(player1)

	////////////////////// Funktionen für Dschungel ////////////////////////////
	case "Triceratops":
		player1, inventory, hp, att, def, rec, player_level = story.Triceratops(player1, inventory, hp, att, def, rec, world, player_level)
	case "Teich":
		player1, inventory, hp, att, def, rec, player_level = story.Teich(player1, inventory, hp, att, def, rec, world, player_level)
	case "EatWoman":
		hp, att, def, rec = story.EatWoman(player1, inventory, hp, att, def, rec)

	////////////////////// Funktionen für Berg ////////////////////////////
	case "MountainPath":
		player1, inventory, hp, att, def, rec, player_level = story.MountainPath(player1, inventory, hp, att, def, rec, world, player_level)
	case "Pterodactylus":
		player1, inventory, hp, att, def, rec, player_level = story.Pterodactylus(player1, inventory, hp, att, def, rec, world, player_level)
	case "Bergsteigerziegen":
		player1, inventory, hp, att, def, rec, player_level = story.Bergsteigerziegen(player1, inventory, hp, att, def, rec, world, player_level)
	}

	// Final Return Statement
	return player1, inventory, hp, att, def, rec, world, player_level
}
