package main

import (
	"fmt"
	"time"
)

// A collection of games to be played
type Calendar struct {
	games []Game
}

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
