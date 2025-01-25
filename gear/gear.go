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
	case "Dagger(Common)":
		g.damage = 1
		g.durability = 10

	case "Spear(Common)":
		g.damage = 2
		g.durability = 10

	/////////////////////// Armor //////////////////////////
	case "Leather Breastplate(Common)":
		g.defense = 2
		g.durability = 9

	/////////////////////// Accessoires ///////////////////
	case "Ring of Strength(Common)":

		g.damage = 1

	case "Ring of Resilience(Common)":

		g.defense = 1

	case "Ring of Recovery(Common)":

		g.recovery = 1

	}

}

///////////////////////////////  Create New Gear  //////////////////////////////////

func NewGear(gearClass string, gearName string) *Gear {

	///////////////////// Create Empty Item ///////////////////////
	var gear *Gear = new(Gear)

	//////////////////// Decide on Empty / Real Item ////////////////////////
	// gearClass of "Weapon", "Armor", and Accessoire is accepted, anything else results in "Empty" "Not Filled"
	switch gearClass {
	case "Weapon":
	case "Armor":
	case "Accessoire":
		gear.gearTyp = gearClass
		gear.name = gearName
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
func ItemDrop() *Gear {
	var rarity string

	var dropNumber int = rand.Intn(10000) + 1
	switch {
	case dropNumber <= 8000 && dropNumber > 3000:
		rarity = "Common"
	case dropNumber <= 9400 && dropNumber > 8000:
		rarity = "Greater"
	case dropNumber <= 9900 && dropNumber > 9400:
		rarity = "Unique"
	case dropNumber <= 9990 && dropNumber > 9900:
		rarity = "Mythical"
	case dropNumber <= 9999 && dropNumber > 9990:
		rarity = "Transcendent"
	case dropNumber == 10000:
		rarity = "Godly"
	default:
		rarity = "Nothing"
	}

	// Creates Item
	var item *Gear

	// Creates Array for deciding itemTyp
	var itemTypArray [3]string = [3]string{"Weapon", "Armor", "Accessoire"}
	// Deciding on ItemTyp through rand.Intn()
	var itemTypNumber int = rand.Intn(3)
	var itemTyp string = itemTypArray[itemTypNumber]

	// Creating ItemArrays of Typ Common
	var weaponArrayCommon [2]string = [2]string{
		"Dagger(Common)",
		"Spear(Common)"}

	var armorArrayCommon [1]string = [1]string{
		"Leather Breastplate(Common)"}

	var accessoireArrayCommon [3]string = [3]string{
		"Ring of Strength(Common)",
		"Ring of Resilience(Common)",
		"Ring of Recovery(Common)"}

	// Creating Variables to store Item Name of corresponding Item
	var itemName string
	var itemArrayNumber int

	// Assigning ItemName according to rarity + rand.Intn()
	switch rarity {
	case "Nothing":

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

	case "Unique":

	case "Mythical":

	case "Transcendent":

	case "Godly":

	}

	// Creating Item based on ItemTyp + ItemName
	item = NewGear(itemTyp, itemName)
	return item
}
