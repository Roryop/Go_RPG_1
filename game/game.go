package game

import (
	"fmt"
	"math/rand"
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
func ChoosePlace(placeArray [3]string, world string) [3]string {

	var place string
	var eventArray [3]string
	var isInLoop bool = true
	for isInLoop {
		fmt.Println("Welchen Ort möchtest du besuchen?")
		fmt.Println(placeArray)
		fmt.Scanln(&place)

		switch {
		case place == "Place1" && world == "cyberpunk":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place1" && world == "cyberpunk":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place1" && world == "cyberpunk":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place2" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place2" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place2" && world == "middleage":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "armageddon":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "prehistory":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "wildwest":
			eventArray = [3]string{}
			isInLoop = false
		}
	}
	fmt.Println("Du hast", strings.ToUpper(place), "ausgewählt")
	return eventArray
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
// Decides Events and Enemy_level based on world + rand.Intn()
// Returns eventArray and Enemy_level
func SetWorldEnemy(world string, world_barrier int) ([3]string, int) {
	var placeArray [3]string
	var enemy_level = 0

	switch world {
	case "cyberpunk":

		placeArray = [3]string{"Slums", "Busimess", "Outer City"}
	case "middleage":

		placeArray = [3]string{"Forest", "Castle", "Village"}
	case "armageddon":

		placeArray = [3]string{"Ground Zero", "Settlement", "Ghost Town"}
	case "prehistory":

		placeArray = [3]string{"Cave", "Jungle", "Mountain"}
	case "wildwest":

		placeArray = [3]string{"Pub", "Sherrif", "Goldmine"}
	default:
		/*

			case 3001:

				enemy_level = 3001
			default:
		*/
		placeArray = [3]string{"Place1", "Place2", "Place3"}

	}

	// Only changes enemy_level based on world_barrier if correct world was chosen, tutorial will stay level 0
	if world == "cyberpunk" || world == "middleage" || world == "armageddon" || world == "prehistory" || world == "wildwest" {
		switch world_barrier {
		case 1:

			//enemy_level between 1 and 3
			enemy_level = rand.Intn(3) + 1
		case 2:

			//enemy_level between 3 and 5
			enemy_level = rand.Intn(3) + 3
		case 3:

			//enemy_level between 6 and 10
			enemy_level = rand.Intn(5) + 6
		case 4:

			//Gegnerlevel between 11 und 20
			enemy_level = rand.Intn(10) + 11
		case 5:

			//Gegnerlevel between 15 und 25
			enemy_level = rand.Intn(11) + 15
		case 6:

			//Gegnerlevel between 30 und 35
			enemy_level = rand.Intn(6) + 30
		case 7:

			//Gegnerlevel between 34 und 36
			enemy_level = rand.Intn(3) + 34
		case 8:

			//Gegnerlevel between 36 und 40
			enemy_level = rand.Intn(5) + 36
		case 9:

			//Gegnerlevel between 40 und 45
			enemy_level = rand.Intn(6) + 40
		case 10:

			//Gegnerlevel between 45 und 50
			enemy_level = rand.Intn(6) + 45
		}
	}
	return placeArray, enemy_level
}
