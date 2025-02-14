package enemy

import (
	"math/rand"
)

// Erstes Wesen
type Wesen struct {
	level int
	hp    int
	att   int
	def   int
	name  string
}

///////////////////////////Create Enemy////////////////////////////////

// Gets Game_level
// Creates Enemy depending on Game_level
// Returns Enemy Name and Enemy Stats
func CreateEnemy(game_level string) (string, [4]int) {
	var enemy_level = SetEnemyLevel(game_level)

	// Decides on Enemy Typ + creates empty Variables for it
	var typ int = rand.Intn(2)
	var enemyStats [4]int
	var enemyName string

	// Creates Stats and Name of Enemy depending on Enemy Typ + Enemy level
	var gegner = NewEnemy()
	gegner.SetEnemyTyp(typ)
	enemyStats = gegner.GetStatsEnemy(enemy_level)
	enemyName = gegner.name

	return enemyName, enemyStats
}

// Gets Enemy_level
// Creates stats based on typ + level
// Returns Stats
func (w *Wesen) GetStatsEnemy(level int) [4]int {

	w.level = level

	var stats [4]int

	switch w.name {
	case "Ork":
		w.hp = 13 + (7 * w.level)
		w.att = 4 + (3 * w.level)
		w.def = 4 + w.level
	case "Wolf":
		w.hp = 11 + (3 * w.level)
		w.att = 6 + (4 * w.level)
		w.def = 3 + (w.level / 2)
	}

	stats[0] = w.level
	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

// Gets Game_level
// Decides Enemy_level based on Game_level + rand.Intn()
// Returns Enemy_level
func SetEnemyLevel(game_level string) int {
	var enemy_level = 0

	//Entscheidet Gegner-level jedes mal neu
	switch game_level {
	case "Cyberpunk":
		//Gegnerlevel zwischen 1 und 3
		enemy_level = rand.Intn(3) + 1
	case "MiddleAges":
		//Gegnerlevel zwischen 3 und 5
		enemy_level = rand.Intn(3) + 3
	case "AndreasLand":
		//Gegnerlevel zwischen 6 und 10
		enemy_level = rand.Intn(5) + 6
	case "Schönwalde":
		//Gegnerlevel zwischen 11 und 20
		enemy_level = rand.Intn(10) + 11
	case "Muschelhausen":
		//Gegnerlevel zwischen 15 und 25
		enemy_level = rand.Intn(11) + 15
	case "Ghetto":
		//Gegnerlevel zwischen 30 und 35
		enemy_level = rand.Intn(6) + 30
	case "Staaken":
		//Gegnerlevel zwischen 34 und 36
		enemy_level = rand.Intn(3) + 34
	case "ChayaLand":
		//Gegnerlevel zwischen 36 und 40
		enemy_level = rand.Intn(5) + 36
	case "BookofRa":
		//Gegnerlevel zwischen 40 und 45
		enemy_level = rand.Intn(6) + 40
	case "HawkTuah":
		//Gegnerlevel zwischen 45 und 50
		enemy_level = rand.Intn(6) + 45
	case "Gyatt":
		enemy_level = 3001
	default:
		enemy_level = 0
	}
	return enemy_level
}

// Gets Enemy Typ
// Sets Name of enemy based on given typ
func (w *Wesen) SetEnemyTyp(typ int) {
	switch typ {
	case 1:
		w.name = "Ork"
	case 2:
		w.name = "Wolf"
	}
}

/////////////////// Creating New Enemy //////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
