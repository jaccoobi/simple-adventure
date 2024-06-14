package main

import (
	"fmt"
	"math/rand/v2"
)

var (
	roomsExplored = 1
	treasureCount = 0
	healthPoints  = 5
	escaped       = false
	monster       = false
	currentRoom   = createRoom()
)

func rollDice(num int, size int) int {
	total := 0
	for i := 0; i < num; i++ {
		total += rand.IntN(size+1) + 1
	}
	return total
}

func hasMonster() bool {
	if rollDice(2, 6) >= 8 {
		return true
	} else {
		return false
	}
}

func hasEscaped() bool {
	if rollDice(2, 6) >= 11 {
		return true
	} else {
		return false
	}
}

func monsterAttack() bool {
	if rollDice(2, 6) >= 9 {
		return true
	} else {
		return false
	}
}

func defeatMonster() bool {
	if rollDice(2, 6) >= 4 {
		return true
	} else {
		return false
	}
}

func hasTreasure() bool {
	if rollDice(2, 6) >= 8 {
		return true
	} else {
		return false
	}
}

func treasure() string {
	t := []string{
		"gold coins",
		"gems",
		"a magic wand",
		"an enchanted sword",
	}

	return t[rand.IntN(4)]
}

func roomSize() string {
	t := []string{
		"huge",
		"large",
		"big",
		"regular",
		"small",
		"tiny",
	}

	return t[rand.IntN(6)]
}

func roomColor() string {
	t := []string{
		"red",
		"blue",
		"green",
		"dark",
		"golden",
		"crystal",
	}

	return t[rand.IntN(6)]
}

func roomType() string {
	t := []string{
		"cave",
		"treasure room",
		"rock cavern",
		"tomb",
		"guard room",
		"lair",
	}

	return t[rand.IntN(6)]
}

func roomDirection() string {
	t := []string{
		"north",
		"south",
		"east",
		"west",
	}

	return t[rand.IntN(4)]
}

func createRoom() string {
	return fmt.Sprintf("You are in a %s %s %s.\nThere is an exit on the %s wall.", roomSize(), roomColor(), roomType(), roomDirection())
}

func playerLoop() {
 	fmt.Println("You are trapped in the dungeon.")
	fmt.Println("Collect treasure and try to escape.")
	fmt.Println("To play, type one of the command choices on each turn.")
	fmt.Println()

	actions := map[string]string{
		"m": "move",
		"s": "search",
	}
	// Main game loop
	for healthPoints > 0 && !escaped {

		fmt.Printf("Room number %d\n", roomsExplored)
		fmt.Println(currentRoom)

		fmt.Println()

		if monster {
			fmt.Println("OH NO! An evil monster is in here with you!")
			actions["f"] = "fight"
		} else {
			delete(actions, "f")
		}
		fmt.Println()

		for k, v := range actions {
			fmt.Printf("%s - %s\n", k, v)
		}

		fmt.Println("What do you do?")

		var input string
		fmt.Scanf("%s", &input)

		fmt.Println()

		if monster && monsterAttack() {
			healthPoints -= 1
			fmt.Println("OUCH! The monster hit you!")
		}

		switch input {
		case "m":
			currentRoom = createRoom()
			roomsExplored += 1
			monster = hasMonster()
			escaped = hasEscaped()
		case "s":
			if hasTreasure() {
				fmt.Printf("You found %s!\n", treasure())
				treasureCount += 1
			} else {
				fmt.Println("You look, but don't find anything.")
			}

			if !monster {
				monster = hasMonster()
			}
		case "f":
			if defeatMonster() {
				monster = false
				fmt.Println("You defeated the monster!")
			} else {
				fmt.Println("You attack and miss!")
			}
		default:
			fmt.Println("I don't know ho to do that!")
			fmt.Println()
		}

		fmt.Println()
	}

	if healthPoints > 0 {
		fmt.Printf("You escaped!\n")
		fmt.Printf("You explored %d rooms.\n", roomsExplored)
		fmt.Printf("You found %d treasures.\n", treasureCount)
	} else {
		fmt.Printf("OH NO! You didn't make it out.\n")
		fmt.Printf("You explored %d rooms before meeting your doom.\n", roomsExplored)
	}
}

func aiLoop()  {
  
}

func main() {
  fmt.Println("Welcome to a simple adventure")
  fmt.Println("Do you want to play the game")
  fmt.Println("or do you want to watch the computer play.")
  fmt.Println()
  fmt.Println("p - player")
  fmt.Println("c - computer")
  fmt.Println()

  fmt.Println("What do you want to do?")
  var i string
  fmt.Scanf("%s", &i)

  fmt.Println()

  if i == "p" {
    playerLoop()
  } else if i == "c" {
    aiLoop()
  } else {
    fmt.Println("I don't know how to do that.")
  }
}
