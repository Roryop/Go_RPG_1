package player

import (
	"fmt"
	"start/gear"
)

type Player struct {
	level int

	//creating experience system
	experience int
	exp_limit  int

	hp          int
	att         int
	def         int
	rec         int
	hpPoints    int
	attPoints   int
	defPoints   int
	bonusPoints int

	stats [7]int
}

func (w *Player) levelUp() int {
	w.level = w.level + 1
	w.experience = 0
	w.exp_limit = w.exp_limit * 2

	fmt.Println("Dein Level ist jetzt", w.level, "!")

	return w.level
}

func (w *Player) Level_Management() {

	if w.level >= 1 && w.experience >= w.exp_limit {
		w.level = w.levelUp()
		w.SetStats()
	} else if w.level >= 1 && w.experience < w.exp_limit {
		var b int = w.exp_limit - w.experience
		fmt.Println("Du brauchst noch", b, "Exp")
	} else {
		w.level = 1
		w.exp_limit = 5
		w.InitEXP()
		w.InitSpStats()
	}
	w.UpdateSpStats()
}

// initializing Special Stats
func (w *Player) InitSpStats() {
	w.hpPoints = 0
	w.attPoints = 0
	w.defPoints = 0
	w.bonusPoints = 1
}

func (w *Player) UpdateSpStats() {
	var allocation int

	if w.bonusPoints >= 1 {
		fmt.Println("Du hast gerade", w.bonusPoints, "Bonuspunkt.")
		fmt.Println("Bitte lege deine Bonus-Punkte fest:")
		fmt.Println("1: HP")
		fmt.Println("2: Att")
		fmt.Println("3: Def")
		fmt.Scanln(&allocation)

		switch allocation {
		case 1:
			w.hpPoints = w.hpPoints + 1
		case 2:
			w.attPoints = w.attPoints + 1
		case 3:
			w.defPoints = w.defPoints + 1

		}
		w.bonusPoints = w.bonusPoints - 1
	}
}

func (w *Player) SetStats() {

	//Kreiert stats Array
	var stats [7]int

	stats[0] = w.level
	stats[1] = w.exp_limit
	//erstellt Regeln für Stats
	w.hp = 10 + (3 * w.level) + (3 * w.hpPoints)
	w.att = 5 + w.level + w.attPoints
	w.def = 3 + w.level + w.defPoints

	w.bonusPoints = w.level - w.hpPoints - w.attPoints - w.defPoints
	//packt stats in array
	stats[2] = w.hp
	stats[3] = w.att
	stats[4] = w.def
	stats[5] = w.rec
	stats[6] = w.bonusPoints
	//gibt stats aus
	w.stats = stats
}

func (w *Player) SeePlayerStats() {
	fmt.Println("Deine Stats sind jetzt:")
	fmt.Println("HP:", w.stats[2])
	fmt.Println("Att:", w.stats[3])
	fmt.Println("Def:", w.stats[4])
	fmt.Println("Rec:", w.stats[5])
}

func (w *Player) GetStat(i int) int {
	return w.stats[i]
}

func (w *Player) SetStatsAccessoires(inventory [10]*gear.InventorySlot) {
	var att, def, rec int = gear.CreateStatsAccessoires(inventory)
	w.stats[3], w.stats[4], w.stats[5] = w.stats[3]+att, w.stats[4]+def, w.stats[5]+rec
}

// initiating EXP_Stat
func (w *Player) InitEXP() {
	w.experience = 0
}

// getting Experience_Value ready for showcasing
func (w *Player) GetEXP() int {
	return w.experience
}

// calculating experience and adding it to player
func (w *Player) CalculateExp(value int) {
	var exp = w.GetEnemyExpValue(value)
	w.EXP_Function(exp)
}

// creating funtion to get EXP_Value from enemies
func (w *Player) GetEnemyExpValue(value int) int {
	var EXP_Value int = value
	return EXP_Value
}

// creating EXP_Function needed for correct LevelUp
func (w *Player) EXP_Function(value int) int {
	w.experience = w.experience + value

	return w.experience
}

func BeginPlayer() *Player {
	var player1 = InitPlayer()

	player1.SetStats() //give lvl1 stats

	player1.SeePlayerStats() //Give out Stats

	return player1
}
func InitPlayer() *Player {
	var player = NewPlayer()
	player.InitSpStats()
	player.Level_Management()

	return player
}

func NewPlayer() *Player {
	var player *Player = new(Player)
	return player
}
