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
func CreateEnemy(world int) (string, [4]int) {
	var _, enemy_level = SetWorldEnemy(world)

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
	}

	stats[0] = w.level
	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

// Gets world
// Decides Events and Enemy_level based on world + rand.Intn()
// Returns eventArray and Enemy_level
func SetWorldEnemy(world int) ([3]string, int) {
	var eventArray [3]string
	var enemy_level = 0

	switch world {
	case 1:

		eventArray = [3]string{"Event1", "Event2", "Event3"} // Please Input Event Names
		//enemy_level between 1 and 3
		enemy_level = rand.Intn(3) + 1
	case 2:

		eventArray = [3]string{"Event1", "Event2", "Event3"}
		//enemy_level between 3 and 5
		enemy_level = rand.Intn(3) + 3
	case 3:
		eventArray = [3]string{"Event1", "Event2", "Event3"}
		//enemy_level between 6 and 10
		enemy_level = rand.Intn(5) + 6
	default:
		/*	//Gegnerlevel zwischen 11 und 20
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
		*/
		eventArray = [3]string{"Event1", "Event2", "Event3"}
		//enemy_level  0
		enemy_level = 0
	}
	return eventArray, enemy_level
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
	}
}

/////////////////// Creating New Enemy //////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
