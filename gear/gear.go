package gear

import "math/rand"

///////////////////////////////  Struct Gear Types  //////////////////////////////////
type Gear struct {
	damage     int
	defense    int
	recovery   int
	durability int
	gearTyp    string
	name       string
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
	case "Accessoire":
		g.durability = -1 //Equipment is destroyed at 0 durability, so Accessoires have -1
	default:
		break
	}

	switch gearName {

	/////////////////////// Weapons ////////////////////////

	////////////// Common ///////////////
	case "Dagger(Common)":
		g.damage = 1
		g.durability = 10

	case "Spear(Common)":
		g.damage = 2
		g.durability = 10

	////////////// Greater ///////////////
	case "Dagger(Greater)":
		g.damage = 9
		g.durability = 10

	case "Spear(Greater)":
		g.damage = 13
		g.durability = 10

	////////////// Unique ///////////////
	case "Rheimzadetz´ Dagger":
		g.damage = 101
		g.durability = 10

	case "Radu´s Greatsword":
		g.damage = 140
		g.defense = 10
		g.durability = 10

	case "Grogerz´ Spear":
		g.damage = 135
		g.durability = 10

	/////////////////////// Armor //////////////////////////

	////////////// Common ///////////////
	case "Leather Breastplate(Common)":
		g.defense = 2
		g.durability = 10

	////////////// Greater ///////////////
	case "Iron Breastplate(Greater)":
		g.defense = 13
		g.durability = 10

	////////////// Unique ///////////////
	case "Rheimzadetz´ Breastplate":
		g.defense = 102
		g.durability = 10

	case "Radu´s Breastplate":
		g.defense = 165
		g.durability = 10

	case "Grogerz´ Breastplate":
		g.defense = 133
		g.durability = 10

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

	var dropNumber int = rand.Intn(10000) + 1
	switch {
	case dropNumber <= 9250-(world_barrier*286) && dropNumber > 5050-(world_barrier*886): //42% , +6%
		rarity = "Common"
	case dropNumber <= 9800-(world_barrier*86) && dropNumber > 9250-(world_barrier*286): //5,5% , +2%
		rarity = "Greater"
	case dropNumber <= 9950-(world_barrier*36) && dropNumber > 9800-(world_barrier*86): //1,5% , +0,5%
		rarity = "Unique"
	case dropNumber <= 9990-(world_barrier*6) && dropNumber > 9950-(world_barrier*36): //0,4% , +0,3%
		rarity = "Mythical"
	case dropNumber <= 9999 && dropNumber > 9990-(world_barrier*6): //0,09% , +0,06%
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
	var weaponArrayCommon [2]string = [2]string{
		"Dagger(Common)",
		"Spear(Common)"}

	var armorArrayCommon [1]string = [1]string{
		"Leather Breastplate(Common)"}

	var accessoireArrayCommon [3]string = [3]string{
		"Ring of Strength(Common)",
		"Ring of Resilience(Common)",
		"Ring of Recovery(Common)"}

	//////////////////////// Item Arrays Rarity Greater //////////////////////
	var weaponArrayGreater [2]string = [2]string{
		"Dagger(Greater)",
		"Spear(Greater)"}

	var armorArrayGreater [1]string = [1]string{
		"Iron Breastplate(Greater)"}

	var accessoireArrayGreater [3]string = [3]string{
		"Ring of Strength(Greater)",
		"Ring of Resilience(Greater)",
		"Ring of Recovery(Greater)"}

	//////////////////////// Item Arrays Rarity Unique //////////////////////
	var weaponArrayUnique [3]string = [3]string{
		"Rheimzadetz´ Dagger",
		"Radu´s Greatsword",
		"Grogerz´ Spear"}

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
