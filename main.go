package main

import (
	"fmt"
	"math/rand"
	"start/enemy"
	"start/gear"
	"start/player"
	"start/story"
)

/*
	func helloWelt() (string, string) {

		var x, y string
		x, y = "Hello", "World"
		return x, y

	}
*/

func main() {

	story.Prologue()
	var b = gear.NewGear("Accessoire", "Lesser Ring of Strength")
	fmt.Println(b)
	/////////////Testing Weapon///////////
	var inventory = gear.NewInventory()
	fmt.Println(inventory)

	var inventorySlot = gear.NewInventorySlot()
	inventorySlot.InputInventorySlot(b, 1)
	inventory[0] = inventorySlot
	fmt.Println(inventory[1])
	gear.GiveInventoryInformation(inventory)

	////////////////////////End Weapon Test///////////////////////

	/*
		var x, y string
		x, y = helloWelt()

		fmt.Println(x, y)
	*/

	//kreiere Spieler

	var player1 = player.BeginPlayer()
	fmt.Println("Player: ", player1)
	player1.SeeStats()
	/*var player1 = player.InitPlayer()

	var pStats [6]int = player1.GetStats() //Create Stats-array + give lvl1 stats

	player1.SeeStats(pStats) //Give out Stats on normal lvl 1 so player can better choose where to invest

	player1.UpdateSpStats() //Updates SpStats

	pStats = player1.GetStats()

	player1.SeeStats(pStats) //Give out Stats after distributing bonusPoint
	//Für HP-Rechnung vom Player:
	*/
	var hp int = player1.GetStat(2)
	var att int = player1.GetStat(3)
	var def int = player1.GetStat(4)

	//das Game_Level auswählen
	var choice int = 0
	for choice != 3 {

		player1.SetStats()
		fmt.Println("Deine Stats nach Levelbeginn sind: ")
		player1.SeeStats()

		var enemy_level int
		hp = player1.GetStat(2) //Healing Player before going into Level

		for i := 0; i < 10 && hp > 0; i++ {
			enemy_level = enemy.SetEnemyLevel()

			//Entscheidet in jedem Durchlauf über Gegner-typ + kreiert stats
			var typ int = rand.Intn(2)
			var enemyStats [4]int
			var enemy_name string

			switch typ {
			case 0:
				var gegner = enemy.NewOrk()
				enemy_name = gegner.GetOrkName()
				enemyStats = gegner.GetStatsEnemy(enemy_level)
			case 1:
				var gegner = enemy.NewWolf()
				enemy_name = gegner.GetWolfName()
				enemyStats = gegner.GetStatsEnemy(enemy_level)
			}

			fmt.Println("Du fightest einen", enemy_name+"!!!")
			fmt.Println("Er hat", enemyStats[1], "HP!")
			fmt.Println("Du hast", hp, "HP!")

			//Fight machen
			for enemyStats[1] > 0 && hp > 0 {

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
					enemyStats[1] = enemyStats[1] - ((att * 100) / (100 + enemyStats[3]))
					fmt.Println("Der", enemy_name, "hat noch", enemyStats[1], "HP")
					//Enemy attacks Player
					if enemyStats[1] > 0 {
						hp = hp - ((enemyStats[2] * 100) / (100 + def))
					}
					fmt.Println("Du hast noch", hp, "HP")

				case 2:

					fmt.Println("Du hast nicht angegriffen")
					//Enemy attacks Player
					hp = hp - ((enemyStats[2] * 100) / (100 + def))
					fmt.Println("Du hast noch", hp, "HP")

				default:
					goto end
				}
			}
			if enemyStats[1] <= 0 {
				player1.CalculateExp((enemyStats[1] + enemyStats[2] + enemyStats[3] + enemyStats[0]) / 4)
				player1.Level_Management()
				player1.SetStats()
				//Spieler will not be healed after fight
				fmt.Println("Deine Stats sind jetzt:")
				fmt.Println("HP:", hp)
				fmt.Println("Att:", att)
				fmt.Println("Def:", def)
			}
		}
	end:
		if choice == 3 {
			break
		}
	}
}
