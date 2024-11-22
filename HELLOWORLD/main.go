package main

import (
	"fmt"
	"math/rand"
	"start/enemy"
	"start/player"
)

/*
	func helloWelt() (string, string) {

		var x, y string
		x, y = "Hello", "World"
		return x, y

	}
*/

func main() {

	/*
		var x, y string
		x, y = helloWelt()

		fmt.Println(x, y)
	*/

	//kreiere Spieler
	var player1 = player.InitPlayer()

	var pStats [6]int = player1.GetStats() //Create Stats-array + give lvl1 stats

	fmt.Println("Deine Stats sind jetzt:") //Give out Stats on normal lvl 1 so player can better choose where to invest
	fmt.Println("HP:", pStats[2])
	fmt.Println("Att:", pStats[3])
	fmt.Println("Def:", pStats[4])

	player1.UpdateSpStats() //Updates SpStats

	pStats = player1.GetStats()

	fmt.Println("Deine Stats sind jetzt:") //Give out Stats after distributing bonusPoint
	fmt.Println("HP:", pStats[2])
	fmt.Println("Att:", pStats[3])
	fmt.Println("Def:", pStats[4])

	//Für HP-Rechnung vom Player:

	var hp int

	//das Game_Level auswählen
	var game_level int
	var choice int = 0
	for choice != 3 {

		pStats = player1.GetStats()

		fmt.Println("In welches Level möchtest du?")
		fmt.Scanln(&game_level)
		fmt.Println("Du hast Level", game_level, "ausgewählt.")

		var enemy_level int

		for i := 0; i < 10 || pStats[2] <= 0; i++ {

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
			}

			//Entscheidet in jedem Durchlauf über Gegner-typ + kreiert stats
			var typ int = rand.Intn(2)
			var enemyStats [4]int
			var enemy_name string
			if typ == 0 {
				var gegner = enemy.NewOrk()
				enemy_name = gegner.GetOrkName()
				enemyStats = gegner.GetStatsOrk(enemy_level)
			} else if typ == 1 {
				var gegner = enemy.NewWolf()
				enemy_name = gegner.GetWolfName()
				enemyStats = gegner.GetStatsWolf(enemy_level)
			}

			fmt.Println("Du fightest einen", enemy_name+"!!!")
			fmt.Println("Er hat", enemyStats[1], "HP!")

			//Fight machen
			for enemyStats[1] > 0 || pStats[2] > 0 {
				//Anfrage wegen Angriff
				fmt.Println("Möchtest du angreifen?")
				fmt.Println("1: Ja")
				fmt.Println("2: Nein")
				fmt.Println("3: End Game")
				fmt.Scanln(&choice)
				fmt.Println(choice)

				//Angreifen
				switch choice {
				case 1:
					//Player attacks Enemy
					enemyStats[1] = enemyStats[1] - ((pStats[3] * 100) / (100 + enemyStats[3]))
					fmt.Println("Der", enemy_name, "hat noch", enemyStats[1], "HP")
					//Enemy attacks Player
					pStats[2] = pStats[2] - ((enemyStats[2] * 100) / (100 + pStats[4]))
					hp = pStats[2]
					fmt.Println("Du hast noch", pStats[2], "HP")
				case 2:
					fmt.Println("Du hast nicht angegriffen")
				default:
					goto end
				}
			}
			if enemyStats[1] <= 0 {
				var EnemyEXP int = player1.GetEnemyExpValue(((enemyStats[1] + enemyStats[2] + enemyStats[3]) + enemyStats[0]) / 5)
				player1.EXP_Function(EnemyEXP)
				player1.Level_Management()
				pStats = player1.GetStats() //Updates Stats so stat SpStats can be added
				pStats[2] = pStats[2] - hp  //Spieler will not be healed after fight
				player1.UpdateSpStats()
				pStats = player1.GetStats() //Updates Stats so that correct values can be given out
				pStats[2] = pStats[2] - hp  //Need to be in double
				fmt.Println("Deine Stats sind jetzt:")
				fmt.Println("HP:", pStats[2])
				fmt.Println("Att:", pStats[3])
				fmt.Println("Def:", pStats[4])
			}
			if pStats[2] <= 0 {
				goto noHP
			}
		noHP:
			if pStats[2] <= 0 {
				break
			}
		}
	end:
		if choice == 3 {
			break
		}
	}
}
