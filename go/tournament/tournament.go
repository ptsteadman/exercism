package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

// TeamRecord represents a team's sum of wins, losses and draws for a season
type TeamRecord struct {
	Name   string
	Wins   int
	Losses int
	Draws  int
}

// GetPoints is used to rank teams against each other
func (record TeamRecord) GetPoints() int {
	return record.Wins*3 + record.Draws
}

// GetMatchesPlayed returns the total matches played for a team
func (record TeamRecord) GetMatchesPlayed() int {
	return record.Wins + record.Draws + record.Losses
}

// Tally reads lines of game results from an input io.Reader r, sums up statistics,
// and writes a sorted results table to io.Writer w
func Tally(r io.Reader, w io.Writer) error {
	records := make(map[string]*TeamRecord) // team name -> record
	scanner := bufio.NewScanner(r)

	// sum up win/loss/draw records based on the input
	for scanner.Scan() {
		l := scanner.Text()
		isCommentOrWhitespace, _ := regexp.MatchString(`^$|^\s+$|^\s+\#|^\#`, l)
		if isCommentOrWhitespace {
			continue
		}
		result := strings.Split(l, ";")
		// validate line
		if len(result) != 3 {
			return errors.New("invalid line of input")
		}
		if result[0] == result[1] {
			return errors.New("team cannot play itself")
		}
		// create new team records if none exist yet
		if _, ok := records[result[0]]; !ok {
			records[result[0]] = &TeamRecord{Name: result[0]}
		}
		if _, ok := records[result[1]]; !ok {
			records[result[1]] = &TeamRecord{Name: result[1]}
		}
		// update tally
		switch result[2] {
		case "win":
			records[result[0]].Wins++
			records[result[1]].Losses++
		case "loss":
			records[result[0]].Losses++
			records[result[1]].Wins++
		case "draw":
			records[result[1]].Draws++
			records[result[0]].Draws++
		default:
			return errors.New("invalid game outcome")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// sort records by points (or alphabetically if tie)
	sortedRecords := []TeamRecord{}
	for _, v := range records {
		sortedRecords = append(sortedRecords, *v)
	}
	sort.Slice(sortedRecords, func(i, j int) bool {
		if sortedRecords[i].GetPoints() == sortedRecords[j].GetPoints() {
			return sortedRecords[i].Name < sortedRecords[j].Name
		}
		return sortedRecords[i].GetPoints() > sortedRecords[j].GetPoints()
	})

	// output tab-formatted results
	writer := tabwriter.NewWriter(w, 1, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Team\t MP \t  W \t  D \t  L \t  P")
	for _, record := range sortedRecords {
		fmt.Fprintf(
			writer,
			"%s        \t  %d\t  %d\t  %d\t  %d\t  %d\n",
			record.Name,
			record.GetMatchesPlayed(),
			record.Wins,
			record.Draws,
			record.Losses,
			record.GetPoints(),
		)
	}
	writer.Flush()

	return nil
}
