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

	var dropNumber int = rand.Intn(10000) + 1
	switch {
	case dropNumber <= 8000 && dropNumber > 3000:
		rarityGrade = 1
	case dropNumber <= 9400 && dropNumber > 8000:
		rarityGrade = 2
	case dropNumber <= 9900 && dropNumber > 9400:
		rarityGrade = 3
	case dropNumber <= 9990 && dropNumber > 9900:
		rarityGrade = 4
	case dropNumber <= 9999 && dropNumber > 9990:
		rarityGrade = 5
	case dropNumber == 10000:
		rarityGrade = 6
	default:
		rarityGrade = 0
	}

	var item *gear.Gear

	var itemTypeArray [3]string = [3]string{"Weapon", "Armor", "Accessoire"}
	var itemTypeNumber int = rand.Intn(3)
	var itemType string = itemTypeArray[itemTypeNumber]

	var weaponArrayCommon [2]string = [2]string{
		"Dagger(Common)",
		"Spear(Common)"}
	var armorArrayCommon [1]string = [1]string{
		"Leather Breastplate(Common)"}
	var accessoireArrayCommon [3]string = [3]string{
		"Ring of Strength(Common)",
		"Ring of Resilience(Common)",
		"Ring of Recovery(Common)"}

	var itemName string
	var itemArrayNumber int

	switch rarityGrade {
	case 0:
		item = gear.NewGear("", "")
	case 1:

		switch itemType {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayCommon))
			itemName = weaponArrayCommon[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayCommon))
			itemName = armorArrayCommon[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayCommon))
			itemName = accessoireArrayCommon[itemArrayNumber]
		}

		item = gear.NewGear(itemType, itemName)
	case 2:
		item = gear.NewGear(itemType, "")
	case 3:
		item = gear.NewGear(itemType, "")
	case 4:
		item = gear.NewGear(itemType, "")
	case 5:
		item = gear.NewGear(itemType, "")
	case 6:
		item = gear.NewGear(itemType, "")
	}

	return item
}
