package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Result map[string]int

type Team struct {
	Name   string
	Result Result
}

// Tally processes the input from the reader and writes the tournament results to the writer.
func Tally(reader io.Reader, writer io.Writer) error {
	res := Result{
		"MP": 0,
		"W":  0,
		"D":  0,
		"L":  0,
		"P":  0,
	}
	teams := make(map[string]*Team)

	var team1, team2, outcome string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.Split(line, ";")

		if len(parts) != 3 {
			return fmt.Errorf("invalid line: %s", line)
		}

		team1, team2, outcome = parts[0], parts[1], parts[2]

		if _, exists := teams[team1]; !exists {
			teams[team1] = &Team{Name: team1, Result: make(Result)} // Initialize the Result map
			for k, v := range res {
				teams[team1].Result[k] = v // Initialize with zero values
			}
		}
		if _, exists := teams[team2]; !exists {
			teams[team2] = &Team{Name: team2, Result: make(Result)}
			for k, v := range res {
				teams[team2].Result[k] = v // Initialize with zero values
			}
		}

		teams[team1].Result["MP"]++
		teams[team2].Result["MP"]++

		switch outcome {
		case "win":
			teams[team1].Result["W"]++
			teams[team1].Result["P"] += 3
			teams[team2].Result["L"]++
		case "loss":
			teams[team2].Result["W"]++
			teams[team2].Result["P"] += 3
			teams[team1].Result["L"]++
		case "draw":
			teams[team1].Result["D"]++
			teams[team1].Result["P"]++
			teams[team2].Result["D"]++
			teams[team2].Result["P"]++
		default:
			return fmt.Errorf("invalid outcome: %s", outcome)
		}
	}

	var teamList []*Team
	for _, team := range teams {
		teamList = append(teamList, team)
	}

	sort.Slice(teamList, func(i, j int) bool {
		if teamList[i].Result["P"] == teamList[j].Result["P"] {
			return teamList[i].Name < teamList[j].Name
		}
		return teamList[i].Result["P"] > teamList[j].Result["P"]
	})

	fmt.Fprintf(writer, "%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, team := range teamList {
		fmt.Fprintf(writer, "%-31s| %2d |  %1d |  %1d |  %1d | %2d\n",
			team.Name,
			team.Result["MP"],
			team.Result["W"],
			team.Result["D"],
			team.Result["L"],
			team.Result["P"])
	}

	return nil
}
