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

func CreateEnemy(game_level int) (string, [4]int) {
	var enemy_level = SetEnemyLevel(game_level)

	//Entscheidet in jedem Durchlauf über Gegner-typ + kreiert stats
	var typ int = rand.Intn(2)
	var enemyStats [4]int
	var enemyName string

	var gegner = NewEnemy()
	gegner.SetEnemyType(typ)
	enemyStats = gegner.GetStatsEnemy(enemy_level)
	enemyName = gegner.name
	return enemyName, enemyStats
}

// Set und Get Funktionen für Attribute von Wesen
func (w *Wesen) GetStatsEnemy(level int) [4]int {

	w.level = level
	var Level int = w.level

	var stats [4]int
	stats[0] = Level

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

	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

func SetEnemyLevel(game_level int) int {
	var enemy_level = 0

	//Entscheidet Gegner-level jedes mal neu
	switch game_level {
	case 1:
		//Gegnerlevel zwischen 1 und 3
		enemy_level = rand.Intn(3) + 1
	case 2:
		//Gegnerlevel zwischen 3 und 5
		enemy_level = rand.Intn(3) + 3
	case 3:
		//Gegnerlevel zwischen 6 und 10
		enemy_level = rand.Intn(5) + 6
	case 4:
		//Gegnerlevel zwischen 11 und 20
		enemy_level = rand.Intn(10) + 11
	case 5:
		//Gegnerlevel zwischen 15 und 25
		enemy_level = rand.Intn(11) + 15
	case 6:
		//Gegnerlevel zwischen 30 und 35
		enemy_level = rand.Intn(6) + 30
	case 7:
		//Gegnerlevel zwischen 34 und 36
		enemy_level = rand.Intn(3) + 34
	case 8:
		//Gegnerlevel zwischen 36 und 40
		enemy_level = rand.Intn(5) + 36
	case 9:
		//Gegnerlevel zwischen 40 und 45
		enemy_level = rand.Intn(6) + 40
	case 10:
		//Gegnerlevel zwischen 45 und 50
		enemy_level = rand.Intn(6) + 45
	case 3001:
		enemy_level = 3001
	default:
		enemy_level = 0
	}
	return enemy_level
}

// Get Name of enemy
func (w *Wesen) SetEnemyType(typ int) {
	switch typ {
	case 1:
		w.name = "Ork"
	case 2:
		w.name = "Wolf"
	}
}

///////////////////Enemy wird erstellt//////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
