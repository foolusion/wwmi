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
var games []Game

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
	Division         string
	Conference       string
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
	games = make([]Game, 0, 100)
	teams = make(map[string]Team)
}

func main() {
	fmt.Printf("Buffalo will make the playoffs\n")
	err := readSchedule("schedule.dat")
	if err != nil {
		fmt.Println(err)
	}
	err = readTeams("teams.dat")
	if err != nil {
		fmt.Println(err)
	}

	WillWeMakeIt("buffalo")
}

// a star?
// for each game
// 	add all possibilities to open list
//	sort them based on score
//	pick best
//	repeat until goal is met
// GOAL:	* in 8th place or better
//		* for all teams in conference lower
//			* games left * max points < wwmi points
func WillWeMakeIt(team string) {
	wwmi := teams[team]

	if wwmi.Conference == "east" {
		removeTeamsByConference("west")
		removeGamesByConference("west")
	} else {
		removeTeamsByConference("east")
		removeGamesByConference("east")
	}

	fmt.Println("NO!!!!")

	for _, v := range teams {
		fmt.Printf("%v, %v\n", strings.Title(v.Name), v.Points)
	}

	for _, g := range games {
		if g.Date.After(time.Now()) {
			fmt.Printf("%v at %v, %v\n", strings.Title(g.Away), strings.Title(g.Home), g.Date)
		}
	}
}

func removeTeamsByConference(conf string) {
	for i := range teams {
		if teams[i].Conference == conf {
			delete(teams, i)
		}
	}
}

func removeGamesByConference(conf string) {
	for i := 0; i < len(games); {
		if v, ok := teams[games[i].Away]; !ok {
			games = append(games[:i], games[i+1:]...)
		} else if v.Conference == conf {
			fmt.Println(games[i])
			games = append(games[:i], games[i+1:]...)
		} else {
			i++
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
		game.Away = strings.ToLower(gameString[1])
		game.Home = strings.ToLower(gameString[2])

		// Add game to Schedule
		games = append(games, game)
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
		team.Name = strings.ToLower(teamString[0])
		team.Division = strings.ToLower(teamString[1])
		team.Conference = strings.ToLower(teamString[2])
		team.RegulationWins, err = strconv.Atoi(teamString[3])
		if err != nil {
			return err
		}
		team.OvertimeWins, err = strconv.Atoi(teamString[4])
		if err != nil {
			return err
		}
		team.ShootoutWins, err = strconv.Atoi(teamString[5])
		if err != nil {
			return err
		}
		team.RegulationLosses, err = strconv.Atoi(teamString[8])
		if err != nil {
			return err
		}
		team.OvertimeLosses, err = strconv.Atoi(teamString[7])
		if err != nil {
			return err
		}
		team.ShootoutLosses, err = strconv.Atoi(teamString[6])
		if err != nil {
			return err
		}
		team.Points = team.RegulationWins*2 + team.OvertimeWins*2 + team.ShootoutWins*2 + team.ShootoutLosses + team.OvertimeLosses
		teams[team.Name] = team
	}
	return err
}
