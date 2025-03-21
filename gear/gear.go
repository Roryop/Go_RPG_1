package gear

import "math/rand"

// /////////////////////////////  Struct Gear Types  //////////////////////////////////
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

	case "Großschwert":
		g.damage = 16

	////////////// Unique ///////////////
	case "Bastard Schwert":
		g.damage = 61

	case "Ankylosaurus Schwanz":
		g.damage = 37
		g.defense = 10

	case "Mantis Klingen":
		g.damage = 48
		g.defense = 15

	//Nicht in WeaponArray (DropPool)
	case "Smith & Wesson 500 Bone Crusher":
		g.damage = 150

	////////////// Mythical ///////////////
	case "Sternenstoff":
		g.damage = 161

	case "Himmlisches Eisen":
		g.damage = 164

	////////////// Transcendent ///////////////
	case "Sternenwaffe":
		g.damage = 400
		g.defense = 100
		g.recovery = 50

	/////////////////////// Armor //////////////////////////

	////////////// Common ///////////////
	case "Enge Leder Hose":
		g.defense = 2

	case "Weihnachtssocken":
		g.defense = 3

	case "Ungewaschener Unterhemd":
		g.defense = 5

	////////////// Greater ///////////////
	case "Eisen Rüstung":
		g.defense = 13

	case "Kettenhemd":
		g.defense = 11

	case "Gold Rüstung":
		g.defense = 15

	////////////// Unique ///////////////
	case "Kevlar Rüstung":
		g.defense = 82

	case "Pyrex Glass Rüstung":
		g.defense = 93

	case "Obsidian Rüstung":
		g.defense = 79

	////////////// Mythical ///////////////
	case "Sternengewand":
		g.defense = 170

	case "Himmlischer Eimer":
		g.defense = 169

	////////////// Transcendent ///////////////
	case "Stern Rüstung":
		g.damage = 50
		g.defense = 500
		g.recovery = 100

	/////////////////////// Accessoires ///////////////////

	////////////// Common ///////////////
	case "Spinat":

		g.damage = 1

	case "Yakult":

		g.defense = 2

	case "Hello Kitty Verband":

		g.recovery = 2

	////////////// Greater ///////////////
	case "Mehr Spinat":

		g.damage = 11

	case "Actimel":

		g.defense = 14

	case "Lecker Bierchen":

		g.recovery = 13

	////////////// Unique ///////////////
	case "Popeyes Eigener Spinat":
		g.damage = 50

	case "Flintstone Vitamine":
		g.defense = 50

	case "Starker Glaube":
		g.recovery = 50

	////////////// Mythical ///////////////
	case "Stern Fragment":
		g.damage = 150

	case "Stern Kern":
		g.defense = 150

	case "Sternstaub":
		g.recovery = 150

	////////////// Transcendent ///////////////
	case "Stern im Gurkenglas":
		g.damage = 300
		g.defense = 300
		g.recovery = 300

	///////////////// Godly /////////////////
	case "Neutronen-Stern":
		g.damage = 1000
		g.defense = 1000
		g.recovery = 1000
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

func (w *Gear) GetGearTyp() string {
	return w.gearTyp
}

// Gets Nothing
// Creates Item based on rand.Intn() + ItemArrays
// Returns Item
func ItemDrop(player_level int) *Gear {
	var rarity string
	var dropLimit = (player_level - 1) / 2 // DropChance Increase for every 2 player levels, starting from 3

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

	var armorArrayCommon [3]string = [3]string{
		"Enge Leder Hose",
		"Weihnachtssocken",
		"Ungewaschener Unterhemd"}

	var accessoireArrayCommon [3]string = [3]string{
		"Spinat",
		"Yakult",
		"Hello Kitty Verband"}

	//////////////////////// Item Arrays Rarity Greater //////////////////////
	var weaponArrayGreater [3]string = [3]string{
		"Flamberge",
		"Großschwert",
		"Rapier"}

	var armorArrayGreater [3]string = [3]string{
		"Eisen Rüstung",
		"Kettenhemd",
		"Gold Rüstung"}

	var accessoireArrayGreater [3]string = [3]string{
		"Mehr Spinat",
		"Actimel",
		"Lecker Bierchen"}

	//////////////////////// Item Arrays Rarity Unique //////////////////////
	var weaponArrayUnique [1]string = [1]string{
		"Bastard Schwert"}

	var armorArrayUnique [3]string = [3]string{
		"Kevlar Rüstung",
		"Pyrex Glass Rüstung",
		"Obsidian Rüstung"}

	var accessoireArrayUnique [3]string = [3]string{
		"Popeyes Eigener Spinat",
		"Flintstone Vitamine",
		"Starker Glaube"}

	//////////////////////// Item Arrays Rarity Mythical //////////////////////
	var weaponArrayMythical [2]string = [2]string{
		"Sternenstoff",
		"Himmlisches Eisen"}
	var armorArrayMythical [2]string = [2]string{
		"Sternengewand",
		"Himmlischer Eimer"}
	var accessoireArrayMythical [3]string = [3]string{
		"Stern Fragment",
		"Stern Kern",
		"Sternstaub"}

	//////////////////////// Item Arrays Rarity Transcendent //////////////////////
	var weaponArrayTranscendent [1]string = [1]string{
		"Sternenwaffe"}
	var armorArrayTranscendent [1]string = [1]string{
		"Stern Rüstung"}
	var accessoireArrayTranscendent [1]string = [1]string{
		"Stern im Gurkenglas"}

	// Creating Variables to store Item Name of corresponding Item
	var itemName string
	var itemArrayNumber int

	// Assigning ItemName according to rarity + rand.Intn()
	switch rarity {
	case "Nothing":
		itemTyp, itemName = "Empty", "Not Filled"
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
		switch itemTyp {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayMythical))
			itemName = weaponArrayMythical[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayMythical))
			itemName = armorArrayMythical[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayMythical))
			itemName = accessoireArrayMythical[itemArrayNumber]
		}
	case "Transcendent":
		switch itemTyp {
		case "Weapon":
			itemArrayNumber = rand.Intn(len(weaponArrayTranscendent))
			itemName = weaponArrayTranscendent[itemArrayNumber]
		case "Armor":
			itemArrayNumber = rand.Intn(len(armorArrayTranscendent))
			itemName = armorArrayTranscendent[itemArrayNumber]
		case "Accessoire":
			itemArrayNumber = rand.Intn(len(accessoireArrayTranscendent))
			itemName = accessoireArrayTranscendent[itemArrayNumber]
		}
	case "Godly":
		itemTyp = "Accessoire"
		itemName = "Neutronen-Stern"
	}

	// Creating Item based on ItemTyp + ItemName
	item = NewGear(itemTyp, itemName)
	return item
}
