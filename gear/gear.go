package gear

import "math/rand"

///////////////////////////////  Struct Gear Types  //////////////////////////////////
type Gear struct {
	damage   int
	defense  int
	recovery int
	gearTyp  string
	name     string
}

///////////////////////////////  Set Attributes  //////////////////////////////////

// Sets Attributes of Items based on name >> "Creating" said Items in "Database"
func (g *Gear) SetAttributesGear() {
	var gearName = g.name

	switch g.gearTyp {
	case "Weapon":
		g.defense = 0
	case "Armor":
		g.damage = 0
	default:
		break
	}

	switch gearName {

	/////////////////////// Weapons ////////////////////////

	////////////// Common ///////////////
	case "Dolch":
		g.damage = 4

	case "Rostige Nagel":
		g.damage = 3

	case "Spitzer Stein":
		g.damage = 1

	case "Sock Mace":
		g.damage = 2

	////////////// Greater ///////////////
	case "Rapier":
		g.damage = 11

	case "Flamberge":
		g.damage = 14

	case "Claymore":
		g.damage = 16

	////////////// Unique ///////////////
	case "Bastard Sword":
		g.damage = 135

	//Nicht in WeaponArray (DropPool)
	case "Ankylosaurus Schwanz":
		g.damage = 140
		g.defense = 10

	case "Mantis Klingen":
		g.damage = 135
		g.defense = 15

	case "smith & wesson 500 bone crusher":
		g.damage = 150

	/////////////////////// Armor //////////////////////////

	////////////// Common ///////////////
	case "Leather Breastplate(Common)":
		g.defense = 2

	////////////// Greater ///////////////
	case "Iron Breastplate(Greater)":
		g.defense = 13

	////////////// Unique ///////////////
	case "Rheimzadetz´ Breastplate":
		g.defense = 102

	case "Radu´s Breastplate":
		g.defense = 165

	case "Grogerz´ Breastplate":
		g.defense = 133

	/////////////////////// Accessoires ///////////////////

	////////////// Common ///////////////
	case "Ring of Strength(Common)":

		g.damage = 1

	case "Ring of Resilience(Common)":

		g.defense = 1

	case "Ring of Recovery(Common)":

		g.recovery = 1

	////////////// Greater ///////////////
	case "Ring of Strength(Greater)":

		g.damage = 10

	case "Ring of Resilience(Greater)":

		g.defense = 10

	case "Ring of Recovery(Greater)":

		g.recovery = 10

	////////////// Unique ///////////////
	case "Rheimzadetz´ Ring of Strength":
		g.damage = 104

	case "Radu´s Ring of Recovery":
		g.recovery = 137

	case "Grogerz´ Ring of Resilience":
		g.defense = 170
	}

}

///////////////////////////////  Create New Gear  //////////////////////////////////

func NewGear(itemTyp string, itemName string) *Gear {

	///////////////////// Create Empty Item ///////////////////////
	var gear *Gear = new(Gear)

	//////////////////// Decide on Empty / Real Item ////////////////////////
	// gearClass of "Weapon", "Armor", and Accessoire is accepted, anything else results in "Empty" "Not Filled"
	switch itemTyp {
	case "Weapon":
	case "Armor":
	case "Accessoire":
		gear.gearTyp = itemTyp
		gear.name = itemName
	default:
		gear.gearTyp = "Empty"
		gear.name = "Not Filled"
	}

	//////////////////////// Set Gear Attributes ////////////////////

	gear.SetAttributesGear()

	//////////////////////// Return Item /////////////////////

	return gear
}

// Gets Nothing
// Creates Item based on rand.Intn() + ItemArrays
// Returns Item
func ItemDrop(world_barrier int) *Gear {
	var rarity string
	var dropLimit = world_barrier - 1

	var dropNumber int = rand.Intn(10000) + 1
	switch {
	case dropNumber <= 9250-(dropLimit*286) && dropNumber > 5050-(dropLimit*886): //42% , +6%
		rarity = "Common"
	case dropNumber <= 9800-(dropLimit*86) && dropNumber > 9250-(dropLimit*286): //5,5% , +2%
		rarity = "Greater"
	case dropNumber <= 9950-(dropLimit*36) && dropNumber > 9800-(dropLimit*86): //1,5% , +0,5%
		rarity = "Unique"
	case dropNumber <= 9990-(dropLimit*6) && dropNumber > 9950-(dropLimit*36): //0,4% , +0,3%
		rarity = "Mythical"
	case dropNumber <= 9999 && dropNumber > 9990-(dropLimit*6): //0,09% , +0,06%
		rarity = "Transcendent"
	case dropNumber == 10000: //0,01%
		rarity = "Godly"
	default:
		rarity = "Nothing" //50,5%
	}

	// Creates Item
	var item *Gear

	// Creates Array for deciding itemTyp
	var itemTypArray [3]string = [3]string{"Weapon", "Armor", "Accessoire"}
	// Deciding on ItemTyp through rand.Intn()
	var itemTypNumber int = rand.Intn(3)
	var itemTyp string = itemTypArray[itemTypNumber]

	//////////////////////// Item Arrays Rarity Common //////////////////////
	var weaponArrayCommon [4]string = [4]string{
		"Dolch",
		"Rostige Nagel",
		"Spitzer Stein",
		"Sock Mace"}

	var armorArrayCommon [1]string = [1]string{
		"Leather Breastplate(Common)"}

	var accessoireArrayCommon [3]string = [3]string{
		"Ring of Strength(Common)",
		"Ring of Resilience(Common)",
		"Ring of Recovery(Common)"}

	//////////////////////// Item Arrays Rarity Greater //////////////////////
	var weaponArrayGreater [3]string = [3]string{
		"Flamberge",
		"Claymore",
		"Rapier"}

	var armorArrayGreater [1]string = [1]string{
		"Iron Breastplate(Greater)"}

	var accessoireArrayGreater [3]string = [3]string{
		"Ring of Strength(Greater)",
		"Ring of Resilience(Greater)",
		"Ring of Recovery(Greater)"}

	//////////////////////// Item Arrays Rarity Unique //////////////////////
	var weaponArrayUnique [3]string = [3]string{
		"Bastard Sword"}

	var armorArrayUnique [3]string = [3]string{
		"Rheimzadetz´ Breastplate",
		"Radu´s Breastplate",
		"Grogerz´ Breastplate"}

	var accessoireArrayUnique [3]string = [3]string{
		"Rheimzadetz´ Ring of Strength",
		"Radu´s Ring of Recovery",
		"Grogerz´ Ring of Resilience"}

	// Creating Variables to store Item Name of corresponding Item
	var itemName string
	var itemArrayNumber int

	// Assigning ItemName according to rarity + rand.Intn()
	switch rarity {
	case "Nothing":
		itemTyp, itemName = "", ""
	case "Common":
		switch itemTyp {
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

	case "Greater":
		switch itemTyp {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayGreater))
			itemName = weaponArrayGreater[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayGreater))
			itemName = armorArrayGreater[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayGreater))
			itemName = accessoireArrayGreater[itemArrayNumber]
		}
	case "Unique":
		switch itemTyp {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayUnique))
			itemName = weaponArrayUnique[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayUnique))
			itemName = armorArrayUnique[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayUnique))
			itemName = accessoireArrayUnique[itemArrayNumber]
		}
	case "Mythical":

	case "Transcendent":

	case "Godly":

	}

	// Creating Item based on ItemTyp + ItemName
	item = NewGear(itemTyp, itemName)
	return item
}
