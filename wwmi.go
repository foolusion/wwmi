package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// The master schedule
var schedule []Game

// A specific game
type Game struct {
	Date time.Time
	Away *Team
	Home *Team
	Result
}

// A representation of a team.
type Team struct {
	Name             string
	RegulationWins   int
	OvertimeWins     int
	ShootoutWins     int
	RegulationLosses int
	OvertimeLosses   int
	ShootoutLosses   int
	Points           int
}

// The results of a game.
// Only one of RegulationWin, OvertimeWin, and ShootoutWIn can be true.
type Result struct {
	HomeWin       bool
	RegulationWin bool
	OvertimeWin   bool
	ShootoutWin   bool
}

func main() {
	fmt.Printf("Buffalo will make the playoffs\n")
}

func readSchedule(filename string) error {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Create a buffered reader
	reader := bufio.NewReader(file)

	for err != nil {
		line, err := reader.ReadLine('\n')

		game := Game{}
		// Split line into game values
		gameString := strings.Split(line, "\t")
		// Parse the date
		game.Date, err = time.Parse("Mon Jan 2 2006", gameString[0])
		if err != nil {
			return err
		}
		// Add teams to game
		game.Away = gameString[1]
		game.Home = gameString[2]

		// Add game to Schedule
		append(schedule, game)
	}
}
