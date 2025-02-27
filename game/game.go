package game

import (
	"fmt"
	"math/rand"
)

// Gets world_barrier
// Player Chooses World
// Returns world
func Chooseworld(world_barrier int) string {

	var world string
	var isInLoop bool = true
	for isInLoop {
		fmt.Println("In welche Welt willst du?")                              // This example is for User deciding world
		fmt.Println("Cyberpunk; Middleage; Armageddon; Prehistory; Wildwest") ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////Gerade hier
		fmt.Scanln(&world)

		switch {
		case world == "Cyberpunk" && world_barrier >= 1: // Namen der Welten bitte Eintragen (hab erstmal 3 geschrieben)
			isInLoop = false
		case world == "Middleage" && world_barrier >= 2: // Den welten fehlen vielleicht Attribute aber KA wie mann die schreibt und ob die überhaupt hier sein sollen
			isInLoop = false
		case world == "Armageddon" && world_barrier >= 3:
			isInLoop = false
		case world == "Prehistory" && world_barrier >= 4:
			isInLoop = false
		case world == "Wildwest" && world_barrier >= 5:
			isInLoop = false
		}
	}
	fmt.Println("Du hast", world, "ausgewählt")
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
		case place == "Place1" && world == "Cyberpunk":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place1" && world == "Cyberpunk":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place1" && world == "Cyberpunk":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place2" && world == "Middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place2" && world == "Middleage":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place2" && world == "Middleage":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "Armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Armageddon":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Armageddon":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "Prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Prehistory":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Prehistory":
			eventArray = [3]string{}
			isInLoop = false

		case place == "Place3" && world == "Wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Wildwest":
			eventArray = [3]string{}
			isInLoop = false
		case place == "Place3" && world == "Wildwest":
			eventArray = [3]string{}
			isInLoop = false
		}
	}
	fmt.Println("Du hast", place, "ausgewählt")
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
	case "Cyberpunk":

		placeArray = [3]string{"Slums", "Busimess", "Outer City"}
	case "Middleage":

		placeArray = [3]string{"Forest", "Castle", "Village"}
	case "Armageddon":

		placeArray = [3]string{"Ground Zero", "Settlement", "Ghost Town"}
	case "Prehistory":

		placeArray = [3]string{"Cave", "Jungle", "Mountain"}
	case "Wildwest":

		placeArray = [3]string{"Pub", "Sherrif", "Goldmine"}
	default:
		/*	case 6:

				//Gegnerlevel zwischen 30 und 35
				enemy_level = rand.Intn(6) + 30
			case 7:

				//Gegnerlevel zwischen 34 und 36
				enemy_level = rand.Intn(3) + 34
			case 8:

				//Gegnerlevel zwischen 36 und 40
				enemy_level = rand.Intn(5) + 36
			case 9:

				//Gegnerlevel zwischen 40 und 45
				enemy_level = rand.Intn(6) + 40
			case 10:

				//Gegnerlevel zwischen 45 und 50
				enemy_level = rand.Intn(6) + 45
			case 3001:

				enemy_level = 3001
			default:
		*/
		placeArray = [3]string{"Place1", "Place2", "Place3"}
		//enemy_level  0
		enemy_level = 0
	}

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
	}
	return placeArray, enemy_level
}
