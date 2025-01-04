package game

import (
	"fmt"
	"math/rand"
	"start/gear"
)

func ChooseGameLevel() int {
	var game_level = 0
	fmt.Println("In welches Level möchtest du?")
	fmt.Scanln(&game_level)
	fmt.Println("Du hast Level", game_level, "ausgewählt.")
	return game_level
}

func ItemDrop() *gear.Gear {
	var rarityGrade int

	var dropNumber int = rand.Intn(101)
	switch {
	case dropNumber <= 30:
		rarityGrade = 1
	case dropNumber <= 50 && dropNumber > 30:
		rarityGrade = 2
	case dropNumber <= 60 && dropNumber > 50:
		rarityGrade = 3
	case dropNumber <= 69 && dropNumber > 60:
		rarityGrade = 4
	case dropNumber <= 77 && dropNumber > 69:
		rarityGrade = 5
	case dropNumber <= 84 && dropNumber > 77:
		rarityGrade = 6
	case dropNumber <= 90 && dropNumber > 84:
		rarityGrade = 7
	case dropNumber <= 95 && dropNumber > 90:
		rarityGrade = 8
	case dropNumber <= 97 && dropNumber > 95:
		rarityGrade = 9
	case dropNumber == 100:
		rarityGrade = 10
	default:
		rarityGrade = 0
	}

	var item *gear.Gear

	switch rarityGrade {
	case 0:
		item = gear.NewGear("", "")
	case 1:
		item = gear.NewGear("Weapon", "Lesser Dagger")
	case 2:
		item = gear.NewGear("", "")
	case 3:
		item = gear.NewGear("", "")
	case 4:
		item = gear.NewGear("", "")
	case 5:
		item = gear.NewGear("", "")
	case 6:
		item = gear.NewGear("", "")
	case 7:
		item = gear.NewGear("", "")
	case 8:
		item = gear.NewGear("", "")
	case 9:
		item = gear.NewGear("", "")
	case 10:
		item = gear.NewGear("", "")
	}

	return item
}
