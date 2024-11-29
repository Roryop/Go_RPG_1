package gear

///////////////////////////////  Struct Gear Types  //////////////////////////////////
type Gear struct {
	recovery int

	damage     int
	defense    int
	durability int
	gearType   string
	name       string
}

type Weapon struct {
	Gear
}

type Armor struct {
	Gear
}

type Accessoire struct {
	Gear
}

///////////////////////////////  Set Attributes  //////////////////////////////////

func (g *Gear) SetAttributesGear(gearName string) {
	switch g.gearType {
	case "Weapon":
		g.defense = 0
	case "Armor":
		g.damage = 0
	default:
		break
	}

	switch gearName {
	case "Lesser Dagger":
		g.damage = 1
		g.durability = 10

	case "Leather Breastplate":
		g.defense = 2
		g.durability = 9
	}
}

///////////////////////////////  Set Accessoire Attributes  //////////////////////////////////

func (a *Accessoire) SetAttributesAccessoire(accessoireName string) {

	a.name = accessoireName
	switch a.name {

	case "Lesser Ring of Strength":

		a.damage = 1

	case "Lesser Ring of Resilience":

		a.defense = 1

	case "Lesser Ring of Recovery":

		a.recovery = 1

	}

	a.durability = -1 //Equipment is destroyed at 0 durability, so Accessoires have -1
}

///////////////////////////////  Create New Gear  //////////////////////////////////

func NewWeapon() *Weapon {
	var weapon = new(Weapon)
	weapon.gearType = "Weapon"
	return weapon
}

func NewArmor() *Armor {
	var armor = new(Armor)
	armor.gearType = "Armor"
	return armor
}

func NewAccessoire() *Accessoire {
	var accessoire = new(Accessoire)
	accessoire.gearType = "Accessoire"
	return accessoire
}
