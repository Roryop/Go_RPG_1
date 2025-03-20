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
func CreateEnemy(world string, player_level int, typ int) (string, [4]int) {
	var enemy_level = SetEnemyLevel(world, player_level)

	// Decides on Enemy Typ + creates empty Variables for it
	var enemyStats [4]int
	var enemyName string

	if typ == 0 {
		typ = rand.Intn(19) + 1
	}

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
	case "Baer":
		w.hp = 13 + (7 * w.level)
		w.att = 6 + (4 * w.level)
		w.def = 3 + w.level
	case "Raeuber":
		w.hp = 12 + (4 * w.level)
		w.att = 7 + (4 * w.level)
		w.def = 3 + w.level
	case "Burgritter":
		w.hp = 8 + (2 * w.level)
		w.att = 8 + (4 * w.level)
		w.def = 7 + (5 * w.level)
	case "betrunkener Dorfbewohner":
		w.hp = 5 + (2 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 6 + (7 * w.level)
	case "Mutantenratte":
		w.hp = 15 + (6 * w.level)
		w.att = 4 + (2 * w.level)
		w.def = 3 + w.level
	case "dreikoepfige Schlange":
		w.hp = 15 + (3 * w.level)
		w.att = 7 + (3 * w.level)
		w.def = 5 + (3 * w.level)
	case "laufende Makrowelle":
		w.hp = 4 + w.level
		w.att = 4 + (5 * w.level)
		w.def = 3 + (3 * w.level)
	case "Pluenderer":
		w.hp = 10 + (3 * w.level)
		w.att = 5 + (5 * w.level)
		w.def = 3 + w.level
	case "Triceratops":
		w.hp = 16 + (7 * w.level)
		w.att = 4 + (4 * w.level)
		w.def = 3 + (2 * w.level)
	case "Pterodactylus":
		w.hp = 7 + (7 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 3 + w.level
	case "prehistorische Bergsteiger Ziege":
		w.hp = 5 + (5 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 6 + (5 * w.level)
	case "Johannes Bleileber":
		w.hp = 10 + (6 * w.level)
		w.att = 6 + (6 * w.level)
		w.def = 5 + w.level
	case "Juan Pared Caballero":
		w.hp = 14 + (6 * w.level)
		w.att = 5 + (4 * w.level)
		w.def = 5 + w.level
	case "Sheriff Joe Swanson":
		w.hp = 12 + (6 * w.level)
		w.att = 5 + (4 * w.level)
		w.def = 5 + w.level
	case "Koyot":
		w.hp = 7 + (3 * w.level)
		w.att = 8 + (3 * w.level)
		w.def = 2 + w.level
	case "Penner":
		w.hp = 5 + (7 * w.level)
		w.att = 2 + (3 * w.level)
		w.def = 0 + w.level
	case "Weißer Monster":
		w.hp = 15 + (7 * w.level)
		w.att = 5 + (5 * w.level)
		w.def = 0 + w.level
	}

	stats[0] = w.level
	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

// Gets world and world_barrier
// Decides enemy_level based on world + world_barrier + rand.Intn()
// Returns enemy_level
func SetEnemyLevel(world string, player_level int) int {
	var enemy_level = 0

	// Only changes enemy_level based on world_barrier if correct world was chosen, tutorial will stay level 0
	if world == "cyberpunk" || world == "middleage" || world == "armageddon" || world == "prehistory" || world == "wildwest" {

		enemy_level = rand.Intn(5) + player_level - 1
		if enemy_level == 0 { // Making sure enemy_level is not 0 outside of Tutorial
			enemy_level = 1
		}

	}

	return enemy_level
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
		w.name = "Baer"
	case 4:
		w.name = "Raeuber"
	case 5:
		w.name = "Burgritter"
	case 6:
		w.name = "betrunkener Dorfbewohner"
	case 7:
		w.name = "Mutantenratte"
	case 8:
		w.name = "dreiköpfige Schlange"
	case 9:
		w.name = "laufende Makrowelle"
	case 10:
		w.name = "Pluenderer"
	case 11:
		w.name = "Triceratops"
	case 12:
		w.name = "Pterodactylus"
	case 13:
		w.name = "prehistorische Bergsteiger Ziege"
	case 14:
		w.name = "Johannes Bleileber"
	case 15:
		w.name = "Juan Pared Caballero"
	case 16:
		w.name = "Sheriff Joe Swanson"
	case 17:
		w.name = "Koyot"
	case 18:
		w.name = "Penner"
	case 19:
		w.name = "Weißer Monster"
	default:
		w.name = "Ork"

	}
}

/////////////////// Creating New Enemy //////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
