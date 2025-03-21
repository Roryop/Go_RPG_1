package story

import (
	"fmt"
	"start/enemy"
	"start/gear"
	"start/player"
	"start/text"
)

func Prologue() {
	text.EmptyLine()
	text.Print("Du wachst im Raum auf.")
	text.Print("Du spührst weder Kälte noch Wärme...")
	text.Print("Siehst weder Licht noch Schatten...")
	text.Print("Hast keinen Boden unter den Füßen und dennoch...")
	text.Print("Spührst keine Bewegung.")
	text.EmptyLine()
	text.ShortWait()
	text.Print("Du weißt nicht wer du bist...")
	text.Print("Du weißt nicht woher du kommst...")
	text.Print("Und du weißt nicht was du in deiner Verganheit hintergelassen hast...")
	text.Pause()
	text.Print("Finde dich wieder auf dem Weg nach Macht und Stärke.")
	text.Print("Denn das das einzige ist was du noch im Kopf übrig hast.")
	text.ShortWait()
	text.Space(20)
	text.ShortWait()
	text.Print("Schmiede dein Schicksal, Spieler.")
	text.ShortWait()
	text.Pause()
	text.Space(30)
}

// Gets Player, Inventory, current stats, world, world_barrier, typ of enemy
// Creates Enemy, Fights against enemy; Updates player Stats
// Returns enemyStats, new current Stats
func Fight(player1 *player.Player, inventory [10]*gear.InventorySlot, hp, att, def, rec int, world string, player_level int, typ int) (*player.Player, [10]*gear.InventorySlot, int, int, int, int, int) {

	var enemyName, enemyStats = enemy.CreateEnemy(world, player_level, typ)

	text.Print("Du fightest einen " + enemyName + "!!!")
	text.Print("Er ist Level " + fmt.Sprint(enemyStats[0]) + "!!!")
	text.Print("Er hat " + fmt.Sprint(enemyStats[1]) + "HP!")
	text.Print("Du hast " + fmt.Sprint(hp) + "HP!")

	for enemyStats[1] > 0 && hp > 0 {
		var choice int

		//Anfrage wegen Angriff
		text.Print("Möchtest du angreifen?")
		fmt.Println("1: Ja")
		fmt.Println("2: Nein")
		fmt.Println("3: Eigene Stats sehen")
		fmt.Println("4: Inventory sehen")
		fmt.Println("5: Wie ein Feigling wegrennen!")
		fmt.Scanln(&choice)

		//Angreifen
		switch choice {
		case 1:
			//Player attacks Enemy
			enemyStats[1] = enemyStats[1] - ((att * 100) / (100 + enemyStats[3]))
			if enemyStats[1] > 0 {
				text.Print(fmt.Sprint(enemyName) + " hat noch " + fmt.Sprint(enemyStats[1]) + "HP")
			} else {
				text.Print(fmt.Sprint(enemyName) + " wurde besiegt")
			}
			//Enemy attacks Player
			if enemyStats[1] > 0 {
				hp = hp - ((enemyStats[2] * 100) / (100 + def))
			}

			if hp > 0 {
				text.Print("Du hast noch " + fmt.Sprint(hp) + "HP")
			}

		case 2:
			text.Print("Du hast nicht angegriffen")
			//Enemy attacks Player
			hp = hp - ((enemyStats[2] * 100) / (100 + def))
			if hp > 0 {
				text.Print("Du hast noch " + fmt.Sprint(hp) + "HP")
			}
		case 3:
			player1.SeePlayerStats(inventory, hp, att, def, rec)
		case 4:
			gear.GiveInventoryInformation(inventory)
		case 5:
			goto end
		}
	}
	///////////////////////// Case Enemy Died ///////////////////////////
	if enemyStats[1] <= 0 {
		//////////////// Enemy Item Drop /////////////
		inventory = gear.AddDropToInventory(inventory, player_level)

		///////////////// Player Management ////////////////
		player1.Exp_Function(enemyStats)                                           // Giving Exp to Player
		hp, att, def, rec = player1.Level_Management(inventory, hp, att, def, rec) // Player will be healed with levelUp + Updating Stats + Updating current Stats
		player_level = player1.GetLevel()                                          // Updating player_level needed for other functions
	}

	//////////////// Clearing Inventory on Player Death ////////////////
	if hp <= 0 {
		text.Print("Du bist gestorben!") // Death message
		inventory = gear.FillEmptyInventory(inventory)
	}
end:
	return player1, inventory, hp, att, def, rec, player_level
}

func Ende(player1 *player.Player) {
	var karma = player1.GetKarma()

	text.Space(50)
	text.ShortWait()
	text.Print("Deine Reise trifft ein Ende.")
	text.ShortWait()
	text.Print("Doch waren deine Taten moralisch vertretbar?")
	text.LongWait()
	text.Space(50)

	switch {
	case karma == 0: // morally neutral ending
		text.Print("Du hast deine Reise auf der Schneide zwischen Gut und Schlecht beendet.")
		text.ShortWait()
		text.Print("Perfectly balanced, as all should be...")
	case karma > 0 && karma <= 30: // morally good ending
		text.Print("Du hast deine Reise auf der besseren Seite beendet.")
		text.ShortWait()
		text.Print("Du fühlst dich zufrieden...")
	case karma > 30: // morally best ending
		text.Print("Du hast deine Reise auf der besten Seite beendet.")
		text.ShortWait()
		text.Print("Du fühlst dich gelassen...")
	case karma < 0 && karma >= -30: // morally bad ending
		text.Print("Du hast deine Reise auf der schlechteren Seite beendet.")
		text.ShortWait()
		text.Print("Du bereust einiges...")
	case karma > 30: // morally worst ending
		text.Print("Du hast deine Reise auf der bösesten Seite beendet.")
		text.ShortWait()
		text.Print("Du fühlst dich gut im Angesicht deiner Gräueltaten...")
	}

	text.ShortWait()
	text.Space(100)
	text.Print("Vielen Dank fürs Spielen")
	text.Print("Credits to Benjamin(40%), Martigga(25%), Charlee(20%), Aleksis(15%)")
	text.LongWait()
}
