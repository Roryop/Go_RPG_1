package player

import (
	"fmt"
)

type Player struct {
	level int

	//creating experience system
	experience int
	exp_limit  int

	hp          int
	att         int
	def         int
	hpPoints    int
	attPoints   int
	defPoints   int
	bonusPoints int
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

func (w *Player) GetStats() [6]int {

	//Kreiert stats Array
	var stats [6]int

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
	stats[5] = w.bonusPoints
	//gibt stats aus
	return stats
}

func (w *Player) SeeStats(pStats [6]int) {
	fmt.Println("Deine Stats sind jetzt:")
	fmt.Println("HP:", pStats[2])
	fmt.Println("Att:", pStats[3])
	fmt.Println("Def:", pStats[4])
}

// initiating EXP_Stat
func (w *Player) InitEXP() {
	w.experience = 0
}

// getting Experience_Value ready for showcasing
func (w *Player) GetEXP() int {
	return w.experience
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

func BeginPlayer() (*Player, [6]int) {
	var player1 = InitPlayer()

	var pStats [6]int = player1.GetStats() //Create Stats-array + give lvl1 stats

	player1.SeeStats(pStats) //Give out Stats on normal lvl 1 so player can better choose where to invest

	player1.UpdateSpStats() //Updates SpStats

	pStats = player1.GetStats()

	player1.SeeStats(pStats) //Give out Stats after distributing bonusPoint

	return player1, pStats
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
