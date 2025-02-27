package enemy

import (
	"math/rand"
	"start/game"
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
func CreateEnemy(world string, world_barrier int) (string, [4]int) {
	var _, enemy_level = game.SetWorldEnemy(world, world_barrier)

	// Decides on Enemy Typ + creates empty Variables for it
	var typ int = rand.Intn(5) + 1
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
func (w *Wesen) GetStatsEnemy(enemy_level int) [4]int {

	w.level = enemy_level

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
	case "Elf":
		w.hp = 8 + (2 * w.level)
		w.att = 6 + (4 * w.level)
		w.def = 3 + w.level
	case "Räuber":
		w.hp = 12 + (4 * w.level)
		w.att = 7 + (4 * w.level)
		w.def = 3 + w.level
	case "E-Gigant":
		w.hp = 8 + (1 * w.level)
		w.att = 8 + (4 * w.level)
		w.def = 5 + (5 * w.level)
	}

	stats[0] = w.level
	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

// Gets Enemy Typ
// Sets Name of enemy based on given typ
// Returns Nothing
func (w *Wesen) SetEnemyTyp(typ int) {
	switch typ {
	case 1:
		w.name = "Ork"
	case 2:
		w.name = "Wolf"
	case 3:
		w.name = "Elf"
	case 4:
		w.name = "Räuber"
	case 5:
		w.name = "E-Gigant"
	default:
		w.name = "Ork"
	}
}

/////////////////// Creating New Enemy //////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
