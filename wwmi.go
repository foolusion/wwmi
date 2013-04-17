package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// The master schedule
var schedule []Game

// The master team list
var teams map[string]Team

// A specific game
type Game struct {
	Date time.Time
	Away string
	Home string
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

func init() {
	schedule = make([]Game, 0, 100)
	teams = make(map[string]Team)
}

func main() {
	fmt.Printf("Buffalo will make the playoffs\n")
	err := readSchedule("schedule.dat")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range schedule {
		if v.Date.After(time.Now()) {
			fmt.Printf("%v at %v, %v\n", v.Away, v.Home, v.Date)
		}
	}
}

func readSchedule(filename string) error {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	done := false
	for !done {
		line, err := reader.ReadString('\n')
		if err != nil {
			done = true
		}

		// Line with less than 2 characters should be skipped
		if len(line) < 2 {
			continue
		}

		// remove the '\n' from the end of the line
		line = line[:len(line)-1]

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
		schedule = append(schedule, game)
	}
	return err
}

func readTeams(filename string) error {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	done := false
	for !done {
		line, err := reader.ReadString('\n')
		if err != nil {
			done = true
		}

		// Line with less than 2 characters should be skiped
		if len(line) < 2 {
			continue
		}

		// remove the '\n' from the end of the line
		line = line[:len(line)-1]

		// Split line into team values
		teamString := strings.Split(line, "\t")

		team := Team{}
		team.Name = teamString[0]
		team.RegulationWins, err = strconv.Atoi(teamString[1])
		if err != nil {
			return err
		}
		team.OvertimeWins, err = strconv.Atoi(teamString[2])
		if err != nil {
			return err
		}
		team.ShootoutWins, err = strconv.Atoi(teamString[3])
		if err != nil {
			return err
		}
		team.RegulationLosses, err = strconv.Atoi(teamString[4])
		if err != nil {
			return err
		}
		team.OvertimeLosses, err = strconv.Atoi(teamString[5])
		if err != nil {
			return err
		}
		team.ShootoutLosses, err = strconv.Atoi(teamString[6])
		if err != nil {
			return err
		}
		team.Points, err = strconv.Atoi(teamString[7])
		if err != nil {
			return err
		}
		teams[team.Name] = team
	}
	return err
}
