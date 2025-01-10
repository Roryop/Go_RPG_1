package gear

///////////////////////////////  Struct Gear Types  //////////////////////////////////
type Gear struct {
	damage     int
	defense    int
	recovery   int
	durability int
	gearType   string
	name       string
}

///////////////////////////////  Set Attributes  //////////////////////////////////

func (g *Gear) SetAttributesGear() {
	var gearName = g.name

	switch g.gearType {
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
	var gear *Gear = new(Gear)

	switch gearClass {
	case "Weapon":
	case "Armor":
	case "Accessoire":
		gear.gearType = gearClass
	default:
		gear.gearType = "empty"
		gearName = "Not Filled"
	}

	gear.name = gearName

	//////////////////////// Set Gear Attributes ////////////////////

	gear.SetAttributesGear()

	//////////////////////// Return die Waffe /////////////////////

	return gear
}
