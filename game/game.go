package game

import (
	"fmt"
	"math/rand"
)

// Gets world_barrier
// Player Chooses World
// Returns world
func Chooseworld(world_barrier int) int {

	var world int
	var isInLoop bool
	for isInLoop {
		fmt.Println("In welche Welt willst du?") // This example is for User deciding world
		fmt.Scanln(&world)

		switch {
		case world == 1 && world_barrier >= 1: // Namen der Welten bitte Eintragen (hab erstmal 3 geschrieben)
			isInLoop = false
		case world == 2 && world_barrier >= 2: // Den welten fehlen vielleicht Attribute aber KA wie mann die schreibt und ob die überhaupt hier sein sollen
			isInLoop = false
		case world == 3 && world_barrier >= 3:
			isInLoop = false
		}
	}
	fmt.Println("Du hast", world, "ausgewählt")
	return world
}

// Gets world and world_barrier
// Decides Events and Enemy_level based on world + rand.Intn()
// Returns eventArray and Enemy_level
func SetWorldEnemy(world int, world_barrier int) ([3]string, int) {
	var placeArray [3]string
	var enemy_level = 0

	switch world {
	case 1:

		placeArray = [3]string{"Place1", "Place2", "Place3"} // Please Input Place Names
	case 2:

		placeArray = [3]string{"Place1", "Place2", "Place3"}
	case 3:

		placeArray = [3]string{"Place1", "Place2", "Place3"}
	case 4:

		placeArray = [3]string{"Place1", "Place2", "Place3"}
	case 5:

		placeArray = [3]string{"Place1", "Place2", "Place3"}
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

// Gets placeArray
// creates
// Chooses random Event
// returns chosen Event
func ChoosePlace(placeArray [3]string) [3]string {
	var eventArray [3]string
	var isInLoop bool = true
	var place string

	for isInLoop {
		fmt.Println("Welchen Ort möchtest du besuchen?") // This example is for User deciding world
		fmt.Println(placeArray)
		fmt.Scanln(&place)

		switch place {
		case "Place1": // Namen der Welten bitte Eintragen (hab erstmal 3 geschrieben)()
			eventArray = [3]string{"Event1", "Event2", "Event3"}
			isInLoop = false
		case "Place2": // Den welten fehlen vielleicht Attribute aber KA wie mann die schreibt und ob die überhaupt hier sein sollen
			eventArray = [3]string{"Event1", "Event2", "Event3"}
			isInLoop = false
		case "Place3":
			eventArray = [3]string{"Event1", "Event2", "Event3"}
			isInLoop = false
		}
	}
	return eventArray
}

func ChooseEvent(eventArray [3]string) [3]string {
	var event string
	var eventNumber int
 
	eventNumber = rand.Intn(3) + 1

	switch eventNumber {
	case 1:
		event = 


	}

	return eventArray
}