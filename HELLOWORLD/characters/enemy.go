package enemy

//Erstes Wesen
type Wesen struct {
	level int
	hp    int
	att   int
	def   int
	name  string
}

type Ork struct {
	Wesen
}

type Wolf struct {
	Wesen
}

//Set und Get Funktionen für Attribute von Wesen
//Ork GetStats
func (w *Ork) GetStatsOrk(level int) [4]int {

	w.level = level
	var Level int = w.level

	var stats [4]int
	stats[0] = Level

	w.hp = 13 + (7 * w.level)
	w.att = 4 + (3 * w.level)
	w.def = 4 + w.level

	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

//Wolf GetStats
func (w *Wolf) GetStatsWolf(level int) [4]int {

	w.level = level
	var Level int = w.level

	var stats [4]int
	stats[0] = Level

	w.hp = 11 + (3 * w.level)
	w.att = 6 + (4 * w.level)
	w.def = 3 + (w.level / 2)

	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

//Get Name of enemy
func (w *Ork) GetOrkName() string {
	w.name = "Ork"
	return w.name
}

func (w *Wolf) GetWolfName() string {
	w.name = "Wolf"
	return w.name
}

//Ork wird erstellt

func NewOrk() *Ork {
	var ork *Ork = new(Ork)
	return ork
}

//Wolf wird erstellt

func NewWolf() *Wolf {
	var wolf *Wolf = new(Wolf)
	return wolf
}